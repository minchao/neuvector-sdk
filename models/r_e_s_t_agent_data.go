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

// RESTAgentData r e s t agent data
//
// swagger:model RESTAgentData
type RESTAgentData struct {

	// enforcer
	// Required: true
	Enforcer *RESTAgent `json:"enforcer"`
}

// Validate validates this r e s t agent data
func (m *RESTAgentData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnforcer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAgentData) validateEnforcer(formats strfmt.Registry) error {

	if err := validate.Required("enforcer", "body", m.Enforcer); err != nil {
		return err
	}

	if m.Enforcer != nil {
		if err := m.Enforcer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enforcer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enforcer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t agent data based on the context it is used
func (m *RESTAgentData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnforcer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAgentData) contextValidateEnforcer(ctx context.Context, formats strfmt.Registry) error {

	if m.Enforcer != nil {

		if err := m.Enforcer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enforcer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enforcer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAgentData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAgentData) UnmarshalBinary(b []byte) error {
	var res RESTAgentData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
