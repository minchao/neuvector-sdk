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

// RESTEULAData r e s t e u l a data
//
// swagger:model RESTEULAData
type RESTEULAData struct {

	// eula
	// Required: true
	Eula *RESTEULA `json:"eula"`
}

// Validate validates this r e s t e u l a data
func (m *RESTEULAData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEula(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTEULAData) validateEula(formats strfmt.Registry) error {

	if err := validate.Required("eula", "body", m.Eula); err != nil {
		return err
	}

	if m.Eula != nil {
		if err := m.Eula.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eula")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eula")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t e u l a data based on the context it is used
func (m *RESTEULAData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEula(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTEULAData) contextValidateEula(ctx context.Context, formats strfmt.Registry) error {

	if m.Eula != nil {

		if err := m.Eula.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eula")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eula")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTEULAData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTEULAData) UnmarshalBinary(b []byte) error {
	var res RESTEULAData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
