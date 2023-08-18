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

// RESTSnifferArgsData r e s t sniffer args data
//
// swagger:model RESTSnifferArgsData
type RESTSnifferArgsData struct {

	// sniffer
	// Required: true
	Sniffer *RESTSnifferArgs `json:"sniffer"`
}

// Validate validates this r e s t sniffer args data
func (m *RESTSnifferArgsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSniffer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSnifferArgsData) validateSniffer(formats strfmt.Registry) error {

	if err := validate.Required("sniffer", "body", m.Sniffer); err != nil {
		return err
	}

	if m.Sniffer != nil {
		if err := m.Sniffer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sniffer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sniffer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t sniffer args data based on the context it is used
func (m *RESTSnifferArgsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSniffer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSnifferArgsData) contextValidateSniffer(ctx context.Context, formats strfmt.Registry) error {

	if m.Sniffer != nil {

		if err := m.Sniffer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sniffer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sniffer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTSnifferArgsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSnifferArgsData) UnmarshalBinary(b []byte) error {
	var res RESTSnifferArgsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}