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

// RESTSystemConfigProxyV2 r e s t system config proxy v2
//
// swagger:model RESTSystemConfigProxyV2
type RESTSystemConfigProxyV2 struct {

	// registry http proxy
	// Required: true
	RegistryHTTPProxy *RESTProxy `json:"registry_http_proxy"`

	// registry http proxy status
	// Example: true
	// Required: true
	RegistryHTTPProxyStatus *bool `json:"registry_http_proxy_status"`

	// registry https proxy
	// Required: true
	RegistryHTTPSProxy *RESTProxy `json:"registry_https_proxy"`

	// registry https proxy status
	// Example: false
	// Required: true
	RegistryHTTPSProxyStatus *bool `json:"registry_https_proxy_status"`
}

// Validate validates this r e s t system config proxy v2
func (m *RESTSystemConfigProxyV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistryHTTPProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryHTTPProxyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryHTTPSProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryHTTPSProxyStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfigProxyV2) validateRegistryHTTPProxy(formats strfmt.Registry) error {

	if err := validate.Required("registry_http_proxy", "body", m.RegistryHTTPProxy); err != nil {
		return err
	}

	if m.RegistryHTTPProxy != nil {
		if err := m.RegistryHTTPProxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_http_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_http_proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigProxyV2) validateRegistryHTTPProxyStatus(formats strfmt.Registry) error {

	if err := validate.Required("registry_http_proxy_status", "body", m.RegistryHTTPProxyStatus); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfigProxyV2) validateRegistryHTTPSProxy(formats strfmt.Registry) error {

	if err := validate.Required("registry_https_proxy", "body", m.RegistryHTTPSProxy); err != nil {
		return err
	}

	if m.RegistryHTTPSProxy != nil {
		if err := m.RegistryHTTPSProxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_https_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_https_proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigProxyV2) validateRegistryHTTPSProxyStatus(formats strfmt.Registry) error {

	if err := validate.Required("registry_https_proxy_status", "body", m.RegistryHTTPSProxyStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t system config proxy v2 based on the context it is used
func (m *RESTSystemConfigProxyV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistryHTTPProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistryHTTPSProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfigProxyV2) contextValidateRegistryHTTPProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistryHTTPProxy != nil {

		if err := m.RegistryHTTPProxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_http_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_http_proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigProxyV2) contextValidateRegistryHTTPSProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistryHTTPSProxy != nil {

		if err := m.RegistryHTTPSProxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_https_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_https_proxy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemConfigProxyV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemConfigProxyV2) UnmarshalBinary(b []byte) error {
	var res RESTSystemConfigProxyV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
