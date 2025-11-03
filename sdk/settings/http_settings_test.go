package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeBaseURL(t *testing.T) {
	// Valid transformations
	cases := []struct{ in, out string }{
		{"https://h", "https://h/api"},
		{"https://h/", "https://h/api"},
		{"https://h//", "https://h/api"},
		{"https://h/api", "https://h/api"},
		{"https://h/api/", "https://h/api"},
		{"https://h/conductor", "https://h/conductor/api"},
		{"https://h/conductor/", "https://h/conductor/api"},
		{"http://[::1]:8080", "http://[::1]:8080/api"},
		{"https://h?x=1", "https://h/api"},  // query dropped
		{"https://h#frag", "https://h/api"}, // fragment dropped
		{"", defaultBaseUrl},                // empty -> default
		{"h", "h"},                          // unchanged
		{"localhost", "localhost"},          // unchanged
	}
	for _, c := range cases {
		hs := &HttpSettings{BaseUrl: c.in}
		hs.normalizeBaseURL()
		assert.Equal(t, c.out, hs.BaseUrl, "%s -> %s (want %s)", c.in, hs.BaseUrl, c.out)
	}
}
