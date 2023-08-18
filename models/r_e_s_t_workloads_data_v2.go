// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RESTWorkloadsDataV2 r e s t workloads data v2
//
// swagger:model RESTWorkloadsDataV2
type RESTWorkloadsDataV2 struct {

	// workloads
	// Required: true
	Workloads []*RESTWorkloadV2 `json:"workloads"`
}

// Validate validates this r e s t workloads data v2
func (m *RESTWorkloadsDataV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkloads(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWorkloadsDataV2) validateWorkloads(formats strfmt.Registry) error {

	if err := validate.Required("workloads", "body", m.Workloads); err != nil {
		return err
	}

	for i := 0; i < len(m.Workloads); i++ {
		if swag.IsZero(m.Workloads[i]) { // not required
			continue
		}

		if m.Workloads[i] != nil {
			if err := m.Workloads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workloads" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workloads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t workloads data v2 based on the context it is used
func (m *RESTWorkloadsDataV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkloads(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWorkloadsDataV2) contextValidateWorkloads(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workloads); i++ {

		if m.Workloads[i] != nil {

			if swag.IsZero(m.Workloads[i]) { // not required
				return nil
			}

			if err := m.Workloads[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workloads" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workloads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTWorkloadsDataV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTWorkloadsDataV2) UnmarshalBinary(b []byte) error {
	var res RESTWorkloadsDataV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}