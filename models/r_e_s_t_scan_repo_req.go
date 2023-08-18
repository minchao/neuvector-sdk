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

// RESTScanRepoReq r e s t scan repo req
//
// swagger:model RESTScanRepoReq
type RESTScanRepoReq struct {

	// base image
	// Example: alpine
	// Required: true
	BaseImage *string `json:"base_image"`

	// metadata
	// Required: true
	Metadata *RESTScanMeta `json:"metadata"`

	// password
	// Example: mypassword
	// Format: password
	Password strfmt.Password `json:"password,omitempty"`

	// registry
	// Example: https://registry.hub.docker.com/
	// Required: true
	Registry *string `json:"registry"`

	// repository
	// Example: alpine
	// Required: true
	Repository *string `json:"repository"`

	// scan layers
	// Example: false
	// Required: true
	ScanLayers *bool `json:"scan_layers"`

	// tag
	// Example: latest
	// Required: true
	Tag *string `json:"tag"`

	// username
	// Example: myusername
	Username string `json:"username,omitempty"`
}

// Validate validates this r e s t scan repo req
func (m *RESTScanRepoReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanLayers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTScanRepoReq) validateBaseImage(formats strfmt.Registry) error {

	if err := validate.Required("base_image", "body", m.BaseImage); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanRepoReq) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *RESTScanRepoReq) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanRepoReq) validateRegistry(formats strfmt.Registry) error {

	if err := validate.Required("registry", "body", m.Registry); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanRepoReq) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanRepoReq) validateScanLayers(formats strfmt.Registry) error {

	if err := validate.Required("scan_layers", "body", m.ScanLayers); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanRepoReq) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t scan repo req based on the context it is used
func (m *RESTScanRepoReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTScanRepoReq) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTScanRepoReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTScanRepoReq) UnmarshalBinary(b []byte) error {
	var res RESTScanRepoReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}