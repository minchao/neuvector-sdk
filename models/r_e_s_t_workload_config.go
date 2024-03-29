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

// RESTWorkloadConfig r e s t workload config
//
// swagger:model RESTWorkloadConfig
type RESTWorkloadConfig struct {

	// quarantine
	// Example: true
	// Required: true
	Quarantine *bool `json:"quarantine"`

	// quarantine reason
	// Example: violation
	QuarantineReason string `json:"quarantine_reason,omitempty"`

	// wire
	// Example: default
	Wire string `json:"wire,omitempty"`
}

// Validate validates this r e s t workload config
func (m *RESTWorkloadConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuarantine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWorkloadConfig) validateQuarantine(formats strfmt.Registry) error {

	if err := validate.Required("quarantine", "body", m.Quarantine); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t workload config based on context it is used
func (m *RESTWorkloadConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTWorkloadConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTWorkloadConfig) UnmarshalBinary(b []byte) error {
	var res RESTWorkloadConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
