package luadns_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luadns/luadns-go"
	"github.com/stretchr/testify/assert"
)

func TestListZonesEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones.index", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))

	zones, err := c.ListZones(context.Background(), &luadns.ListParams{Query: "example.com"})
	assert.NoError(t, err)
	assert.Len(t, zones, 1)

	zone := zones[0]
	assert.Equal(t, zone.ID, int64(5))
	assert.Equal(t, zone.Name, "example.org")
}

func TestCreateZoneEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones.create", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	f := &luadns.Zone{Name: "example.dev"}

	zone, err := c.CreateZone(context.Background(), f)
	assert.NoError(t, err)
	assert.Equal(t, zone.ID, int64(75247))
	assert.Equal(t, zone.Name, "example.dev")
	assert.Len(t, zone.Records, 5)
}

func TestCreateZoneEndpointWithInvalidData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones.create:err", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	f := &luadns.Zone{Name: ""}

	_, err := c.CreateZone(context.Background(), f)
	assert.EqualError(t, err, "Invalid data for name: Required; Invalid data for name: invalid name")
}

func TestCreateZoneEndpointWithForbiddenResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones.create:err-forbidden", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	f := &luadns.Zone{Name: "example.org"}

	_, err := c.CreateZone(context.Background(), f)
	assert.EqualError(t, err, "Forbidden: Zone 'example.org' is taken already.")
}

func TestGetZoneEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5.show", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))

	zone, err := c.GetZone(context.Background(), 5)
	assert.NoError(t, err)
	assert.Equal(t, zone.ID, int64(5))
	assert.Equal(t, zone.Name, "example.org")
	assert.Len(t, zone.Records, 11)
}

func TestUpdateZoneEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5.update", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))
	f := &luadns.Zone{Name: "example.org"}

	zone, err := c.UpdateZone(context.Background(), 5, f)
	assert.NoError(t, err)
	assert.Equal(t, zone.ID, int64(5))
	assert.Equal(t, zone.Name, "example.org")
	assert.Len(t, zone.Records, 5)
}

func TestDeleteZoneEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/zones/5.delete", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))

	zone, err := c.DeleteZone(context.Background(), 5)
	assert.NoError(t, err)
	assert.Equal(t, zone.ID, int64(5))
	assert.Equal(t, zone.Name, "example.org")
	assert.Len(t, zone.Records, 11)
}
