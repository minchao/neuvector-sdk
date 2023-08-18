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

// RESTRolePermission r e s t role permission
//
// swagger:model RESTRolePermission
type RESTRolePermission struct {

	// id
	// Example: ci_scan
	// Required: true
	ID *string `json:"id"`

	// read
	// Example: false
	// Required: true
	Read *bool `json:"read"`

	// write
	// Example: true
	// Required: true
	Write *bool `json:"write"`
}

// Validate validates this r e s t role permission
func (m *RESTRolePermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTRolePermission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RESTRolePermission) validateRead(formats strfmt.Registry) error {

	if err := validate.Required("read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *RESTRolePermission) validateWrite(formats strfmt.Registry) error {

	if err := validate.Required("write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t role permission based on context it is used
func (m *RESTRolePermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTRolePermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTRolePermission) UnmarshalBinary(b []byte) error {
	var res RESTRolePermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
