package subnets

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paybyphone/phpipam-sdk-go/phpipam"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
)

var testCreateSubnetInput = Subnet{
	SubnetAddress:  "10.10.3.0",
	Mask:           24,
	SectionID:      1,
	MasterSubnetID: 2,
}

const testCreateSubnetOutputExpected = `Subnet created`
const testCreateSubnetOutputJSON = `
{
  "code": 201,
  "success": true,
  "data": "Subnet created"
}
`

var testGetSubnetByIDOutputExpected = Subnet{
	ID:             8,
	SubnetAddress:  "10.10.3.0",
	Mask:           24,
	SectionID:      1,
	MasterSubnetID: 2,
}

const testGetSubnetByIDOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": {
    "id": "8",
    "subnet": "10.10.3.0",
    "mask": "24",
    "sectionId": "1",
    "description": null,
    "firewallAddressObject": null,
    "vrfId": null,
    "masterSubnetId": "2",
    "allowRequests": "0",
    "vlanId": null,
    "showName": "0",
    "device": "0",
    "permissions": null,
    "pingSubnet": "0",
    "discoverSubnet": "0",
    "DNSrecursive": "0",
    "DNSrecords": "0",
    "nameserverId": "0",
    "scanAgent": null,
    "isFolder": "0",
    "isFull": "0",
    "tag": "2",
    "editDate": null,
    "links": [
      {
        "rel": "self",
        "href": "/api/test/subnets/8/",
        "methods": [
          "GET",
          "POST",
          "DELETE",
          "PATCH"
        ]
      },
      {
        "rel": "addresses",
        "href": "/api/test/subnets/8/addresses/",
        "methods": [
          "GET"
        ]
      },
      {
        "rel": "usage",
        "href": "/api/test/subnets/8/usage/",
        "methods": [
          "GET"
        ]
      },
      {
        "rel": "first_free",
        "href": "/api/test/subnets/8/first_free/",
        "methods": [
          "GET"
        ]
      },
      {
        "rel": "slaves",
        "href": "/api/test/subnets/8/slaves/",
        "methods": [
          "GET"
        ]
      },
      {
        "rel": "slaves_recursive",
        "href": "/api/test/subnets/8/slaves_recursive/",
        "methods": [
          "GET"
        ]
      },
      {
        "rel": "truncate",
        "href": "/api/test/subnets/8/truncate/",
        "methods": [
          "DELETE"
        ]
      },
      {
        "rel": "resize",
        "href": "/api/test/subnets/8/resize/",
        "methods": [
          "PATCH"
        ]
      },
      {
        "rel": "split",
        "href": "/api/test/subnets/8/split/",
        "methods": [
          "PATCH"
        ]
      }
    ]
  }
}
`

var testGetSubnetsByCIDROutputExpected = []Subnet{
	Subnet{
		ID:             8,
		SubnetAddress:  "10.10.3.0",
		Mask:           24,
		SectionID:      1,
		MasterSubnetID: 2,
	},
}

const testGetSubnetsByCIDROutputJSON = `
{
  "code": 200,
  "success": true,
  "data": [
    {
      "id": "8",
      "subnet": "10.10.3.0",
      "mask": "24",
      "sectionId": "1",
      "description": null,
      "firewallAddressObject": null,
      "vrfId": null,
      "masterSubnetId": "2",
      "allowRequests": "0",
      "vlanId": null,
      "showName": "0",
      "device": "0",
      "permissions": null,
      "pingSubnet": "0",
      "discoverSubnet": "0",
      "DNSrecursive": "0",
      "DNSrecords": "0",
      "nameserverId": "0",
      "scanAgent": null,
      "isFolder": "0",
      "isFull": "0",
      "tag": "2",
      "editDate": null,
      "links": [
        {
          "rel": "self",
          "href": "/api/test/subnets/8/"
        }
      ]
    }
  ]
}
`

var testUpdateSubnetInput = Subnet{
	ID:          8,
	Description: "foobat",
}

const testUpdateSubnetOutputExpected = `Subnet updated`
const testUpdateSubnetOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Subnet updated"
}
`

const testDeleteSubnetOutputExpected = `Subnet deleted`
const testDeleteSubnetOutputJSON = `
{
  "code": 200,
  "success": true,
  "data": "Subnet deleted"
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

func TestCreateSubnet(t *testing.T) {
	ts := httpCreatedTestServer(testCreateSubnetOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testCreateSubnetInput
	expected := testCreateSubnetOutputExpected
	actual, err := client.CreateSubnet(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetSubnetByID(t *testing.T) {
	ts := httpOKTestServer(testGetSubnetByIDOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetSubnetByIDOutputExpected
	actual, err := client.GetSubnetByID(8)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestGetSubnetsByCIDR(t *testing.T) {
	ts := httpOKTestServer(testGetSubnetsByCIDROutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testGetSubnetsByCIDROutputExpected
	actual, err := client.GetSubnetsByCIDR("10.10.3.0/24")
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestUpdateSubnet(t *testing.T) {
	ts := httpOKTestServer(testUpdateSubnetOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	in := testUpdateSubnetInput
	expected := testUpdateSubnetOutputExpected
	actual, err := client.UpdateSubnet(in)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

func TestDeleteSubnet(t *testing.T) {
	ts := httpOKTestServer(testDeleteSubnetOutputJSON)
	defer ts.Close()
	sess := fullSessionConfig()
	sess.Config.Endpoint = ts.URL
	client := New(sess)

	expected := testDeleteSubnetOutputExpected
	actual, err := client.DeleteSubnet(8)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}
