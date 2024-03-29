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

// RESTAdmissionState r e s t admission state
//
// swagger:model RESTAdmissionState
type RESTAdmissionState struct {

	// adm client mode
	// Example: service
	AdmClientMode string `json:"adm_client_mode,omitempty"`

	// adm client mode options
	AdmClientModeOptions *RESTAdmissionStateAdmClientModeOptions `json:"adm_client_mode_options,omitempty"`

	// adm svc type
	AdmSvcType string `json:"adm_svc_type,omitempty"`

	// ctrl states
	CtrlStates *RESTAdmissionStateCtrlStates `json:"ctrl_states,omitempty"`

	// default action
	// Example: allow
	DefaultAction string `json:"default_action,omitempty"`

	// enable
	// Example: true
	Enable bool `json:"enable,omitempty"`

	// mode
	// Example: Protect
	Mode string `json:"mode,omitempty"`
}

// Validate validates this r e s t admission state
func (m *RESTAdmissionState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmClientModeOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCtrlStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdmissionState) validateAdmClientModeOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.AdmClientModeOptions) { // not required
		return nil
	}

	if m.AdmClientModeOptions != nil {
		if err := m.AdmClientModeOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adm_client_mode_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adm_client_mode_options")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAdmissionState) validateCtrlStates(formats strfmt.Registry) error {
	if swag.IsZero(m.CtrlStates) { // not required
		return nil
	}

	if m.CtrlStates != nil {
		if err := m.CtrlStates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ctrl_states")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ctrl_states")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t admission state based on the context it is used
func (m *RESTAdmissionState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdmClientModeOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCtrlStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdmissionState) contextValidateAdmClientModeOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.AdmClientModeOptions != nil {

		if swag.IsZero(m.AdmClientModeOptions) { // not required
			return nil
		}

		if err := m.AdmClientModeOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adm_client_mode_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adm_client_mode_options")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAdmissionState) contextValidateCtrlStates(ctx context.Context, formats strfmt.Registry) error {

	if m.CtrlStates != nil {

		if swag.IsZero(m.CtrlStates) { // not required
			return nil
		}

		if err := m.CtrlStates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ctrl_states")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ctrl_states")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdmissionState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdmissionState) UnmarshalBinary(b []byte) error {
	var res RESTAdmissionState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTAdmissionStateAdmClientModeOptions r e s t admission state adm client mode options
//
// swagger:model RESTAdmissionStateAdmClientModeOptions
type RESTAdmissionStateAdmClientModeOptions struct {

	// service
	// Example: service
	Service string `json:"service,omitempty"`

	// url
	// Example: service:xyz-svc-admission-webhook.xyz.svc
	URL string `json:"url,omitempty"`
}

// Validate validates this r e s t admission state adm client mode options
func (m *RESTAdmissionStateAdmClientModeOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t admission state adm client mode options based on context it is used
func (m *RESTAdmissionStateAdmClientModeOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdmissionStateAdmClientModeOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdmissionStateAdmClientModeOptions) UnmarshalBinary(b []byte) error {
	var res RESTAdmissionStateAdmClientModeOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTAdmissionStateCtrlStates r e s t admission state ctrl states
//
// swagger:model RESTAdmissionStateCtrlStates
type RESTAdmissionStateCtrlStates struct {

	// validate type
	// Example: validate
	ValidateType string `json:"validate,omitempty"`

	// states
	// Example: true
	States bool `json:"states,omitempty"`
}

// Validate validates this r e s t admission state ctrl states
func (m *RESTAdmissionStateCtrlStates) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t admission state ctrl states based on context it is used
func (m *RESTAdmissionStateCtrlStates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdmissionStateCtrlStates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdmissionStateCtrlStates) UnmarshalBinary(b []byte) error {
	var res RESTAdmissionStateCtrlStates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
