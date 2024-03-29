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

// RESTDlpSensorsData r e s t dlp sensors data
//
// swagger:model RESTDlpSensorsData
type RESTDlpSensorsData struct {

	// sensors
	// Required: true
	Sensors []*RESTDlpSensor `json:"sensors"`
}

// Validate validates this r e s t dlp sensors data
func (m *RESTDlpSensorsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSensors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTDlpSensorsData) validateSensors(formats strfmt.Registry) error {

	if err := validate.Required("sensors", "body", m.Sensors); err != nil {
		return err
	}

	for i := 0; i < len(m.Sensors); i++ {
		if swag.IsZero(m.Sensors[i]) { // not required
			continue
		}

		if m.Sensors[i] != nil {
			if err := m.Sensors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sensors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sensors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t dlp sensors data based on the context it is used
func (m *RESTDlpSensorsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSensors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTDlpSensorsData) contextValidateSensors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sensors); i++ {

		if m.Sensors[i] != nil {

			if swag.IsZero(m.Sensors[i]) { // not required
				return nil
			}

			if err := m.Sensors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sensors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sensors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTDlpSensorsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTDlpSensorsData) UnmarshalBinary(b []byte) error {
	var res RESTDlpSensorsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
