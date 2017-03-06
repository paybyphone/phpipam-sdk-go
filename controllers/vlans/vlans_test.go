package vlans

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paybyphone/phpipam-sdk-go/phpipam"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
)

var testCreateVLANInput = VLAN{
	Name:   "foolan",
	Number: 1000,
}

const testCreateVLANOutputExpected = `Vlan created`
const testCreateVLANOutputJSON = `
{
  "code": 201,
  "success": true,
  "data": "Vlan created"
}
`

var testGetVLANByIDOutputExpected = VLAN{
	ID:       3,
	DomainID: 1,
	Name:     "foolan",
	Number:   1000,
}

const testGetVLANByIDOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": {
    "id": "3",
    "domainId": "1",
    "name": "foolan",
    "number": "1000",
    "description": null,
    "editDate": null,
    "links": [
      {
        "rel": "self",
        "href": "/api/test/vlans/3/",
        "methods": [
          "GET",
          "POST",
          "DELETE",
          "PATCH"
        ]
      },
      {
        "rel": "subnets",
        "href": "/api/test/vlans/3/subnets/",
        "methods": [
          "GET"
        ]
      }
    ]
  }
}
`

var testGetVLANsByNumberOutputExpected = []VLAN{
	VLAN{
		ID:       3,
		DomainID: 1,
		Name:     "foolan",
		Number:   1000,
	},
}

const testGetVLANsByNumberOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": [
    {
      "id": "3",
      "domainId": "1",
      "name": "foolan",
      "number": "1000",
      "description": null,
      "editDate": null,
      "links": [
        {
          "rel": "self",
          "href": "/api/test/vlans/3/"
        }
      ]
    }
  ]
}
`

var testUpdateVLANInput = VLAN{
	ID:   3,
	Name: "bazlan",
}

const testUpdateVLANOutputExpected = `Vlan updated`
const testUpdateVLANOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Vlan updated"
}
`

const testDeleteVLANOutputExpected = `Vlan deleted`
const testDeleteVLANOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Vlan deleted"
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

func TestCreateVLAN(t *testing.T) {
	ts := httpCreatedTestServer(testCreateVLANOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testCreateVLANInput
	expected := testCreateVLANOutputExpected
	actual, err := client.CreateVLAN(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetVLANByID(t *testing.T) {
	ts := httpOKTestServer(testGetVLANByIDOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetVLANByIDOutputExpected
	actual, err := client.GetVLANByID(3)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetVLANsByNumber(t *testing.T) {
	ts := httpOKTestServer(testGetVLANsByNumberOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetVLANsByNumberOutputExpected
	actual, err := client.GetVLANsByNumber(1000)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestUpdateVLAN(t *testing.T) {
	ts := httpOKTestServer(testUpdateVLANOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testUpdateVLANInput
	expected := testUpdateVLANOutputExpected
	actual, err := client.UpdateVLAN(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestDeleteVLAN(t *testing.T) {
	ts := httpOKTestServer(testDeleteVLANOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testDeleteVLANOutputExpected
	actual, err := client.DeleteVLAN(3)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}
