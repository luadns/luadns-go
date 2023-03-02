package luadns_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luadns/luadns-go"
	"github.com/stretchr/testify/assert"
)

func TestBadStatusCodeResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/users/me.show:err-bad-code", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient(context.Background(), "joe@example.com", "password", luadns.SetBaseURL(server.URL))

	_, err := c.Me()
	assert.Error(t, err)
	assert.EqualError(t, err, "Server returned bad status code (502)")
}

func TestBadContentResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/users/me.show:err-bad-content", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient(context.Background(), "joe@example.com", "password", luadns.SetBaseURL(server.URL))

	_, err := c.Me()
	assert.Error(t, err)
	assert.EqualError(t, err, "Server returned bad content type (text/html)")
}

func TestRateLimitedResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/users/me.show:err-too-many", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient(context.Background(), "joe@example.com", "password", luadns.SetBaseURL(server.URL))

	_, err := c.Me()
	assert.EqualError(t, err, "Too many requests, retry after 1693221300 unix time")

	rerr := err.(*luadns.ErrTooManyRequests)
	assert.Equal(t, rerr.Limit, int64(3))
	assert.Equal(t, rerr.Reset, int64(1693221300))
}
