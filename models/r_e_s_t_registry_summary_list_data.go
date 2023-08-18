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

// RESTRegistrySummaryListData r e s t registry summary list data
//
// swagger:model RESTRegistrySummaryListData
type RESTRegistrySummaryListData struct {

	// summarys
	// Required: true
	Summarys []*RESTRegistrySummary `json:"summarys"`
}

// Validate validates this r e s t registry summary list data
func (m *RESTRegistrySummaryListData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummarys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTRegistrySummaryListData) validateSummarys(formats strfmt.Registry) error {

	if err := validate.Required("summarys", "body", m.Summarys); err != nil {
		return err
	}

	for i := 0; i < len(m.Summarys); i++ {
		if swag.IsZero(m.Summarys[i]) { // not required
			continue
		}

		if m.Summarys[i] != nil {
			if err := m.Summarys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summarys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summarys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t registry summary list data based on the context it is used
func (m *RESTRegistrySummaryListData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSummarys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTRegistrySummaryListData) contextValidateSummarys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Summarys); i++ {

		if m.Summarys[i] != nil {

			if swag.IsZero(m.Summarys[i]) { // not required
				return nil
			}

			if err := m.Summarys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summarys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summarys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTRegistrySummaryListData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTRegistrySummaryListData) UnmarshalBinary(b []byte) error {
	var res RESTRegistrySummaryListData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}