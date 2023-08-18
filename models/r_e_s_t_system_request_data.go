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

// RESTSystemRequestData r e s t system request data
//
// swagger:model RESTSystemRequestData
type RESTSystemRequestData struct {

	// request
	// Required: true
	Request *RESTSystemRequest `json:"request"`
}

// Validate validates this r e s t system request data
func (m *RESTSystemRequestData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemRequestData) validateRequest(formats strfmt.Registry) error {

	if err := validate.Required("request", "body", m.Request); err != nil {
		return err
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t system request data based on the context it is used
func (m *RESTSystemRequestData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemRequestData) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {

		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemRequestData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemRequestData) UnmarshalBinary(b []byte) error {
	var res RESTSystemRequestData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}