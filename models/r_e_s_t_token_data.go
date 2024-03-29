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

// RESTTokenData r e s t token data
//
// swagger:model RESTTokenData
type RESTTokenData struct {

	// password days until expire
	// Required: true
	PasswordDaysUntilExpire *int64 `json:"password_days_until_expire"`

	// password hours until expire
	// Required: true
	PasswordHoursUntilExpire *int64 `json:"password_hours_until_expire"`

	// token
	// Required: true
	Token *RESTToken `json:"token"`
}

// Validate validates this r e s t token data
func (m *RESTTokenData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePasswordDaysUntilExpire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordHoursUntilExpire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTTokenData) validatePasswordDaysUntilExpire(formats strfmt.Registry) error {

	if err := validate.Required("password_days_until_expire", "body", m.PasswordDaysUntilExpire); err != nil {
		return err
	}

	return nil
}

func (m *RESTTokenData) validatePasswordHoursUntilExpire(formats strfmt.Registry) error {

	if err := validate.Required("password_hours_until_expire", "body", m.PasswordHoursUntilExpire); err != nil {
		return err
	}

	return nil
}

func (m *RESTTokenData) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t token data based on the context it is used
func (m *RESTTokenData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTTokenData) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {

		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTTokenData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTTokenData) UnmarshalBinary(b []byte) error {
	var res RESTTokenData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
