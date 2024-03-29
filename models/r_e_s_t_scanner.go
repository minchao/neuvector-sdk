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

// RESTScanner r e s t scanner
//
// swagger:model RESTScanner
type RESTScanner struct {

	// cvedb create time
	// Example: 2018-06-20T19:00:53Z
	// Required: true
	// Format: date-time
	CvedbCreateTime *strfmt.DateTime `json:"cvedb_create_time"`

	// cvedb version
	// Example: 1.011
	// Required: true
	CvedbVersion *string `json:"cvedb_version"`

	// id
	// Example: github
	// Required: true
	ID *string `json:"id"`

	// port
	// Example: 51764
	// Required: true
	Port *uint16 `json:"port"`

	// server
	// Example: 10.1.5.1
	// Required: true
	Server *string `json:"server"`
}

// Validate validates this r e s t scanner
func (m *RESTScanner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCvedbCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCvedbVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTScanner) validateCvedbCreateTime(formats strfmt.Registry) error {

	if err := validate.Required("cvedb_create_time", "body", m.CvedbCreateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("cvedb_create_time", "body", "date-time", m.CvedbCreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanner) validateCvedbVersion(formats strfmt.Registry) error {

	if err := validate.Required("cvedb_version", "body", m.CvedbVersion); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanner) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanner) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanner) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t scanner based on context it is used
func (m *RESTScanner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTScanner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTScanner) UnmarshalBinary(b []byte) error {
	var res RESTScanner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
