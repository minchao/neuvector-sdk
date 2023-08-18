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

// RESTProcessProfileEntryConfig r e s t process profile entry config
//
// swagger:model RESTProcessProfileEntryConfig
type RESTProcessProfileEntryConfig struct {

	// action
	// Example: monitor
	// Required: true
	Action *string `json:"action"`

	// group
	// Example: myGroup
	// Required: true
	Group *string `json:"group"`

	// name
	// Example: myEntryConfig
	// Required: true
	Name *string `json:"name"`

	// path
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this r e s t process profile entry config
func (m *RESTProcessProfileEntryConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTProcessProfileEntryConfig) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *RESTProcessProfileEntryConfig) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *RESTProcessProfileEntryConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RESTProcessProfileEntryConfig) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t process profile entry config based on context it is used
func (m *RESTProcessProfileEntryConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTProcessProfileEntryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTProcessProfileEntryConfig) UnmarshalBinary(b []byte) error {
	var res RESTProcessProfileEntryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}