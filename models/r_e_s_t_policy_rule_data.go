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

// RESTPolicyRuleData r e s t policy rule data
//
// swagger:model RESTPolicyRuleData
type RESTPolicyRuleData struct {

	// rule
	// Required: true
	Rule *RESTPolicyRule `json:"rule"`
}

// Validate validates this r e s t policy rule data
func (m *RESTPolicyRuleData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTPolicyRuleData) validateRule(formats strfmt.Registry) error {

	if err := validate.Required("rule", "body", m.Rule); err != nil {
		return err
	}

	if m.Rule != nil {
		if err := m.Rule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t policy rule data based on the context it is used
func (m *RESTPolicyRuleData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTPolicyRuleData) contextValidateRule(ctx context.Context, formats strfmt.Registry) error {

	if m.Rule != nil {

		if err := m.Rule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTPolicyRuleData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTPolicyRuleData) UnmarshalBinary(b []byte) error {
	var res RESTPolicyRuleData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
