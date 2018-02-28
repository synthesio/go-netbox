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

// Circuit circuit
// swagger:model Circuit
type Circuit struct {

	// Circuit ID
	// Required: true
	// Max Length: 50
	Cid *string `json:"cid"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Commit rate (Kbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	CommitRate *int64 `json:"commit_rate,omitempty"`

	// Created
	// Read Only: true
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Date installed
	InstallDate strfmt.Date `json:"install_date,omitempty"`

	// Last updated
	// Read Only: true
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// provider
	// Required: true
	Provider *NestedProvider `json:"provider"`

	// status
	// Required: true
	Status *CircuitStatus `json:"status"`

	// tenant
	// Required: true
	Tenant *NestedTenant `json:"tenant"`

	// type
	// Required: true
	Type *NestedCircuitType `json:"type"`
}

// Validate validates this circuit
func (m *Circuit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommitRate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Circuit) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	if err := validate.MaxLength("cid", "body", string(*m.Cid), 50); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateCommitRate(formats strfmt.Registry) error {

	if swag.IsZero(m.CommitRate) { // not required
		return nil
	}

	if err := validate.MinimumInt("commit_rate", "body", int64(*m.CommitRate), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("commit_rate", "body", int64(*m.CommitRate), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *Circuit) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	if m.Provider != nil {

		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	if m.Tenant != nil {

		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Circuit) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {

		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Circuit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Circuit) UnmarshalBinary(b []byte) error {
	var res Circuit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
