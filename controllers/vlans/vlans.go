// Package vlans provides types and methods for working with the VLAN
// controller.
package vlans

import (
	"fmt"

	"github.com/paybyphone/phpipam-sdk-go/phpipam/client"
	"github.com/paybyphone/phpipam-sdk-go/phpipam/session"
)

// VLAN represents a PHPIPAM VLAN.
type VLAN struct {
	// The VLAN ID. This is the entry ID in the PHPIPAM database, and not the
	// VLAN number, which is represented by the Number field.
	ID int `json:"id,string,omitempty"`

	// The Layer 2 domain identifier of the VLAN.
	DomainID int `json:"domainId,string,omitempty"`

	// The VLAN name/label.
	Name string `json:"name,omitempty"`

	// The VLAN number.
	Number int `json:"number,string,omitempty"`

	// A detailed description of the VLAN.
	Description string `json:"description,omitempty"`

	// The date of the last edit to this resource.
	EditDate string `json:"editDate,omitempty"`
}

// Controller is the base client for the VLAN controller.
type Controller struct {
	client.Client
}

// New returns a new instance of the client for the VLAN controller.
func New(sess *session.Session) *Controller {
	c := &Controller{
		Client: *client.NewClient(sess),
	}
	return c
}

// CreateVLAN creates a VLAN by sending a POST request.
func (c *Controller) CreateVLAN(in VLAN) (message string, err error) {
	err = c.SendRequest("POST", "/vlans/", &in, &message)
	return
}

// GetVLANByID GETs a VLAN via its ID in the PHPIPAM database.
func (c *Controller) GetVLANByID(id int) (out VLAN, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/vlans/%d/", id), &struct{}{}, &out)
	return
}

// GetVLANByNumber GETs a VLAN via its VLAN number.
func (c *Controller) GetVLANByNumber(id int) (out VLAN, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/vlans/search/%d/", id), &struct{}{}, &out)
	return
}

// UpdateVLAN updates a VLAN by sending a PATCH request.
func (c *Controller) UpdateVLAN(in VLAN) (message string, err error) {
	err = c.SendRequest("PATCH", "/vlans/", &in, &message)
	return
}

// DeleteVLAN deletes a VLAN by sending a DELETE request.
func (c *Controller) DeleteVLAN(in VLAN) (message string, err error) {
	err = c.SendRequest("DELETE", "/vlans/", &in, &message)
	return
}
