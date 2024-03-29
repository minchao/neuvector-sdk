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

// RESTServerRoleGroupsConfig r e s t server role groups config
//
// swagger:model RESTServerRoleGroupsConfig
type RESTServerRoleGroupsConfig struct {

	// groups
	// Example: ["reader1","reader2"]
	// Required: true
	Groups []string `json:"groups"`

	// name
	// Example: reader
	// Required: true
	Name *string `json:"name"`

	// role
	// Example: reader
	// Required: true
	Role *string `json:"role"`
}

// Validate validates this r e s t server role groups config
func (m *RESTServerRoleGroupsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTServerRoleGroupsConfig) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerRoleGroupsConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerRoleGroupsConfig) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t server role groups config based on context it is used
func (m *RESTServerRoleGroupsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTServerRoleGroupsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerRoleGroupsConfig) UnmarshalBinary(b []byte) error {
	var res RESTServerRoleGroupsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
