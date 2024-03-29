package luadns

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	jsonMime = "application/json"              // Accept, Content-Type header
	timeout  = time.Duration(10 * time.Second) // Request timeout
)

// JSONClient represents a REST client using JSON format.
type JSONClient struct {
	ctx      context.Context
	client   *http.Client
	username string
	password string
}

// NewJSONClient initializes JSON client.
func NewJSONClient(ctx context.Context, client *http.Client) *JSONClient {
	return &JSONClient{
		ctx:    ctx,
		client: client,
	}
}

// NewAuthJSONClient initializes JSON client using Basic authentication.
func NewAuthJSONClient(ctx context.Context, username, password string) *JSONClient {
	client := &http.Client{
		Transport: &Transport{
			Transport: &http.Transport{},
			username:  username,
			password:  password,
		},
		Timeout: timeout,
	}
	return NewJSONClient(ctx, client)
}

// Post executes a POST request using JSON body and returns JSON response.
func (c *JSONClient) Post(url string, attrs interface{}) ([]byte, error) {
	json, err := c.marshalJSON(attrs)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewBuffer(json)

	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

// Get executes a GET request and returns JSON response.
func (c *JSONClient) Get(url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

// Put executes a PUT request using JSON body and returns JSON response.
func (c *JSONClient) Put(url string, data interface{}) ([]byte, error) {
	json, err := c.marshalJSON(data)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewBuffer(json)

	req, err := http.NewRequestWithContext(c.ctx, http.MethodPut, url, payload)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

// Delete executes a DELETE request and returns JSON response.
func (c *JSONClient) Delete(url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodDelete, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// do executes HTTP request, checks for proper response and returns the response body.
func (c *JSONClient) do(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = c.checkStatusCode(resp)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, jsonMime) {
		return nil, &ErrBadContentType{ContentType: contentType}
	}

	return body, nil
}

// checkStatusCode checks the HTTP status code and maps to corresponding error.
func (c *JSONClient) checkStatusCode(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusTooManyRequests:
		limit, err := c.getRatelimitValue(resp, "X-Ratelimit-Limit")
		if err != nil {
			return err
		}
		reset, err := c.getRatelimitValue(resp, "X-Ratelimit-Reset")
		if err != nil {
			return err
		}
		return &ErrTooManyRequests{
			Limit: limit,
			Reset: reset,
		}
	default:
		return &ErrBadStatusCode{StatusCode: resp.StatusCode}
	}
}

// getRatelimitValue parses values from X-Ratelimit-* headers.
func (c *JSONClient) getRatelimitValue(resp *http.Response, key string) (int64, error) {
	value := resp.Header.Get(key)

	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (c *JSONClient) marshalJSON(payload interface{}) ([]byte, error) {
	return json.Marshal(payload)
}
