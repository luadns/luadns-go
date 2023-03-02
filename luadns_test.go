package luadns_test

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func endpointPath(filename string) string {
	ext := path.Ext(filename)
	name := filename[0 : len(filename)-len(ext)]
	return name
}

func readHTTPFixture(name string) ([]byte, error) {
	filename := "testdata/http" + name
	return os.ReadFile(filename)
}

func sendHTTPFixture(t *testing.T, filename string, w http.ResponseWriter, r *http.Request) {
	path := endpointPath(filename)
	assert.Equal(t, r.URL.Path, path)
	assert.Equal(t, r.Header.Get("Accept"), "application/json")
	assert.Equal(t, r.Header.Get("Authorization"), "Basic am9lQGV4YW1wbGUuY29tOnBhc3N3b3Jk")

	data, err := readHTTPFixture(filename)
	assert.NoError(t, err)

	resp, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(data)), r)
	assert.NoError(t, err)

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Set(name, value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	assert.NoError(t, err)

	flusher.Flush()
}
