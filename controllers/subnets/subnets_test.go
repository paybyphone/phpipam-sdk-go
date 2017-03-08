package subnets

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paybyphone/phpipam-sdk-go/phpipam"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
	"github.com/paybyphone/phpipam-sdk-go/testacc"
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
	client := NewController(sess)

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
	client := NewController(sess)

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
	client := NewController(sess)

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
	client := NewController(sess)

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
	client := NewController(sess)

	expected := testDeleteSubnetOutputExpected
	actual, err := client.DeleteSubnet(8)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}

// testAccSubnetCRUDCreate tests the creation part of the subnets controller
// CRUD acceptance test.
func testAccSubnetCRUDCreate(t *testing.T, s Subnet) {
	sess := session.NewSession()
	c := NewController(sess)

	if _, err := c.CreateSubnet(s); err != nil {
		t.Fatalf("Create: Error creating subnet: %s", err)
	}
}

// testAccSubnetCRUDReadByCIDR tests the read part of the subnets controller
// acceptance test, by fetching the subnet by CIDR. This is the first part of
// the 2-part read test, and also returns the ID of the subnet so that the
// test fixutre can be updated.
func testAccSubnetCRUDReadByCIDR(t *testing.T, s Subnet) int {
	sess := session.NewSession()
	c := NewController(sess)

	out, err := c.GetSubnetsByCIDR(fmt.Sprintf("%s/%d", s.SubnetAddress, s.Mask))
	if err != nil {
		t.Fatalf("Can't get subnet by CIDR: %s", err)
	}

	for _, v := range out {
		// We don't have an ID yet here, so set it.
		s.ID = v.ID
		if reflect.DeepEqual(s, v) {
			return v.ID
		}
	}

	t.Fatalf("ReadByCIDR: Could not find subnet %#v in %#v", s, out)
	return 0
}

// testAccSubnetCRUDReadByID tests the read part of the subnets controller
// acceptance test, by fetching the subnet by ID. This is the second part of
// the 2-part read test.
func testAccSubnetCRUDReadByID(t *testing.T, s Subnet) {
	sess := session.NewSession()
	c := NewController(sess)

	out, err := c.GetSubnetByID(s.ID)
	if err != nil {
		t.Fatalf("Can't find subnet by ID: %s", err)
	}

	if !reflect.DeepEqual(s, out) {
		t.Fatalf("ReadByID: Expected %#v, got %#v", s, out)
	}
}

// testAccSubnetCRUDUpdate tests the update part of the subnets controller
// acceptance test.
func testAccSubnetCRUDUpdate(t *testing.T, s Subnet) {
	sess := session.NewSession()
	c := NewController(sess)

	// Address or mask can't be in an update request.
	params := s
	params.SubnetAddress = ""
	params.Mask = 0

	if _, err := c.UpdateSubnet(params); err != nil {
		t.Fatalf("Error updating subnet: %s", err)
	}

	// Assert update
	out, err := c.GetSubnetByID(s.ID)

	if err != nil {
		t.Fatalf("Error fetching subnet after update: %s", err)
	}

	// Update updated date in original
	s.EditDate = out.EditDate

	if !reflect.DeepEqual(s, out) {
		t.Fatalf("Error after update: expected %#v, got %#v", s, out)
	}
}

// testAccSubnetCRUDDelete tests the delete part of the subnets controller
// acceptance test.
func testAccSubnetCRUDDelete(t *testing.T, s Subnet) {
	sess := session.NewSession()
	c := NewController(sess)

	if _, err := c.DeleteSubnet(s.ID); err != nil {
		t.Fatalf("Error deleting subnet: %s", err)
	}

	// check to see if subnet is actually gone
	if _, err := c.GetSubnetByID(s.ID); err == nil {
		t.Fatalf("Subnet still present after delete")
	}
}

// TestAccSubnetCRUD runs a full create-read-update-delete test for a PHPIPAM
// subnet.
func TestAccSubnetCRUD(t *testing.T) {
	testacc.VetAccConditions(t)

	subnet := testCreateSubnetInput
	testAccSubnetCRUDCreate(t, subnet)
	subnet.ID = testAccSubnetCRUDReadByCIDR(t, subnet)
	testAccSubnetCRUDReadByID(t, subnet)
	subnet.Description = "Updating subnet!"
	testAccSubnetCRUDUpdate(t, subnet)
	testAccSubnetCRUDDelete(t, subnet)
}
