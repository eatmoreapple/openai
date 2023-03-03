package openai

import (
	"net/http"
)

// Client is the OpenAI API client.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client
	// API key used for authentication.
	APIKey string
}

// NewClient returns a new OpenAI API client
func NewClient(apikey string) *Client {
	return &Client{
		client: http.DefaultClient,
		APIKey: apikey,
	}
}
