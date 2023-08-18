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

// RESTApikeyCreation r e s t apikey creation
//
// swagger:model RESTApikeyCreation
type RESTApikeyCreation struct {

	// apikey name
	// Example: token-cicd-scan
	// Required: true
	ApikeyName *string `json:"apikey_name"`

	// description
	// Example: cicd image scanning
	Description string `json:"description,omitempty"`

	// expiration hours
	// Example: 1
	ExpirationHours uint32 `json:"expiration_hours,omitempty"`

	// expiration type
	// Example: never
	// Required: true
	ExpirationType *string `json:"expiration_type"`

	// role
	// Example: admin
	// Required: true
	Role *string `json:"role"`

	// role domains
	RoleDomains *RESTApikeyCreationRoleDomains `json:"role_domains,omitempty"`
}

// Validate validates this r e s t apikey creation
func (m *RESTApikeyCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApikeyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTApikeyCreation) validateApikeyName(formats strfmt.Registry) error {

	if err := validate.Required("apikey_name", "body", m.ApikeyName); err != nil {
		return err
	}

	return nil
}

func (m *RESTApikeyCreation) validateExpirationType(formats strfmt.Registry) error {

	if err := validate.Required("expiration_type", "body", m.ExpirationType); err != nil {
		return err
	}

	return nil
}

func (m *RESTApikeyCreation) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *RESTApikeyCreation) validateRoleDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleDomains) { // not required
		return nil
	}

	if m.RoleDomains != nil {
		if err := m.RoleDomains.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_domains")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_domains")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t apikey creation based on the context it is used
func (m *RESTApikeyCreation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTApikeyCreation) contextValidateRoleDomains(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleDomains != nil {

		if swag.IsZero(m.RoleDomains) { // not required
			return nil
		}

		if err := m.RoleDomains.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_domains")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_domains")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTApikeyCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTApikeyCreation) UnmarshalBinary(b []byte) error {
	var res RESTApikeyCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTApikeyCreationRoleDomains r e s t apikey creation role domains
//
// swagger:model RESTApikeyCreationRoleDomains
type RESTApikeyCreationRoleDomains struct {

	// domains
	Domains []string `json:"domains"`

	// role
	// Example: admin
	Role string `json:"role,omitempty"`
}

// Validate validates this r e s t apikey creation role domains
func (m *RESTApikeyCreationRoleDomains) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t apikey creation role domains based on context it is used
func (m *RESTApikeyCreationRoleDomains) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTApikeyCreationRoleDomains) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTApikeyCreationRoleDomains) UnmarshalBinary(b []byte) error {
	var res RESTApikeyCreationRoleDomains
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
