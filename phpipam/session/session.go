// Package session provides session management utility and token storage.
package session

import (
	"github.com/imdario/mergo"
	"github.com/paybyphone/phpipam-sdk-go/phpipam"
	"net/http"
	"sync"
)

// timeLayout represents the datetime format returned by the PHPIPAM api.
const timeLayout = "2006-01-02 15:04:05"

// Token represents a PHPIPAM session token.
type Token struct {
	// The token string.
	String string `json:"token"`
}

// Session represents a PHPIPAM session.
type Session struct {
	// The session's configuration.
	Config phpipam.Config

	// The session token.
	Token Token

	// A session shares this client if provided
	HttpClient *http.Client

	sync.RWMutex // protect updates of HttpClient

}

// NewSession creates a new session based off supplied configs. It is up to the
// client for each controller implementation to log in and refresh the token.
// This is provided in the base client.Client implementation.
func NewSession(configs ...phpipam.Config) *Session {
	s := &Session{
		Config: phpipam.DefaultConfigProvider(),
	}
	for _, v := range configs {
		mergo.Merge(&s.Config, v, mergo.WithOverride)
	}

	return s
}

// SetHttpClient sets a specific http client with a particular transport etc
// For backwards compatibility, a CheckRedirect function will be added at call time if not present
func (s *Session) SetHttpClient(hc *http.Client) {
	s.Lock()
	s.HttpClient = hc
	s.Unlock()
}

