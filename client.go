package openai

import (
	"crypto/tls"
	"net/http"
)

// Client is the OpenAI API C.
type Client struct {
	// HTTP client used to communicate with the API.
	Client *http.Client
	// API key used for authentication.
	APIKey string
}

// NewClient returns a new OpenAI API C
func NewClient(apikey string, client *http.Client) *Client {
	return &Client{
		Client: client,
		APIKey: apikey,
	}
}

// DefaultClient returns a new OpenAI API C with the default HTTP client.
func DefaultClient(apikey string) *Client {
	client := &http.Client{}
	tp := http.DefaultTransport.(*http.Transport).Clone()
	// Ignore certificate verification.
	tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client.Transport = tp
	return NewClient(apikey, client)
}
