package vectordb

import (
	"github.com/pashpashpash/vault/chunk"
)

type QueryMatch struct {
	ID       string            `json:"id"`
	Score    float32           `json:"score"` // Use "score" instead of "distance"
	Metadata map[string]string `json:"metadata"`
}

type VectorDB interface {
	UpsertEmbeddings(embeddings [][]float32, chunks []chunk.Chunk, uuid string) error
	Retrieve(questionEmbedding []float32, topK int, uuid string) ([]QueryMatch, error)
}
