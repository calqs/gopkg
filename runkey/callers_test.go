package runkey

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// requestFrom is a request from remoteAddr, through a reverse proxy when
// forwardedFor is set.
func requestFrom(remoteAddr, forwardedFor string) *http.Request {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/boards", nil)
	req.RemoteAddr = remoteAddr
	if forwardedFor != "" {
		req.Header.Set("X-Forwarded-For", forwardedFor)
	}
	return req
}

func Test_Loopback_Allows_This_Machine_Only(t *testing.T) {
	cases := map[string]struct {
		remoteAddr, forwardedFor string
		allowed                  bool
	}{
		"IPv4 loopback":                  {remoteAddr: "127.0.0.1:51000", allowed: true},
		"IPv6 loopback":                  {remoteAddr: "[::1]:51000", allowed: true},
		"IPv4 loopback mapped in IPv6":   {remoteAddr: "[::ffff:127.0.0.1]:51000", allowed: true},
		"another machine":                {remoteAddr: "203.0.113.7:51000"},
		"the reverse proxy on this host": {remoteAddr: "127.0.0.1:51000", forwardedFor: "203.0.113.7"},
		"no port":                        {remoteAddr: "127.0.0.1"},
	}
	for name, c := range cases {
		assert.Equal(t, c.allowed, Loopback().Allow(requestFrom(c.remoteAddr, c.forwardedFor)), name)
	}
}

func Test_ParseCallers_Takes_Its_Networks(t *testing.T) {
	callers, err := ParseCallers("172.16.0.0/12, 192.168.0.0/16")
	require.NoError(t, err)
	assert.True(t, callers.Allow(requestFrom("172.17.0.1:51000", "")))
	assert.True(t, callers.Allow(requestFrom("192.168.65.1:51000", "")))
	// loopback is no longer implied
	assert.False(t, callers.Allow(requestFrom("127.0.0.1:51000", "")))
}

func Test_ParseCallers_Defaults_To_Loopback(t *testing.T) {
	callers, err := ParseCallers(" ")
	require.NoError(t, err)
	assert.Equal(t, Loopback(), callers)
}

func Test_ParseCallers_Refuses_A_Bad_Network(t *testing.T) {
	_, err := ParseCallers("127.0.0.1")
	assert.Error(t, err)
}

func Test_A_Zero_Callers_Allows_Nothing(t *testing.T) {
	assert.False(t, Callers{}.Allow(requestFrom("127.0.0.1:51000", "")))
}
