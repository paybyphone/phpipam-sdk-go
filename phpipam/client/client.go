// Package client contains generic client structs and methods that are
// designed to be used by specific PHPIPAM services and resources.
package client

import (
	"fmt"

	"github.com/paybyphone/phpipam-sdk-go/phpipam/request"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
)

// Client encompasses a generic client object that is further extended by
// services. Any common configuration and functionality goes here.
type Client struct {
	// The session for this client.
	Session *session.Session
}

// NewClient creates a new client.
func NewClient(s *session.Session) *Client {
	c := &Client{
		Session: s,
	}
	return c
}

// loginSession logs in a session via the user controller. This is the only
// valid operation if the session does not have a token yet.
func loginSession(s *session.Session) error {
	var out session.Token
	r := request.NewRequest(s)
	r.Method = "POST"
	r.URI = fmt.Sprintf("/%s/user/", s.Config.AppID)
	r.Input = &struct{}{}
	r.Output = &out
	if err := r.Send(); err != nil {
		return err
	}
	s.Token = out
	return nil
}

// refreshSession refreshes the session by sending a PATCH to the user
// controller, refreshing the existing token.
func refreshSession(s *session.Session) error {
	var out session.Token
	r := request.NewRequest(s)
	r.Method = "PATCH"
	r.URI = fmt.Sprintf("/%s/user/", s.Config.AppID)
	r.Input = &struct{}{}
	r.Output = &out
	if err := r.Send(); err != nil {
		return err
	}
	s.Token = out
	return nil
}

// SendRequest sends a request to a request.Request object.  It's expected that
// references to specific data types are passed - no checking is done to make
// sure that references are passed.
//
// This function also wraps session management into the workflow, logging in
// and refreshing session tokens as needed.
func (c *Client) SendRequest(method, uri string, in, out interface{}) error {
	// Check to make sure our session is ok first.
	switch {
	case c.Session.Token.String == "":
		if err := loginSession(c.Session); err != nil {
			return fmt.Errorf("Error logging into PHPIPAM: %s", err)
		}
	case c.Session.IsExpired():
		if err := refreshSession(c.Session); err != nil {
			return fmt.Errorf("Error refreshing PHPIPAM session token: %s", err)
		}
	}

	r := request.NewRequest(c.Session)
	r.Method = method
	r.URI = uri
	r.Input = in
	r.Output = out
	if err := r.Send(); err != nil {
		return err
	}
	return nil
}
