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

// RESTSigstoreRootOfTrustPost r e s t sigstore root of trust post
//
// swagger:model RESTSigstoreRootOfTrustPost
type RESTSigstoreRootOfTrustPost struct {

	// comment
	// Example: example comment
	Comment string `json:"comment,omitempty"`

	// is private
	// Example: true
	IsPrivate bool `json:"is_private,omitempty"`

	// name
	// Example: example_name
	// Required: true
	Name *string `json:"name"`

	// rekor public key
	// Example: -----BEGIN PUBLIC KEY-----XXXXXXXXXX-----END PUBLIC KEY-----
	RekorPublicKey string `json:"rekor_public_key,omitempty"`

	// root cert
	// Example: -----BEGIN CERTIFICATE-----XXXXXXXXXX-----END CERTIFICATE-----
	RootCert string `json:"root_cert,omitempty"`

	// sct public key
	// Example: -----BEGIN PUBLIC KEY-----XXXXXXXXXX-----END PUBLIC KEY-----
	SctPublicKey string `json:"sct_public_key,omitempty"`
}

// Validate validates this r e s t sigstore root of trust post
func (m *RESTSigstoreRootOfTrustPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSigstoreRootOfTrustPost) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t sigstore root of trust post based on context it is used
func (m *RESTSigstoreRootOfTrustPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTSigstoreRootOfTrustPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSigstoreRootOfTrustPost) UnmarshalBinary(b []byte) error {
	var res RESTSigstoreRootOfTrustPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}