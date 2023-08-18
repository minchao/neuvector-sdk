// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTSigstoreRootOfTrustPatch r e s t sigstore root of trust patch
//
// swagger:model RESTSigstoreRootOfTrustPatch
type RESTSigstoreRootOfTrustPatch struct {

	// comment
	// Example: example comment
	Comment string `json:"comment,omitempty"`

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

// Validate validates this r e s t sigstore root of trust patch
func (m *RESTSigstoreRootOfTrustPatch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t sigstore root of trust patch based on context it is used
func (m *RESTSigstoreRootOfTrustPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTSigstoreRootOfTrustPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSigstoreRootOfTrustPatch) UnmarshalBinary(b []byte) error {
	var res RESTSigstoreRootOfTrustPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
