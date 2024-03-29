// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTLicenseKey r e s t license key
//
// swagger:model RESTLicenseKey
type RESTLicenseKey struct {

	// license key
	LicenseKey string `json:"license_key,omitempty"`
}

// Validate validates this r e s t license key
func (m *RESTLicenseKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t license key based on context it is used
func (m *RESTLicenseKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTLicenseKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTLicenseKey) UnmarshalBinary(b []byte) error {
	var res RESTLicenseKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
