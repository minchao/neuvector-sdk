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

// RESTAwsLambdaRes r e s t aws lambda res
//
// swagger:model RESTAwsLambdaRes
type RESTAwsLambdaRes struct {

	// aws region resource
	// Required: true
	AwsRegionResource *RESTAwsLambdaResAwsRegionResource `json:"aws_region_resource"`

	// status
	// Example: data_lost
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this r e s t aws lambda res
func (m *RESTAwsLambdaRes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsRegionResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAwsLambdaRes) validateAwsRegionResource(formats strfmt.Registry) error {

	if err := validate.Required("aws_region_resource", "body", m.AwsRegionResource); err != nil {
		return err
	}

	if m.AwsRegionResource != nil {
		if err := m.AwsRegionResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_region_resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_region_resource")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAwsLambdaRes) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t aws lambda res based on the context it is used
func (m *RESTAwsLambdaRes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsRegionResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAwsLambdaRes) contextValidateAwsRegionResource(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsRegionResource != nil {

		if err := m.AwsRegionResource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_region_resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_region_resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAwsLambdaRes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAwsLambdaRes) UnmarshalBinary(b []byte) error {
	var res RESTAwsLambdaRes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTAwsLambdaResAwsRegionResource r e s t aws lambda res aws region resource
//
// swagger:model RESTAwsLambdaResAwsRegionResource
type RESTAwsLambdaResAwsRegionResource struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value *RESTAwsLambdaResDetail `json:"value,omitempty"`
}

// Validate validates this r e s t aws lambda res aws region resource
func (m *RESTAwsLambdaResAwsRegionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAwsLambdaResAwsRegionResource) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_region_resource" + "." + "value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_region_resource" + "." + "value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this r e s t aws lambda res aws region resource based on the context it is used
func (m *RESTAwsLambdaResAwsRegionResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAwsLambdaResAwsRegionResource) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_region_resource" + "." + "value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_region_resource" + "." + "value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAwsLambdaResAwsRegionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAwsLambdaResAwsRegionResource) UnmarshalBinary(b []byte) error {
	var res RESTAwsLambdaResAwsRegionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
