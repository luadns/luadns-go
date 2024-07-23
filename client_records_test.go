package luadns_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luadns/luadns-go"
	"github.com/stretchr/testify/assert"
)

func TestListRecordsEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5/records.index", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	records, err := c.ListRecords(context.Background(), &luadns.Zone{ID: 5}, &luadns.ListParams{Query: "example.org."})
	assert.NoError(t, err)
	assert.Len(t, records, 11)

	record := records[0]
	assert.Equal(t, record.ID, int64(115014343))
	assert.Equal(t, record.Name, "example.org.")
	assert.Equal(t, record.Type, "SOA")
	assert.Equal(t, record.Content, "ns1.luadns.net. hostmaster.luadns.net. 1692975563 1200 120 604800 3600")
	assert.Equal(t, record.TTL, uint32(3600))
}

func TestCreateRecordEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5/records.create", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	f := &luadns.Record{Name: "example.org.", Type: "TXT", Content: "Hello, world!", TTL: 3600}

	record, err := c.CreateRecord(context.Background(), &luadns.Zone{ID: 5}, f)
	assert.NoError(t, err)
	assert.Equal(t, record.ID, int64(115087858))
	assert.Equal(t, record.Name, "example.org.")
	assert.Equal(t, record.Type, "TXT")
	assert.Equal(t, record.Content, "Hello, world!")
	assert.Equal(t, record.TTL, uint32(3600))
}

func TestCreateRecordEndpointWithInvalidData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5/records.create:err", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	f := &luadns.Record{Name: "example.org.", Type: "A", Content: "invalid", TTL: 3600}
	_, err := c.CreateRecord(context.Background(), &luadns.Zone{ID: 5}, f)
	assert.EqualError(t, err, "Invalid data for content: invalid IPv4 address")
}

func TestGetRecordEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5/records/115014348.show", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))

	record, err := c.GetRecord(context.Background(), &luadns.Zone{ID: 5}, 115014348)
	assert.NoError(t, err)
	assert.Equal(t, record.ID, int64(115014348))
	assert.Equal(t, record.Name, "example.org.")
	assert.Equal(t, record.Type, "A")
	assert.Equal(t, record.Content, "1.1.1.1")
	assert.Equal(t, record.TTL, uint32(86400))
}

func TestUpdateRecordEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5/records/115014348.update", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	f := &luadns.Record{Name: "example.org.", Type: "A", Content: "2.2.2.2", TTL: 86400}

	record, err := c.UpdateRecord(context.Background(), &luadns.Zone{ID: 5}, 115014348, f)
	assert.NoError(t, err)
	assert.Equal(t, record.ID, int64(115014348))
	assert.Equal(t, record.Name, "example.org.")
	assert.Equal(t, record.Type, "A")
	assert.Equal(t, record.Content, "2.2.2.2")
	assert.Equal(t, record.TTL, uint32(86400))
}

func TestDeleteRecordEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5/records/115014348.delete", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))

	record, err := c.DeleteRecord(context.Background(), &luadns.Zone{ID: 5}, 115014348)
	assert.NoError(t, err)
	assert.Equal(t, record.ID, int64(115014348))
	assert.Equal(t, record.Name, "example.org.")
	assert.Equal(t, record.Type, "A")
	assert.Equal(t, record.Content, "1.1.1.1")
	assert.Equal(t, record.TTL, uint32(86400))
}
