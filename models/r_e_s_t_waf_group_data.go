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

// RESTWafGroupData r e s t waf group data
//
// swagger:model RESTWafGroupData
type RESTWafGroupData struct {

	// waf group
	// Required: true
	WafGroup *RESTWafGroup `json:"waf_group"`
}

// Validate validates this r e s t waf group data
func (m *RESTWafGroupData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWafGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWafGroupData) validateWafGroup(formats strfmt.Registry) error {

	if err := validate.Required("waf_group", "body", m.WafGroup); err != nil {
		return err
	}

	if m.WafGroup != nil {
		if err := m.WafGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("waf_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("waf_group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t waf group data based on the context it is used
func (m *RESTWafGroupData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWafGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWafGroupData) contextValidateWafGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.WafGroup != nil {

		if err := m.WafGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("waf_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("waf_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTWafGroupData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTWafGroupData) UnmarshalBinary(b []byte) error {
	var res RESTWafGroupData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
