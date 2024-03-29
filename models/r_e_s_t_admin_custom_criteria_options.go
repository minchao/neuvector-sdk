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

// RESTAdminCustomCriteriaOptions r e s t admin custom criteria options
//
// swagger:model RESTAdminCustomCriteriaOptions
type RESTAdminCustomCriteriaOptions struct {

	// ops
	// Example: ["exist","notExist"]
	// Required: true
	Ops []string `json:"ops"`

	// values
	// Example: ["true","false"]
	Values []string `json:"values"`

	// valuetype
	// Example: key
	// Required: true
	Valuetype *string `json:"valuetype"`
}

// Validate validates this r e s t admin custom criteria options
func (m *RESTAdminCustomCriteriaOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValuetype(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdminCustomCriteriaOptions) validateOps(formats strfmt.Registry) error {

	if err := validate.Required("ops", "body", m.Ops); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdminCustomCriteriaOptions) validateValuetype(formats strfmt.Registry) error {

	if err := validate.Required("valuetype", "body", m.Valuetype); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t admin custom criteria options based on context it is used
func (m *RESTAdminCustomCriteriaOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdminCustomCriteriaOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdminCustomCriteriaOptions) UnmarshalBinary(b []byte) error {
	var res RESTAdminCustomCriteriaOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
