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

// RESTAdmCtrlRulesTestResult r e s t adm ctrl rules test result
//
// swagger:model RESTAdmCtrlRulesTestResult
type RESTAdmCtrlRulesTestResult struct {

	// allowed
	// Example: false
	// Required: true
	Allowed *bool `json:"allowed"`

	// index
	// Example: 1
	// Required: true
	Index *int64 `json:"index"`

	// kind
	// Example: Deployment
	// Required: true
	Kind *string `json:"kind"`

	// message
	// Example: Assessment Creation of Kubernetes Deployment resource (iperfserver) is denied because of deny rule id 1000 with criteria: (allow privilege escalation = true) [Notice: the requested image(s) are not scanned: quay.io/nvlab/iperf]
	// Required: true
	Message *string `json:"message"`

	// name
	// Example: iperfserver
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this r e s t adm ctrl rules test result
func (m *RESTAdmCtrlRulesTestResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdmCtrlRulesTestResult) validateAllowed(formats strfmt.Registry) error {

	if err := validate.Required("allowed", "body", m.Allowed); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmCtrlRulesTestResult) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmCtrlRulesTestResult) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmCtrlRulesTestResult) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmCtrlRulesTestResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t adm ctrl rules test result based on context it is used
func (m *RESTAdmCtrlRulesTestResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdmCtrlRulesTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdmCtrlRulesTestResult) UnmarshalBinary(b []byte) error {
	var res RESTAdmCtrlRulesTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
