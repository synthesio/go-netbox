// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Aggregate aggregate
// swagger:model Aggregate
type Aggregate struct {

	// Created
	// Read Only: true
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Date added
	DateAdded strfmt.Date `json:"date_added,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Family
	// Required: true
	Family *int64 `json:"family"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Prefix
	// Required: true
	Prefix *string `json:"prefix"`

	// rir
	// Required: true
	Rir *NestedRIR `json:"rir"`
}

// Validate validates this aggregate
func (m *Aggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFamily(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRir(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Aggregate) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

var aggregateTypeFamilyPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[4,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aggregateTypeFamilyPropEnum = append(aggregateTypeFamilyPropEnum, v)
	}
}

// prop value enum
func (m *Aggregate) validateFamilyEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, aggregateTypeFamilyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Aggregate) validateFamily(formats strfmt.Registry) error {

	if err := validate.Required("family", "body", m.Family); err != nil {
		return err
	}

	// value enum
	if err := m.validateFamilyEnum("family", "body", *m.Family); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) validateRir(formats strfmt.Registry) error {

	if err := validate.Required("rir", "body", m.Rir); err != nil {
		return err
	}

	if m.Rir != nil {

		if err := m.Rir.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rir")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Aggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Aggregate) UnmarshalBinary(b []byte) error {
	var res Aggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
