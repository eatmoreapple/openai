package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type role string

const (
	RoleUser      role = "user"
	RoleSystem    role = "system"
	RoleAssistant role = "assistant"
)

// completionModel is the model used for completion.
type completionModel string

const (
	CompletionModelGPT35Turbo completionModel = "gpt-3.5-turbo"
)

// CompletionMessage is a message in a completion request.
type CompletionMessage struct {
	Role    role   `json:"role"`
	Content string `json:"content"`
}

// CompletionMessages is a slice of CompletionMessage.
type CompletionMessages []CompletionMessage

// CompletionRequest is the request to the OpenAI API's completion endpoint.
type CompletionRequest struct {
	Model    completionModel    `json:"model"`
	Messages CompletionMessages `json:"messages"`
}

// CompletionResponse is the response from the OpenAI API's completion endpoint.
type CompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// completionURL is the URL for the OpenAI API's completion endpoint.
const completionURL = "https://api.openai.com/v1/chat/completions"

// Completion returns a completion response for the given request.
func (c *Client) Completion(ctx context.Context, req CompletionRequest) (*CompletionResponse, error) {
	var buf = bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, completionURL, buf)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	buf.Reset()
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		// TODO: parse the response body and return a more specific error
		return nil, errors.New(buf.String())
	}
	var completionResponse CompletionResponse
	if err := json.NewDecoder(buf).Decode(&completionResponse); err != nil {
		return nil, err
	}
	return &completionResponse, nil
}

// CompletionWithPrompt returns a completion response for the given prompt.
func (c *Client) CompletionWithPrompt(ctx context.Context, prompt string) (*CompletionResponse, error) {
	return c.CompletionWithHistory(ctx, prompt)
}

// CompletionWithHistory returns a completion response for the given prompt and history.
func (c *Client) CompletionWithHistory(ctx context.Context, prompt string, histories ...CompletionMessage) (*CompletionResponse, error) {
	req := CompletionRequest{
		Model: CompletionModelGPT35Turbo,
		Messages: append(
			CompletionMessages{
				{
					Role:    RoleUser,
					Content: prompt,
				},
			},
			histories...,
		),
	}
	return c.Completion(ctx, req)
}
