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

// RESTSystemConfigNetSvcV2 r e s t system config net svc v2
//
// swagger:model RESTSystemConfigNetSvcV2
type RESTSystemConfigNetSvcV2 struct {

	// detect unmanaged wl
	// Example: true
	// Required: true
	DetectUnmanagedWl *bool `json:"detect_unmanaged_wl"`

	// disable net policy
	// Example: false
	// Required: true
	DisableNetPolicy *bool `json:"disable_net_policy"`

	// net service status
	// Example: true
	// Required: true
	NetServiceStatus *bool `json:"net_service_status"`

	// new service profile baseline
	// Example: zero-drift
	// Required: true
	NewServiceProfileBaseline *string `json:"new_service_profile_baseline"`
}

// Validate validates this r e s t system config net svc v2
func (m *RESTSystemConfigNetSvcV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetectUnmanagedWl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisableNetPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetServiceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewServiceProfileBaseline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfigNetSvcV2) validateDetectUnmanagedWl(formats strfmt.Registry) error {

	if err := validate.Required("detect_unmanaged_wl", "body", m.DetectUnmanagedWl); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfigNetSvcV2) validateDisableNetPolicy(formats strfmt.Registry) error {

	if err := validate.Required("disable_net_policy", "body", m.DisableNetPolicy); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfigNetSvcV2) validateNetServiceStatus(formats strfmt.Registry) error {

	if err := validate.Required("net_service_status", "body", m.NetServiceStatus); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfigNetSvcV2) validateNewServiceProfileBaseline(formats strfmt.Registry) error {

	if err := validate.Required("new_service_profile_baseline", "body", m.NewServiceProfileBaseline); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t system config net svc v2 based on context it is used
func (m *RESTSystemConfigNetSvcV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemConfigNetSvcV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemConfigNetSvcV2) UnmarshalBinary(b []byte) error {
	var res RESTSystemConfigNetSvcV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}