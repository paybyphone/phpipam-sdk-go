// Package phpipam contains any top-level configuration structures
// necessary to work with the rest of the SDK and API.
package phpipam

import (
	"os"
	"strings"
)

// The default PHPIPAM API endpoint.
const defaultAPIAddress = "http://localhost/api"

// Config contains the configuration for connecting to the PHPIPAM API.
//
//
// Supplying Configuration to Controllers
//
// All controller constructors (ie: VLANs, subnets, addresses, etc) take zero or
// more of these structs as configuration, like so:
//
//   cfg := phpipam.Config{
//     Username:     "jdoe",
//     Password:     "password",
//     AppID:        "appid",
//   }
//   sess := session.New(cfg)
//   ctlr := ipaddr.New(sess)
//
// Note that default options are set for EmailAddress, Password, and AppKey.
// See the DefaultConfigProvider method for more details.
type Config struct {
	// The application ID required for API requests. This needs to be created in
	// the PHPIPAM console.
	AppID string

	// The API endpoint.
	Endpoint string

	// The password for the PHPIPAM account.
	Password string

	// The user name for the PHPIPAM account.
	Username string
}

// DefaultConfigProvider supplies a default configuration:
//  * AppID defaults to PHPIPAM_APP_ID, if set, otherwise empty
//  * Endpoint defaults to PHPIPAM_ENDPOINT_ADDR, otherwise http://localhost/api
//  * Password defaults to PHPIPAM_PASSWORD, if set, otherwise empty
//  * Username defaults to PHPIPAM_USER_NAME, if set, otherwise empty
//
// This essentially loads an initial config state for any given
// API service.
func DefaultConfigProvider() Config {
	env := os.Environ()
	cfg := Config{
		Endpoint: defaultAPIAddress,
	}

	for _, v := range env {
		d := strings.Split(v, "=")
		switch d[0] {
		case "PHPIPAM_APP_ID":
			cfg.AppID = d[1]
		case "PHPIPAM_ENDPOINT_ADDR":
			cfg.Endpoint = d[1]
		case "PHPIPAM_PASSWORD":
			cfg.Password = d[1]
		case "PHPIPAM_USER_NAME":
			cfg.Username = d[1]
		}
	}
	return cfg
}
