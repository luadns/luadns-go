package luadns_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luadns/luadns-go"
	"github.com/stretchr/testify/assert"
)

func TestUsersMeEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sendHTTPFixture(t, "/users/me.show", w, r)
	}))
	defer server.Close()

	c := luadns.NewClient("joe@example.com", "password", luadns.SetBaseURL(server.URL))

	user, err := c.Me(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, user.Email, "joe@example.com")
	assert.Equal(t, user.Name, "Example User")
	assert.Equal(t, user.RepoURI, "")
	assert.True(t, user.APIEnabled)
	assert.False(t, user.TFA)
	assert.Regexp(t, "^ssh-rsa AAAAB3NzaC1yc2.*", user.DeployKey)
	assert.Equal(t, user.TTL, uint32(300))
	assert.Equal(t, user.Package, "Free")
	assert.Equal(t, user.NameServers, []string{"ns1.luadns.net.", "ns2.luadns.net.", "ns3.luadns.net.", "ns4.luadns.net."})
}
