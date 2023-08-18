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

// RESTIDName r e s t ID name
//
// swagger:model RESTIDName
type RESTIDName struct {

	// display name
	// Required: true
	DisplayName *string `json:"display_name"`

	// domains
	// Example: ["domain1","domain2"]
	// Required: true
	Domains []string `json:"domains"`

	// id
	// Required: true
	ID *string `json:"id"`

	// policy mode
	// Required: true
	PolicyMode *string `json:"policy_mode"`
}

// Validate validates this r e s t ID name
func (m *RESTIDName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTIDName) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *RESTIDName) validateDomains(formats strfmt.Registry) error {

	if err := validate.Required("domains", "body", m.Domains); err != nil {
		return err
	}

	return nil
}

func (m *RESTIDName) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RESTIDName) validatePolicyMode(formats strfmt.Registry) error {

	if err := validate.Required("policy_mode", "body", m.PolicyMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t ID name based on context it is used
func (m *RESTIDName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTIDName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTIDName) UnmarshalBinary(b []byte) error {
	var res RESTIDName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
