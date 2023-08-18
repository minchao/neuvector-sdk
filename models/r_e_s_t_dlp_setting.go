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

// RESTDlpSetting r e s t dlp setting
//
// swagger:model RESTDlpSetting
type RESTDlpSetting struct {

	// action
	// Example: log
	// Required: true
	Action *string `json:"action"`

	// comment
	// Example: logging sensor ssn
	Comment string `json:"comment,omitempty"`

	// name
	// Example: sensor.ssn
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this r e s t dlp setting
func (m *RESTDlpSetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTDlpSetting) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *RESTDlpSetting) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t dlp setting based on context it is used
func (m *RESTDlpSetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTDlpSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTDlpSetting) UnmarshalBinary(b []byte) error {
	var res RESTDlpSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}