package postapi

import (
	"errors"
	"fmt"
	"io/ioutil"
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
// MaxTokensPerChunk is the maximum number of tokens allowed in a single chunk for OpenAI embeddings
const MaxTokensPerChunk = 1000
const EmbeddingModel = "text-embedding-ada-002"

// Take file content, split it into sentences using neurosnap/sentences library
// Start building chunks up to the MaxTokensPerChunk token limit
// Use tiktoken-go to estimate token count
// Once a chunk is full, move on to the next chunk until the entire content is covered
// Use a dynamically set stride so that Chunks overlap each other by 10 sentences
// Skip any sentence that is longer than MaxTokensPerChunk
// All returned Chunks must have a non-empty Text contents
// All lengths of files/sentences should be handled, including edge cases
// There should be no chance of looping forever
func CreateChunks(fileContent string, title string) ([]Chunk, error) {
	chunks := []Chunk{}

	// Initialize sentence tokenizer
	tokenizer, _ := english.NewSentenceTokenizer(nil)
	sentences := tokenizer.Tokenize(fileContent)

	// Get tiktoken encoding for the model
	tiktoken, err := tke.EncodingForModel(EmbeddingModel)
	if err != nil {
		return []Chunk{}, fmt.Errorf("getEncoding: %v", err)
	}

	chunkStart := 0

	for chunkStart < len(sentences) {
		tokenCount := 0
		chunkText := ""
		chunkSentences := 0

		for i := chunkStart; i < len(sentences) && tokenCount < MaxTokensPerChunk; i++ {
			sentence := sentences[i].Text
			tiktokens := tiktoken.Encode(sentence, nil, nil)
			sentenceTokenCount := len(tiktokens)

			if sentenceTokenCount > MaxTokensPerChunk {
				continue // Skip sentence if longer than MaxTokensPerChunk
			}

			if tokenCount+sentenceTokenCount <= MaxTokensPerChunk {
				tokenCount += sentenceTokenCount
				chunkText += " " + sentence
				chunkSentences++
			} else {
				break
			}
		}

		trimmedText := strings.TrimSpace(chunkText)
		if len(trimmedText) > 0 {
			chunks = append(chunks, Chunk{
				Start: chunkStart,
				End:   chunkStart + tokenCount,
				Title: title,
				Text:  trimmedText,
			})
		}

		// Calculate stride dynamically based on chunk sentences
		sentenceStride := chunkSentences / 5
		if sentenceStride == 0 {
			sentenceStride = 1
		}

		// Move chunkStart forward by sentenceStride
		chunkStart += sentenceStride
	}

	if len(chunks) == 0 {
		return nil, errors.New("no chunks created")
	}

	return chunks, nil
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
