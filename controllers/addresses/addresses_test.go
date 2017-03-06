package addresses

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paybyphone/phpipam-sdk-go/phpipam"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
)

var testCreateAddressInput = Address{
	SubnetID:    3,
	IPAddress:   "10.10.1.10",
	Description: "foobar",
}

const testCreateAddressOutputExpected = `Address created`
const testCreateAddressOutputJSON = `
{
  "code": 201,
  "success": true,
  "data": "Address created"
}
`

var testGetAddressByIDOutputExpected = Address{
	ID:          11,
	SubnetID:    3,
	IPAddress:   "10.10.1.10",
	Description: "foobar",
	Tag:         2,
}

const testGetAddressByIDOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": {
    "id": "11",
    "subnetId": "3",
    "ip": "10.10.1.10",
    "is_gateway": null,
    "description": "foobar",
    "hostname": null,
    "mac": null,
    "owner": null,
    "tag": "2",
    "deviceId": null,
    "port": null,
    "note": null,
    "lastSeen": null,
    "excludePing": null,
    "PTRignore": null,
    "PTR": "0",
    "firewallAddressObject": null,
    "editDate": null,
    "links": [
      {
        "rel": "self",
        "href": "/api/test/addresses/11/",
        "methods": [
          "GET",
          "POST",
          "DELETE",
          "PATCH"
        ]
      },
      {
        "rel": "ping",
        "href": "/api/test/addresses/11/ping/",
        "methods": [
          "GET"
        ]
      }
    ]
  }
}
`

var testGetAddressesByIPOutputExpected = []Address{
	Address{
		ID:          11,
		SubnetID:    3,
		IPAddress:   "10.10.1.10",
		Description: "foobar",
		Tag:         2,
	},
}

const testGetAddressesByIPOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": [
    {
      "id": "11",
      "subnetId": "3",
      "ip": "10.10.1.10",
      "is_gateway": null,
      "description": "foobar",
      "hostname": null,
      "mac": null,
      "owner": null,
      "tag": "2",
      "deviceId": null,
      "port": null,
      "note": null,
      "lastSeen": null,
      "excludePing": null,
      "PTRignore": null,
      "PTR": "0",
      "firewallAddressObject": null,
      "editDate": null,
      "links": [
        {
          "rel": "self",
          "href": "/api/test/addresses/11/"
        }
      ]
    }
  ]
}
`

var testUpdateAddressInput = Address{
	ID:          11,
	Description: "bazboop",
}

const testUpdateAddressOutputExpected = `Address updated`
const testUpdateAddressOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Address updated"
}
`

const testDeleteAddressOutputExpected = `Address deleted`
const testDeleteAddressOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Address deleted"
}
`

func newHTTPTestServer(f func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(f))
	return ts
}

func httpOKTestServer(output string) *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, output, http.StatusOK)
	})
}

func httpCreatedTestServer(output string) *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, output, http.StatusCreated)
	})
}

func fullSessionConfig() *session.Session {
	return &session.Session{
		Config: phpipam.Config{
			AppID:    "0123456789abcdefgh",
			Password: "changeit",
			Username: "nobody",
		},
		Token: session.Token{
			String:  "foobarbazboop",
			Expires: "2999-12-31 23:59:59",
		},
	}
}

func TestCreateAddress(t *testing.T) {
	ts := httpCreatedTestServer(testCreateAddressOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testCreateAddressInput
	expected := testCreateAddressOutputExpected
	actual, err := client.CreateAddress(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetAddressByID(t *testing.T) {
	ts := httpOKTestServer(testGetAddressByIDOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetAddressByIDOutputExpected
	actual, err := client.GetAddressByID(11)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetAddressesByIP(t *testing.T) {
	ts := httpOKTestServer(testGetAddressesByIPOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetAddressesByIPOutputExpected
	actual, err := client.GetAddressesByIP("10.10.1.10/24")
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestUpdateAddress(t *testing.T) {
	ts := httpOKTestServer(testUpdateAddressOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testUpdateAddressInput
	expected := testUpdateAddressOutputExpected
	actual, err := client.UpdateAddress(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestDeleteAddress(t *testing.T) {
	ts := httpOKTestServer(testDeleteAddressOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testDeleteAddressOutputExpected
	actual, err := client.DeleteAddress(11, false)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}
