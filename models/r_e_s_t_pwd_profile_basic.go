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

// RESTPwdProfileBasic r e s t pwd profile basic
//
// swagger:model RESTPwdProfileBasic
type RESTPwdProfileBasic struct {

	// min digit count
	// Example: 0
	// Required: true
	MinDigitCount *int64 `json:"min_digit_count"`

	// min len
	// Example: 6
	// Required: true
	MinLen *int64 `json:"min_len"`

	// min lowercase count
	// Example: 0
	// Required: true
	MinLowercaseCount *int64 `json:"min_lowercase_count"`

	// min special count
	// Example: 0
	// Required: true
	MinSpecialCount *int64 `json:"min_special_count"`

	// min uppercase count
	// Example: 0
	// Required: true
	MinUppercaseCount *int64 `json:"min_uppercase_count"`
}

// Validate validates this r e s t pwd profile basic
func (m *RESTPwdProfileBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMinDigitCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinLen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinLowercaseCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinSpecialCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinUppercaseCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTPwdProfileBasic) validateMinDigitCount(formats strfmt.Registry) error {

	if err := validate.Required("min_digit_count", "body", m.MinDigitCount); err != nil {
		return err
	}

	return nil
}

func (m *RESTPwdProfileBasic) validateMinLen(formats strfmt.Registry) error {

	if err := validate.Required("min_len", "body", m.MinLen); err != nil {
		return err
	}

	return nil
}

func (m *RESTPwdProfileBasic) validateMinLowercaseCount(formats strfmt.Registry) error {

	if err := validate.Required("min_lowercase_count", "body", m.MinLowercaseCount); err != nil {
		return err
	}

	return nil
}

func (m *RESTPwdProfileBasic) validateMinSpecialCount(formats strfmt.Registry) error {

	if err := validate.Required("min_special_count", "body", m.MinSpecialCount); err != nil {
		return err
	}

	return nil
}

func (m *RESTPwdProfileBasic) validateMinUppercaseCount(formats strfmt.Registry) error {

	if err := validate.Required("min_uppercase_count", "body", m.MinUppercaseCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t pwd profile basic based on context it is used
func (m *RESTPwdProfileBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTPwdProfileBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTPwdProfileBasic) UnmarshalBinary(b []byte) error {
	var res RESTPwdProfileBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
