package gotfy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Publisher sends notification messages to a Ntfy server.
type Publisher interface {
	// Send publishes the given message to the configured Ntfy server.
	Send(ctx context.Context, m Message) (*SendResponse, error)
}

type publisher struct {
	server     url.URL
	headers    http.Header
	httpClient HttpClient
}

type HttpClient interface {
	// Do sends an HTTP request and returns an HTTP response.
	Do(req *http.Request) (*http.Response, error)
}

// PublisherOpts contains the configuration options for a new Publisher.
type PublisherOpts struct {
	Server     *url.URL
	Auth       Authorization
	Headers    http.Header
	HttpClient HttpClient
}

// NewPublisher creates a publisher for the given Ntfy server URL.
func NewPublisher(opts PublisherOpts) Publisher {
	retv := publisher{}

	if opts.Server == nil || opts.Server.String() == "" {
		retv.server = url.URL{
			Scheme: "https",
			Host:   "ntfy.sh",
		}
	} else {
		retv.server = *opts.Server
	}

	if opts.Headers == nil {
		retv.headers = make(http.Header)
	} else {
		retv.headers = opts.Headers.Clone()
	}
	retv.headers.Set("Content-Type", "application/json")
	retv.headers.Set("Accept", "application/json")

	if opts.Auth != nil {
		retv.headers.Set("Authorization", opts.Auth.Header())
	}

	if opts.HttpClient == nil {
		retv.httpClient = http.DefaultClient
	} else {
		retv.httpClient = opts.HttpClient
	}

	return &retv
}

// Send publishes the given message to the configured Ntfy server.
func (p *publisher) Send(ctx context.Context, m Message) (*SendResponse, error) {
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message to JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.server.String(), bytes.NewReader(buf))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header = p.headers.Clone()

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to send message: HTTP %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}

	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var pubResp SendResponse
	if err = json.Unmarshal(buf, &pubResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response from JSON: %w", err)
	}

	return &pubResp, nil
}
