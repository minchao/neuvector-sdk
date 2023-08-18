// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTSystemConfigDataV2 r e s t system config data v2
//
// swagger:model RESTSystemConfigDataV2
type RESTSystemConfigDataV2 struct {

	// config
	Config *RESTSystemConfigV2 `json:"config,omitempty"`

	// fed config
	FedConfig *RESTFedSystemConfig `json:"fed_config,omitempty"`
}

// Validate validates this r e s t system config data v2
func (m *RESTSystemConfigDataV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFedConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfigDataV2) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigDataV2) validateFedConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.FedConfig) { // not required
		return nil
	}

	if m.FedConfig != nil {
		if err := m.FedConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fed_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fed_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t system config data v2 based on the context it is used
func (m *RESTSystemConfigDataV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFedConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfigDataV2) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigDataV2) contextValidateFedConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.FedConfig != nil {

		if swag.IsZero(m.FedConfig) { // not required
			return nil
		}

		if err := m.FedConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fed_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fed_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemConfigDataV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemConfigDataV2) UnmarshalBinary(b []byte) error {
	var res RESTSystemConfigDataV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}