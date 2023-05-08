package postapi

import (
	"github.com/pashpashpash/vault/vectordb"

	cache "github.com/patrickmn/go-cache"
	openai "github.com/sashabaranov/go-openai"
)

type HandlerContext struct {
	openAIClient *openai.Client
	cache        *cache.Cache
	vectorDB     vectordb.VectorDB
}

func NewHandlerContext(openAIClient *openai.Client, vectorDB vectordb.VectorDB) *HandlerContext {
	return &HandlerContext{
		openAIClient: openAIClient,
		cache:        cache.New(cache.NoExpiration, cache.NoExpiration),
		vectorDB:     vectorDB,
	}
}
