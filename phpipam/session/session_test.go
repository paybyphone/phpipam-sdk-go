package session

import (
	"reflect"
	"testing"

	"github.com/paybyphone/phpipam-sdk-go/phpipam"
)

func phpipamConfig() phpipam.Config {
	return phpipam.Config{
		AppID:    "0123456789abcdefgh",
		Password: "changeit",
		Username: "nobody",
	}
}

func fullSessionConfig() *Session {
	return &Session{
		Config: phpipamConfig(),
		Token: Token{
			String:  "foobarbazboop",
			Expires: "2999-12-31 23:59:59",
		},
	}
}

func TestNewSession(t *testing.T) {
	cfg := phpipamConfig()

	expected := &Session{
		Config: phpipamConfig(),
	}
	expected.Config.Endpoint = "http://localhost/api"

	actual := NewSession(cfg)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected session to be %s, got %s", expected, actual)
	}
}

func TestIsExpiredFalse(t *testing.T) {
	sess := fullSessionConfig()
	if sess.IsExpired() {
		t.Fatal("Token should not be expired, unless we have travelled far into the future")
	}
}

func TestIsExpiredTrue(t *testing.T) {
	sess := fullSessionConfig()
	sess.Token.Expires = "1999-12-31 23:59:59"
	if !sess.IsExpired() {
		t.Fatal("Token should be expired, unless we have gone back in time")
	}
}
