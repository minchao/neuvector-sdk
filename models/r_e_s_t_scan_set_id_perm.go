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

// RESTScanSetIDPerm r e s t scan set Id perm
//
// swagger:model RESTScanSetIdPerm
type RESTScanSetIDPerm struct {

	// evidence
	// Required: true
	Evidence *string `json:"evidence"`

	// path
	// Required: true
	Path *string `json:"path"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this r e s t scan set Id perm
func (m *RESTScanSetIDPerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
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

func (m *RESTScanSetIDPerm) validateEvidence(formats strfmt.Registry) error {

	if err := validate.Required("evidence", "body", m.Evidence); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanSetIDPerm) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanSetIDPerm) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t scan set Id perm based on context it is used
func (m *RESTScanSetIDPerm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTScanSetIDPerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTScanSetIDPerm) UnmarshalBinary(b []byte) error {
	var res RESTScanSetIDPerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}