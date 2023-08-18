// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RESTBenchItem r e s t bench item
//
// swagger:model RESTBenchItem
type RESTBenchItem struct {

	// automated
	// Example: true
	// Required: true
	Automated *bool `json:"automated"`

	// catalog
	// Example: docker
	// Required: true
	Catalog *string `json:"catalog"`

	// description
	// Example: General Configuration
	// Required: true
	Description *string `json:"description"`

	// group
	// Example: nv.calico
	// Required: true
	Group *string `json:"group"`

	// level
	// Example: INFO
	// Required: true
	Level *string `json:"level"`

	// message
	// Required: true
	Message []string `json:"message"`

	// profile
	// Example: Level 1
	// Required: true
	Profile *string `json:"profile"`

	// remediation
	// Required: true
	Remediation *string `json:"remediation"`

	// scored
	// Example: true
	// Required: true
	Scored *bool `json:"scored"`

	// test number
	// Example: 1
	// Required: true
	TestNumber *string `json:"test_number"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this r e s t bench item
func (m *RESTBenchItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutomated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemediation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScored(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTBenchItem) validateAutomated(formats strfmt.Registry) error {

	if err := validate.Required("automated", "body", m.Automated); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateCatalog(formats strfmt.Registry) error {

	if err := validate.Required("catalog", "body", m.Catalog); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("profile", "body", m.Profile); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateRemediation(formats strfmt.Registry) error {

	if err := validate.Required("remediation", "body", m.Remediation); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateScored(formats strfmt.Registry) error {

	if err := validate.Required("scored", "body", m.Scored); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateTestNumber(formats strfmt.Registry) error {

	if err := validate.Required("test_number", "body", m.TestNumber); err != nil {
		return err
	}

	return nil
}

func (m *RESTBenchItem) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t bench item based on context it is used
func (m *RESTBenchItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTBenchItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTBenchItem) UnmarshalBinary(b []byte) error {
	var res RESTBenchItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}