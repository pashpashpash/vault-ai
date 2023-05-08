package postapi

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pashpashpash/vault/chunk"
	"github.com/pkoukk/tiktoken-go"
	openai "github.com/sashabaranov/go-openai"
)

const formatPrompt = `Take this unstructured data (pasted at the end of the prompt), and format it into a JSON object with the following structure:

{summary: string, answers[{shortAnswer: string, explanation: string, questionSummaryTwoWordsMax: string}]}

I want the short answer to be Yes/No/Uncertain, and the explanation to be what is put after it.
If there is no short answer, infer whether it is Yes/No/Uncertain based on the explanation. The only acceptable contents of the shortAnswer string is "Yes", "No", or "Uncertain". Each answer object must have non-empty shortAnswer, explanation, and questionSummaryTwoWordsMax variables that are non-null. For the question summary, provide a one to two word string that summarizes the question the answer is responding to. The questionSummaryTwoWordsMax string must not exceed two words. The size of the answers array must be the same as the number of questions.

For some context, here were the input questions:
%s
Unstructured Data:
%s
`

type OpenAIResponse struct {
	Response string `json:"response"`
	Tokens   int    `json:"tokens"`
}

func callOpenAI(client *openai.Client, prompt string, model string,
	instructions string, maxTokens int) (string, int, error) {
	// set request details
	temperature := float32(0.7)
	topP := float32(1.0)
	frequencyPenalty := float32(0.0)
	presencePenalty := float32(0.6)
	stop := []string{"Human:", "AI:"}

	var assistantMessage string
	var tokens int
	var err error
	if model == openai.GPT3TextDavinci003 {
		prompt = "System Instructions:\n" + instructions + "\n\nPrompt:\n" + prompt
		assistantMessage, tokens, err = useCompletionAPI(client, prompt, model, temperature,
			maxTokens, topP, frequencyPenalty, presencePenalty, stop)
	} else {
		assistantMessage, tokens, err = useChatCompletionAPI(client, prompt, model, instructions, temperature,
			maxTokens, topP, frequencyPenalty, presencePenalty, stop)
	}

	return assistantMessage, tokens, err
}

func useChatCompletionAPI(client *openai.Client, prompt, modelParam string, instructions string, temperature float32, maxTokens int, topP float32, frequencyPenalty, presencePenalty float32, stop []string) (string, int, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: instructions,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            modelParam,
			Messages:         messages,
			Temperature:      temperature,
			MaxTokens:        maxTokens,
			TopP:             topP,
			FrequencyPenalty: frequencyPenalty,
			PresencePenalty:  presencePenalty,
			Stop:             stop,
		},
	)

	if err != nil {
		return "", 0, err
	}

	return resp.Choices[0].Message.Content, resp.Usage.TotalTokens, nil
}

func useCompletionAPI(client *openai.Client, prompt, modelParam string,
	temperature float32, maxTokens int, topP float32,
	frequencyPenalty, presencePenalty float32,
	stop []string) (string, int, error) {
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:            modelParam,
			Prompt:           prompt,
			Temperature:      temperature,
			MaxTokens:        maxTokens,
			TopP:             topP,
			FrequencyPenalty: frequencyPenalty,
			PresencePenalty:  presencePenalty,
			Stop:             stop,
		},
	)

	if err != nil {
		return "", 0, err
	}

	return resp.Choices[0].Text, resp.Usage.TotalTokens, nil
}

func callEmbeddingAPIWithRetry(client *openai.Client, texts []string, embedModel openai.EmbeddingModel,
	maxRetries int) (*openai.EmbeddingResponse, error) {
	var err error
	var res openai.EmbeddingResponse

	for i := 0; i < maxRetries; i++ {
		res, err = client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
			Input: texts,
			Model: embedModel,
		})

		if err == nil {
			return &res, nil
		}

		time.Sleep(5 * time.Second)
	}

	return nil, err
}

func getEmbeddings(client *openai.Client, chunks []chunk.Chunk, batchSize int,
	embedModel openai.EmbeddingModel) ([][]float32, error) {
	embeddings := make([][]float32, 0, len(chunks))

	for i := 0; i < len(chunks); i += batchSize {
		iEnd := min(len(chunks), i+batchSize)

		texts := make([]string, 0, iEnd-i)
		for _, chunk := range chunks[i:iEnd] {
			texts = append(texts, chunk.Text)
		}

		log.Println("[getEmbeddings] Feeding texts to Openai to get embedding...\n", texts)

		res, err := callEmbeddingAPIWithRetry(client, texts, embedModel, 3)
		if err != nil {
			return nil, err
		}

		embeds := make([][]float32, len(res.Data))
		for i, record := range res.Data {
			embeds[i] = record.Embedding
		}

		embeddings = append(embeddings, embeds...)
	}

	return embeddings, nil
}

func getEmbedding(client *openai.Client, text string, embedModel openai.EmbeddingModel) ([]float32, error) {
	res, err := callEmbeddingAPIWithRetry(client, []string{text}, embedModel, 3)
	if err != nil {
		return nil, err
	}

	return res.Data[0].Embedding, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildPrompt(contexts []string, question string) (string, error) {
	tokenLimit := 3750
	promptStart := "Answer the question based on the context below.\n\nContext:\n"
	promptEnd := fmt.Sprintf("\n\nQuestion: %s\nAnswer:", question)

	// Get tiktoken encoding for the model
	tke, err := tiktoken.EncodingForModel("davinci")
	if err != nil {
		return "", fmt.Errorf("getEncoding: %v", err)
	}

	// Count tokens for the question
	questionTokens := tke.Encode(question, nil, nil)
	currentTokenCount := len(questionTokens)

	var prompt string
	for i := range contexts {
		// Count tokens for the current context
		contextTokens := tke.Encode(contexts[i], nil, nil)
		currentTokenCount += len(contextTokens)

		if currentTokenCount >= tokenLimit {
			prompt = promptStart + strings.Join(contexts[:i], "\n\n---\n\n") + promptEnd
			break
		} else if i == len(contexts)-1 {
			prompt = promptStart + strings.Join(contexts, "\n\n---\n\n") + promptEnd
		}
	}

	return prompt, nil
}
