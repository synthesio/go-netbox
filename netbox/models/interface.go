// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Interface Interface
// swagger:model Interface
type Interface struct {

	// circuit termination
	// Required: true
	CircuitTermination *InterfaceCircuitTermination `json:"circuit_termination"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// form factor
	// Required: true
	FormFactor *InterfaceFormFactor `json:"form_factor"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Interface connection
	// Read Only: true
	InterfaceConnection string `json:"interface_connection,omitempty"`

	// Is connected
	// Read Only: true
	IsConnected string `json:"is_connected,omitempty"`

	// lag
	// Required: true
	Lag *NestedInterface `json:"lag"`

	// MAC Address
	MacAddress string `json:"mac_address,omitempty"`

	// OOB Management
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// mode
	// Required: true
	Mode *InterfaceMode `json:"mode"`

	// MTU
	// Maximum: 32767
	// Minimum: 0
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	Name *string `json:"name"`

	// tagged vlans
	// Required: true
	TaggedVlans InterfaceTaggedVlans `json:"tagged_vlans"`

	// untagged vlan
	// Required: true
	UntaggedVlan *InterfaceVLAN `json:"untagged_vlan"`
}

// Validate validates this interface
func (m *Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCircuitTermination(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFormFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUntaggedVlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interface) validateCircuitTermination(formats strfmt.Registry) error {

	if err := validate.Required("circuit_termination", "body", m.CircuitTermination); err != nil {
		return err
	}

	if m.CircuitTermination != nil {

		if err := m.CircuitTermination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("circuit_termination")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {

		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateFormFactor(formats strfmt.Registry) error {

	if err := validate.Required("form_factor", "body", m.FormFactor); err != nil {
		return err
	}

	if m.FormFactor != nil {

		if err := m.FormFactor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("form_factor")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateLag(formats strfmt.Registry) error {

	if err := validate.Required("lag", "body", m.Lag); err != nil {
		return err
	}

	if m.Lag != nil {

		if err := m.Lag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lag")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if m.Mode != nil {

		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateTaggedVlans(formats strfmt.Registry) error {

	if err := validate.Required("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	if err := m.TaggedVlans.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tagged_vlans")
		}
		return err
	}

	return nil
}

func (m *Interface) validateUntaggedVlan(formats strfmt.Registry) error {

	if err := validate.Required("untagged_vlan", "body", m.UntaggedVlan); err != nil {
		return err
	}

	if m.UntaggedVlan != nil {

		if err := m.UntaggedVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interface) UnmarshalBinary(b []byte) error {
	var res Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
