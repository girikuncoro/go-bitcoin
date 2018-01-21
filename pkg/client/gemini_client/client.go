package gemini_client

import (
	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"
)

type GeminiClient struct {
	baseClient *base_client.BaseClient
}

func NewGeminiClient(baseClient *base_client.BaseClient) *GeminiClient {
	return &GeminiClient{
		baseClient: baseClient,
	}
}
