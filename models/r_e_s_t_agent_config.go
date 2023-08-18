// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTAgentConfig r e s t agent config
//
// swagger:model RESTAgentConfig
type RESTAgentConfig struct {

	// debug
	// Example: ["packet","log","parser","tcp","session","timer","error"]
	Debug []string `json:"debug"`

	// disable kvcctl
	// Example: false
	DisableKvcctl bool `json:"disable_kvcctl,omitempty"`

	// disable nvprotect
	// Example: false
	DisableNvprotect bool `json:"disable_nvprotect,omitempty"`
}

// Validate validates this r e s t agent config
func (m *RESTAgentConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t agent config based on context it is used
func (m *RESTAgentConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAgentConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAgentConfig) UnmarshalBinary(b []byte) error {
	var res RESTAgentConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
