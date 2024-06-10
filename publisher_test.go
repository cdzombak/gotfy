package gotfy

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type FakeHttpClient struct {
	CheckDo  func(req *http.Request)
	Response func() (*http.Response, error)
}

func (m *FakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	if m.CheckDo != nil {
		m.CheckDo(req)
	}
	if m.Response != nil {
		return m.Response()
	}

	respReq := req.Clone(context.Background())
	respReq.Body = nil
	return &http.Response{
		Status:     "204 No content",
		StatusCode: 204,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Close:      true,
		Request:    respReq,
	}, nil
}

func Test_Publisher_SetsJSONHeaders(t *testing.T) {
	r := require.New(t)
	c := FakeHttpClient{}
	sut := NewPublisher(PublisherOpts{
		HttpClient: &c,
	})

	c.CheckDo = func(req *http.Request) {
		r.Equal("application/json", req.Header.Get("Content-Type"))
		r.Equal("application/json", req.Header.Get("Accept"))
	}
	_, _ = sut.Send(context.Background(), Message{})
}

func Test_Publisher_AuthToken(t *testing.T) {
	r := require.New(t)
	c := FakeHttpClient{}
	sut := NewPublisher(PublisherOpts{
		HttpClient: &c,
		Auth:       AccessToken("tk_0123456789"),
	})

	c.CheckDo = func(req *http.Request) {
		r.Equal("Bearer tk_0123456789", req.Header.Get("Authorization"))
	}
	_, _ = sut.Send(context.Background(), Message{})
}
