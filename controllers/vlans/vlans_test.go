package vlans

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paybyphone/phpipam-sdk-go/phpipam"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
	"github.com/paybyphone/phpipam-sdk-go/testacc"
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

// testAccVLANCRUDCreate tests the creation part of the vlans controller
// CRUD acceptance test.
func testAccVLANCRUDCreate(t *testing.T, v VLAN) {
	sess := session.NewSession()
	c := New(sess)

	if _, err := c.CreateVLAN(v); err != nil {
		t.Fatalf("Create: Error creating vlan: %s", err)
	}
}

// testAccVLANCRUDReadByNumber tests the read part of the vlans controller
// acceptance test, by fetching the vlan by number. This is the first part of
// the 2-part read test, and also returns the ID of the vlan so that the
// test fixutre can be updated.
func testAccVLANCRUDReadByNumber(t *testing.T, v VLAN) int {
	sess := session.NewSession()
	c := New(sess)

	out, err := c.GetVLANsByNumber(v.Number)
	if err != nil {
		t.Fatalf("Can't get vlan by number: %s", err)
	}

	for _, val := range out {
		// We don't have an ID yet here, so set it.
		v.ID = val.ID
		if reflect.DeepEqual(v, val) {
			return val.ID
		}
	}

	t.Fatalf("ReadByNumber: Could not find vlan %#v in %#v", v, out)
	return 0
}

// testAccVLANCRUDReadByID tests the read part of the vlans controller
// acceptance test, by fetching the vlan by ID. This is the second part of
// the 2-part read test.
func testAccVLANCRUDReadByID(t *testing.T, v VLAN) {
	sess := session.NewSession()
	c := New(sess)

	out, err := c.GetVLANByID(v.ID)
	if err != nil {
		t.Fatalf("Can't find vlan by ID: %s", err)
	}

	if !reflect.DeepEqual(v, out) {
		t.Fatalf("ReadByID: Expected %#v, got %#v", v, out)
	}
}

// testAccVLANCRUDUpdate tests the update part of the vlans controller
// acceptance test.
func testAccVLANCRUDUpdate(t *testing.T, v VLAN) {
	sess := session.NewSession()
	c := New(sess)

	if _, err := c.UpdateVLAN(v); err != nil {
		t.Fatalf("Error updating vlan: %s", err)
	}

	// Assert update
	out, err := c.GetVLANByID(v.ID)

	if err != nil {
		t.Fatalf("Error fetching vlan after update: %s", err)
	}

	// Update updated date in original
	v.EditDate = out.EditDate

	if !reflect.DeepEqual(v, out) {
		t.Fatalf("Error after update: expected %#v, got %#v", v, out)
	}
}

// testAccVLANCRUDDelete tests the delete part of the vlans controller
// acceptance test.
func testAccVLANCRUDDelete(t *testing.T, v VLAN) {
	sess := session.NewSession()
	c := New(sess)

	if _, err := c.DeleteVLAN(v.ID); err != nil {
		t.Fatalf("Error deleting vlan: %s", err)
	}

	// check to see if vlan is actually gone
	if _, err := c.GetVLANByID(v.ID); err == nil {
		t.Fatalf("VLAN still present after delete")
	}
}

// TestAccVLANCRUD runs a full create-read-update-delete test for a PHPIPAM
// vlan.
func TestAccVLANCRUD(t *testing.T) {
	testacc.VetAccConditions(t)

	vlan := testCreateVLANInput
	testAccVLANCRUDCreate(t, vlan)
	// Add the domain ID here as 1 is the default.
	vlan.DomainID = 1
	vlan.ID = testAccVLANCRUDReadByNumber(t, vlan)
	testAccVLANCRUDReadByID(t, vlan)
	vlan.Name = "bazlan"
	testAccVLANCRUDUpdate(t, vlan)
	testAccVLANCRUDDelete(t, vlan)
}
