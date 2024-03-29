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

// RESTAwsFuncSummary r e s t aws func summary
//
// swagger:model RESTAwsFuncSummary
type RESTAwsFuncSummary struct {

	// function id
	// Example: project_id
	// Required: true
	FunctionID *string `json:"function_id"`

	// function name
	// Example: demo_project
	// Required: true
	FunctionName *string `json:"function_name"`

	// high
	// Example: 5
	// Required: true
	High *int64 `json:"high"`

	// medium
	// Example: 5
	// Required: true
	Medium *int64 `json:"medium"`

	// permission level
	// Required: true
	PermissionLevel *string `json:"permission_level"`

	// scan result
	// Required: true
	ScanResult *string `json:"scan_result"`

	// status
	// Example: data_lost
	// Required: true
	Status *string `json:"status"`

	// version
	// Example: 1.0
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this r e s t aws func summary
func (m *RESTAwsFuncSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunctionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHigh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissionLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAwsFuncSummary) validateFunctionID(formats strfmt.Registry) error {

	if err := validate.Required("function_id", "body", m.FunctionID); err != nil {
		return err
	}

	return nil
}

func (m *RESTAwsFuncSummary) validateFunctionName(formats strfmt.Registry) error {

	if err := validate.Required("function_name", "body", m.FunctionName); err != nil {
		return err
	}

	return nil
}

func (m *RESTAwsFuncSummary) validateHigh(formats strfmt.Registry) error {

	if err := validate.Required("high", "body", m.High); err != nil {
		return err
	}

	return nil
}

func (m *RESTAwsFuncSummary) validateMedium(formats strfmt.Registry) error {

	if err := validate.Required("medium", "body", m.Medium); err != nil {
		return err
	}

	return nil
}

func (m *RESTAwsFuncSummary) validatePermissionLevel(formats strfmt.Registry) error {

	if err := validate.Required("permission_level", "body", m.PermissionLevel); err != nil {
		return err
	}

	return nil
}

func (m *RESTAwsFuncSummary) validateScanResult(formats strfmt.Registry) error {

	if err := validate.Required("scan_result", "body", m.ScanResult); err != nil {
		return err
	}

	return nil
}

func (m *RESTAwsFuncSummary) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *RESTAwsFuncSummary) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t aws func summary based on context it is used
func (m *RESTAwsFuncSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAwsFuncSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAwsFuncSummary) UnmarshalBinary(b []byte) error {
	var res RESTAwsFuncSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
