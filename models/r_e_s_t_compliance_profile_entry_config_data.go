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

// RESTComplianceProfileEntryConfigData r e s t compliance profile entry config data
//
// swagger:model RESTComplianceProfileEntryConfigData
type RESTComplianceProfileEntryConfigData struct {

	// config
	// Required: true
	Config *RESTComplianceProfileEntry `json:"config"`
}

// Validate validates this r e s t compliance profile entry config data
func (m *RESTComplianceProfileEntryConfigData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTComplianceProfileEntryConfigData) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t compliance profile entry config data based on the context it is used
func (m *RESTComplianceProfileEntryConfigData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTComplianceProfileEntryConfigData) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTComplianceProfileEntryConfigData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTComplianceProfileEntryConfigData) UnmarshalBinary(b []byte) error {
	var res RESTComplianceProfileEntryConfigData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
