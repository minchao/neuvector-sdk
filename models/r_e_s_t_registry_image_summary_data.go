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

// RESTRegistryImageSummaryData r e s t registry image summary data
//
// swagger:model RESTRegistryImageSummaryData
type RESTRegistryImageSummaryData struct {

	// images
	// Required: true
	Images []*RESTRegistryImageSummary `json:"images"`
}

// Validate validates this r e s t registry image summary data
func (m *RESTRegistryImageSummaryData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTRegistryImageSummaryData) validateImages(formats strfmt.Registry) error {

	if err := validate.Required("images", "body", m.Images); err != nil {
		return err
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t registry image summary data based on the context it is used
func (m *RESTRegistryImageSummaryData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTRegistryImageSummaryData) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {

			if swag.IsZero(m.Images[i]) { // not required
				return nil
			}

			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTRegistryImageSummaryData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTRegistryImageSummaryData) UnmarshalBinary(b []byte) error {
	var res RESTRegistryImageSummaryData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
