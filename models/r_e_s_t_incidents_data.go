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

// RESTIncidentsData r e s t incidents data
//
// swagger:model RESTIncidentsData
type RESTIncidentsData struct {

	// incidents
	// Required: true
	Incidents []*Incident `json:"incidents"`
}

// Validate validates this r e s t incidents data
func (m *RESTIncidentsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTIncidentsData) validateIncidents(formats strfmt.Registry) error {

	if err := validate.Required("incidents", "body", m.Incidents); err != nil {
		return err
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t incidents data based on the context it is used
func (m *RESTIncidentsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIncidents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTIncidentsData) contextValidateIncidents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Incidents); i++ {

		if m.Incidents[i] != nil {

			if swag.IsZero(m.Incidents[i]) { // not required
				return nil
			}

			if err := m.Incidents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTIncidentsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTIncidentsData) UnmarshalBinary(b []byte) error {
	var res RESTIncidentsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
