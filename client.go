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
	// MaxRetries is the maximum number of times to retry a request.
	MaxRetries int
}

func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	if c.MaxRetries <= 0 {
		c.MaxRetries = 0
	}
	for i := 0; i <= c.MaxRetries; i++ {
		resp, err = c.Client.Do(req)
		if err == nil {
			break
		}
	}
	return resp, err
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
