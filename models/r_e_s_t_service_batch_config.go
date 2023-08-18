// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTServiceBatchConfig r e s t service batch config
//
// swagger:model RESTServiceBatchConfig
type RESTServiceBatchConfig struct {

	// baseline profile
	BaselineProfile string `json:"baseline_profile,omitempty"`

	// not scored
	// Example: false
	NotScored bool `json:"not_scored,omitempty"`

	// policy mode
	// Example: Monitor
	PolicyMode string `json:"policy_mode,omitempty"`

	// services
	// Example: ["iperfserver","iperfclient"]
	Services []string `json:"services"`
}

// Validate validates this r e s t service batch config
func (m *RESTServiceBatchConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t service batch config based on context it is used
func (m *RESTServiceBatchConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTServiceBatchConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServiceBatchConfig) UnmarshalBinary(b []byte) error {
	var res RESTServiceBatchConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}