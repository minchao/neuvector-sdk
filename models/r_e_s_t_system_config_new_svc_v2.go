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

// RESTSystemConfigNewSvcV2 r e s t system config new svc v2
//
// swagger:model RESTSystemConfigNewSvcV2
type RESTSystemConfigNewSvcV2 struct {

	// new service policy mode
	// Example: Discover
	// Required: true
	NewServicePolicyMode *string `json:"new_service_policy_mode"`

	// new service profile baseline
	// Example: zero-drift
	// Required: true
	NewServiceProfileBaseline *string `json:"new_service_profile_baseline"`
}

// Validate validates this r e s t system config new svc v2
func (m *RESTSystemConfigNewSvcV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewServicePolicyMode(formats); err != nil {
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

func (m *RESTSystemConfigNewSvcV2) validateNewServicePolicyMode(formats strfmt.Registry) error {

	if err := validate.Required("new_service_policy_mode", "body", m.NewServicePolicyMode); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfigNewSvcV2) validateNewServiceProfileBaseline(formats strfmt.Registry) error {

	if err := validate.Required("new_service_profile_baseline", "body", m.NewServiceProfileBaseline); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t system config new svc v2 based on context it is used
func (m *RESTSystemConfigNewSvcV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemConfigNewSvcV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemConfigNewSvcV2) UnmarshalBinary(b []byte) error {
	var res RESTSystemConfigNewSvcV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
