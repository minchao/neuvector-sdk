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

// RESTAuditsData r e s t audits data
//
// swagger:model RESTAuditsData
type RESTAuditsData struct {

	// audits
	// Required: true
	Audits []*Audit `json:"audits"`
}

// Validate validates this r e s t audits data
func (m *RESTAuditsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAuditsData) validateAudits(formats strfmt.Registry) error {

	if err := validate.Required("audits", "body", m.Audits); err != nil {
		return err
	}

	for i := 0; i < len(m.Audits); i++ {
		if swag.IsZero(m.Audits[i]) { // not required
			continue
		}

		if m.Audits[i] != nil {
			if err := m.Audits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("audits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("audits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t audits data based on the context it is used
func (m *RESTAuditsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAuditsData) contextValidateAudits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Audits); i++ {

		if m.Audits[i] != nil {

			if swag.IsZero(m.Audits[i]) { // not required
				return nil
			}

			if err := m.Audits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("audits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("audits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAuditsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAuditsData) UnmarshalBinary(b []byte) error {
	var res RESTAuditsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
