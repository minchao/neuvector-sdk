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

// RESTProcessProfileData r e s t process profile data
//
// swagger:model RESTProcessProfileData
type RESTProcessProfileData struct {

	// process profile
	// Required: true
	ProcessProfile *RESTProcessProfile `json:"process_profile"`
}

// Validate validates this r e s t process profile data
func (m *RESTProcessProfileData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTProcessProfileData) validateProcessProfile(formats strfmt.Registry) error {

	if err := validate.Required("process_profile", "body", m.ProcessProfile); err != nil {
		return err
	}

	if m.ProcessProfile != nil {
		if err := m.ProcessProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("process_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("process_profile")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t process profile data based on the context it is used
func (m *RESTProcessProfileData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTProcessProfileData) contextValidateProcessProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessProfile != nil {

		if err := m.ProcessProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("process_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("process_profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTProcessProfileData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTProcessProfileData) UnmarshalBinary(b []byte) error {
	var res RESTProcessProfileData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
