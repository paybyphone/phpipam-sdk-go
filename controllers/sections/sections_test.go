package sections

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paybyphone/phpipam-sdk-go/phpipam"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
)

var testListSectionsOutputExpected = []Section{
	Section{
		ID:          2,
		Name:        "IPv6",
		Description: "Section for IPv6 addresses",
		Permissions: "{\"3\":\"1\",\"2\":\"2\"}",
	},
	Section{
		ID:   3,
		Name: "foobar",
	},
	Section{
		ID:          1,
		Name:        "Customers",
		Description: "Section for customers",
		Permissions: "{\"3\":\"1\",\"2\":\"2\"}",
	},
}

const testListSectionsOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": [
    {
      "id": "2",
      "name": "IPv6",
      "description": "Section for IPv6 addresses",
      "masterSection": "0",
      "permissions": "{\"3\":\"1\",\"2\":\"2\"}",
      "strictMode": "0",
      "subnetOrdering": null,
      "order": null,
      "editDate": null,
      "showVLAN": "0",
      "showVRF": "0",
      "DNS": null,
      "links": [
        {
          "rel": "self",
          "href": "/api/test/sections/2/"
        }
      ]
    },
    {
      "id": "3",
      "name": "foobar",
      "description": null,
      "masterSection": "0",
      "permissions": null,
      "strictMode": "0",
      "subnetOrdering": null,
      "order": null,
      "editDate": null,
      "showVLAN": "0",
      "showVRF": "0",
      "DNS": null,
      "links": [
        {
          "rel": "self",
          "href": "/api/test/sections/3/"
        }
      ]
    },
    {
      "id": "1",
      "name": "Customers",
      "description": "Section for customers",
      "masterSection": "0",
      "permissions": "{\"3\":\"1\",\"2\":\"2\"}",
      "strictMode": "0",
      "subnetOrdering": null,
      "order": null,
      "editDate": null,
      "showVLAN": "0",
      "showVRF": "0",
      "DNS": null,
      "links": [
        {
          "rel": "self",
          "href": "/api/test/sections/1/"
        }
      ]
    }
  ]
}
`

var testCreateSectionInput = Section{
	Name: "foobar",
}

const testCreateSectionOutputExpected = `Section created`
const testCreateSectionOutputJSON = `
{
  "code": 201,
  "success": true,
  "data": "Section created"
}
`

var testGetSectionOutputExpected = Section{
	ID:          1,
	Name:        "Customers",
	Description: "Section for customers",
	Permissions: "{\"3\":\"1\",\"2\":\"2\"}",
}

const testGetSectionOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": {
    "id": "1",
    "name": "Customers",
    "description": "Section for customers",
    "masterSection": "0",
    "permissions": "{\"3\":\"1\",\"2\":\"2\"}",
    "strictMode": "0",
    "subnetOrdering": null,
    "order": null,
    "editDate": null,
    "showVLAN": "0",
    "showVRF": "0",
    "DNS": null,
    "links": [
      {
        "rel": "self",
        "href": "/api/test/sections/1/",
        "methods": [
          "GET",
          "POST",
          "DELETE",
          "PATCH"
        ]
      },
      {
        "rel": "subnets",
        "href": "/api/test/sections/1/subnets/",
        "methods": [
          "GET"
        ]
      }
    ]
  }
}
`

var testUpdateSectionInput = Section{
	ID:   3,
	Name: "foobaz",
}

const testUpdateSectionOutputJSON = `
{
  "code": 200,
  "success": true
}
`

var testDeleteSectionInput = Section{
	ID: 3,
}

const testDeleteSectionOutputJSON = `
{
  "code": 200,
  "success": true
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

func TestListSections(t *testing.T) {
	ts := httpOKTestServer(testListSectionsOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testListSectionsOutputExpected
	actual, err := client.ListSections()
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestCreateSection(t *testing.T) {
	ts := httpCreatedTestServer(testCreateSectionOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testCreateSectionInput
	expected := testCreateSectionOutputExpected
	actual, err := client.CreateSection(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetSectionByID(t *testing.T) {
	ts := httpOKTestServer(testGetSectionOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetSectionOutputExpected
	actual, err := client.GetSectionByID(1)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetSectionByName(t *testing.T) {
	ts := httpOKTestServer(testGetSectionOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetSectionOutputExpected
	actual, err := client.GetSectionByName("Customers")
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestUpdateSection(t *testing.T) {
	ts := httpOKTestServer(testUpdateSectionOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testUpdateSectionInput
	err := client.UpdateSection(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}
}
func TestDeleteSection(t *testing.T) {
	ts := httpOKTestServer(testUpdateSectionOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testDeleteSectionInput
	err := client.DeleteSection(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}
}
