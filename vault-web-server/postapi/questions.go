package postapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pashpashpash/vault/form"
	openai "github.com/sashabaranov/go-openai"
)

type Context struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type Answer struct {
	Answer  string    `json:"answer"`
	Context []Context `json:"context"`
	Tokens  int       `json:"tokens"`
}

// Handle Requests For Question
func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	form := new(form.QuestionForm)

	if errs := FormParseVerify(form, "QuestionForm", w, r); errs != nil {
		return
	}

	log.Println("[QuestionHandler] Question:", form.Question)
	log.Println("[QuestionHandler] Model:", form.Model)
	log.Println("[QuestionHandler] UUID:", form.UUID)
	log.Println("[QuestionHandler] ApiKey:", form.ApiKey)

	clientToUse := DEFAULT_OPENAI_CLIENT
	if form.ApiKey != "" {
		log.Println("[QuestionHandler] Using provided custom API key:", form.ApiKey)
		clientToUse = openai.NewClient(form.ApiKey)
	}

	// step 1: Feed question to openai embeddings api to get an embedding back
	questionEmbedding, err := getEmbedding(clientToUse, form.Question, openai.AdaEmbeddingV2)
	if err != nil {
		log.Println("[QuestionHandler ERR] OpenAI get embedding request error\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[QuestionHandler] Question Embedding Length:", len(questionEmbedding))

	// step 2: Query Pinecone using questionEmbedding to get context matches
	matches, err := retrieve(questionEmbedding, 4, form.UUID)
	if err != nil {
		log.Println("[QuestionHandler ERR] Pinecone query error\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("[QuestionHandler] Got matches from Pinecone:", matches)

	// Extract context text and titles from the matches
	contexts := make([]Context, len(matches))
	for i, match := range matches {
		contexts[i].Text = match.Metadata["text"]
		contexts[i].Title = match.Metadata["title"]
	}
	log.Println("[QuestionHandler] Retrieved context from Pinecone:\n", contexts)

	// step 3: Structure the prompt with a context section + question, using top x results from pinecone as the context
	contextTexts := make([]string, len(contexts))
	for i, context := range contexts {
		contextTexts[i] = context.Text
	}
	prompt, err := buildPrompt(contextTexts, form.Question)
	if err != nil {
		log.Println("[QuestionHandler ERR] Error building prompt\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	model := openai.GPT3Dot5Turbo
	if form.Model == "GPT Davinci" {
		model = openai.GPT3TextDavinci003
	}

	log.Printf("[QuestionHandler] Sending OpenAI api request...\nPrompt:%s\n", prompt)
	openAIResponse, tokens, err := callOpenAI(clientToUse, prompt, model,
		"You are a helpful assistant answering questions based on the context provided.",
		512)

	if err != nil {
		log.Println("[QuestionHandler ERR] OpenAI answer questions request error\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("[QuestionHandler] OpenAI response:\n", openAIResponse)
	response := OpenAIResponse{openAIResponse, tokens}

	answer := Answer{response.Response, contexts, response.Tokens}
	jsonResponse, err := json.Marshal(answer)
	if err != nil {
		log.Println("[QuestionHandler ERR] OpenAI response marshalling error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
