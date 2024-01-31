package syncapi

import (
	"net/http"
)

const apiKeyHeader = "X-API-Key"

type ApiClient struct {
	apiKey string
	client *http.Client
	host   string
}

func NewApiClient(apiKey string, host string) (*ApiClient, error) {
	return &ApiClient{
		apiKey: apiKey,
		client: &http.Client{},
		host:   host,
	}, nil
}
