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

// RESTAwsResourceConfig r e s t aws resource config
//
// swagger:model RESTAwsResourceConfig
type RESTAwsResourceConfig struct {

	// acc id
	AccID string `json:"acc_id,omitempty"`

	// acc key
	AccKey string `json:"acc_key,omitempty"`

	// project name
	// Example: demo_project
	// Required: true
	ProjectName *string `json:"project_name"`

	// region list
	// Example: ["us-east-1","us-west-1","us-west-2"]
	RegionList []string `json:"region_list"`
}

// Validate validates this r e s t aws resource config
func (m *RESTAwsResourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAwsResourceConfig) validateProjectName(formats strfmt.Registry) error {

	if err := validate.Required("project_name", "body", m.ProjectName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t aws resource config based on context it is used
func (m *RESTAwsResourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAwsResourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAwsResourceConfig) UnmarshalBinary(b []byte) error {
	var res RESTAwsResourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
