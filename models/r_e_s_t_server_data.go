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

// RESTServerData r e s t server data
//
// swagger:model RESTServerData
type RESTServerData struct {

	// server
	// Required: true
	Server *RESTServer `json:"server"`
}

// Validate validates this r e s t server data
func (m *RESTServerData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTServerData) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t server data based on the context it is used
func (m *RESTServerData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTServerData) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {

		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTServerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerData) UnmarshalBinary(b []byte) error {
	var res RESTServerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
