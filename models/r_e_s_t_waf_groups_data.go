// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RESTWafGroupsData r e s t waf groups data
//
// swagger:model RESTWafGroupsData
type RESTWafGroupsData struct {

	// waf groups
	// Required: true
	WafGroups []*RESTWafGroup `json:"waf_groups"`
}

// Validate validates this r e s t waf groups data
func (m *RESTWafGroupsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWafGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWafGroupsData) validateWafGroups(formats strfmt.Registry) error {

	if err := validate.Required("waf_groups", "body", m.WafGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.WafGroups); i++ {
		if swag.IsZero(m.WafGroups[i]) { // not required
			continue
		}

		if m.WafGroups[i] != nil {
			if err := m.WafGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("waf_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("waf_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t waf groups data based on the context it is used
func (m *RESTWafGroupsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWafGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWafGroupsData) contextValidateWafGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WafGroups); i++ {

		if m.WafGroups[i] != nil {

			if swag.IsZero(m.WafGroups[i]) { // not required
				return nil
			}

			if err := m.WafGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("waf_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("waf_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTWafGroupsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTWafGroupsData) UnmarshalBinary(b []byte) error {
	var res RESTWafGroupsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
