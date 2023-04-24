package postapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strings"

	s "github.com/neurosnap/sentences"
	"github.com/neurosnap/sentences/english"
	tke "github.com/pkoukk/tiktoken-go"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"io"

	"code.sajari.com/docconv"
)

type Chunk struct {
	Start int
	End   int
	Title string
	Text  string
}

// MaxTokensPerChunk is the maximum number of tokens allowed in a single chunk for OpenAI embeddings
const MaxTokensPerChunk = 500
const EmbeddingModel = "text-embedding-ada-002"

func CreateChunks(fileContent string, title string) ([]Chunk, error) {
	tokenizer, _ := english.NewSentenceTokenizer(nil)
	sentences := tokenizer.Tokenize(fileContent)

	log.Println("[CreateChunks] getting tiktoken for", EmbeddingModel, "...")
	// Get tiktoken encoding for the model
	tiktoken, err := tke.EncodingForModel(EmbeddingModel)
	if err != nil {
		return []Chunk{}, fmt.Errorf("getEncoding: %v", err)
	}

	newData := make([]Chunk, 0)
	position := 0
	i := 0

	for i < len(sentences) {
		chunkTokens := 0
		chunkSentences := []*s.Sentence{}

		// Add sentences to the chunk until the token limit is reached
		for i < len(sentences) {
			tiktokens := tiktoken.Encode(sentences[i].Text, nil, nil)
			tokenCount := len(tiktokens)
			fmt.Printf(
				"[CreateChunks] #%d Token count: %d | Total number of sentences: %d | Sentence: %s\n",
				i, tokenCount, len(sentences), sentences[i].Text)

			if chunkTokens+tokenCount <= MaxTokensPerChunk {
				chunkSentences = append(chunkSentences, sentences[i])
				chunkTokens += tokenCount
				i++
			} else {
				log.Println("[CreateChunks] Adding this sentence would exceed max token limit. Breaking....")
				break
			}
		}

		if len(chunkSentences) > 0 {
			text := strings.Join(sentencesToStrings(chunkSentences), "")

			start := position
			end := position + len(text)

			fmt.Printf("[CreateChunks] Created chunk and adding it to the array...\nText: %s\n",
				text)

			newData = append(newData, Chunk{
				Start: start,
				End:   end,
				Title: title,
				Text:  text,
			})
			fmt.Printf("[CreateChunks] New chunk array length: %d\n",
				len(newData))
			position = end

			// Set the stride for overlapping chunks
			stride := len(chunkSentences) / 2
			if stride < 1 {
				stride = 1
			}

			oldI := i
			i -= stride

			// Check if the next sentence would still fit within the token limit
			nextTokens := tiktoken.Encode(sentences[i].Text, nil, nil)
			nextTokenCount := len(nextTokens)

			if chunkTokens+nextTokenCount <= MaxTokensPerChunk {
				// Increment i without applying the stride
				i = oldI + 1
			} else if i == oldI {
				// Ensure i is always incremented to avoid an infinite loop
				i++
			}

		}
	}

	return newData, nil
}

func sentencesToStrings(sentences []*s.Sentence) []string {
	strs := make([]string, len(sentences))
	for i, s := range sentences {
		strs[i] = s.Text
	}
	return strs
}

func getTextFromFile(f multipart.File) (string, error) {
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	utf16bom := unicode.BOMOverride(unicode.UTF8.NewDecoder())
	fileString, _, err := transform.String(utf16bom, string(fileBytes))
	if err != nil {
		return "", err
	}

	fmt.Printf("[getTextFromFile] fileBytes length: %d | fileString: %+v \n", len(fileBytes), fileString)

	return fileString, nil
}

// extract human-readable text from a given pdf with support for spaces/whitespace.
func extractTextFromPDF(f multipart.File, fileSize int64) (string, error) {
	// Reset the file reader's position
	_, err := f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	// Convert the uploaded file to a human-readable text
	bodyResult, _, err := docconv.ConvertPDF(f)
	if err != nil {
		return "", err
	}

	// Remove extra whitespace and newlines
	text := strings.TrimSpace(bodyResult)

	return text, nil
}
