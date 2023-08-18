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

// RESTAdminCriteriaTemplate r e s t admin criteria template
//
// swagger:model RESTAdminCriteriaTemplate
type RESTAdminCriteriaTemplate struct {

	// kind
	// Example: podTemplate
	// Required: true
	Kind *string `json:"kind"`

	// rawjson
	// Example: {\"key\": \"value\"}
	// Required: true
	Rawjson *string `json:"rawjson"`
}

// Validate validates this r e s t admin criteria template
func (m *RESTAdminCriteriaTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRawjson(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdminCriteriaTemplate) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdminCriteriaTemplate) validateRawjson(formats strfmt.Registry) error {

	if err := validate.Required("rawjson", "body", m.Rawjson); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t admin criteria template based on context it is used
func (m *RESTAdminCriteriaTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdminCriteriaTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdminCriteriaTemplate) UnmarshalBinary(b []byte) error {
	var res RESTAdminCriteriaTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}