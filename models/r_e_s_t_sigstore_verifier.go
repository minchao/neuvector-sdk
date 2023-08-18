// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RESTSigstoreVerifier r e s t sigstore verifier
//
// swagger:model RESTSigstoreVerifier
type RESTSigstoreVerifier struct {

	// cert issuer
	// Example: https://github.com/login/oauth
	CertIssuer string `json:"cert_issuer,omitempty"`

	// cert subject
	// Example: cert.subject@example.com
	CertSubject string `json:"cert_subject,omitempty"`

	// comment
	// Example: example comment
	Comment string `json:"comment,omitempty"`

	// name
	// Example: example name
	// Required: true
	Name *string `json:"name"`

	// public key
	// Example: -----BEGIN PUBLIC KEY-----XXXXXXXXXX-----END PUBLIC KEY-----
	PublicKey string `json:"public_key,omitempty"`

	// verifier type
	// Required: true
	// Enum: [keypair keyless]
	VerifierType *string `json:"verifier_type"`
}

// Validate validates this r e s t sigstore verifier
func (m *RESTSigstoreVerifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifierType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSigstoreVerifier) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var rESTSigstoreVerifierTypeVerifierTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["keypair","keyless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTSigstoreVerifierTypeVerifierTypePropEnum = append(rESTSigstoreVerifierTypeVerifierTypePropEnum, v)
	}
}

const (

	// RESTSigstoreVerifierVerifierTypeKeypair captures enum value "keypair"
	RESTSigstoreVerifierVerifierTypeKeypair string = "keypair"

	// RESTSigstoreVerifierVerifierTypeKeyless captures enum value "keyless"
	RESTSigstoreVerifierVerifierTypeKeyless string = "keyless"
)

// prop value enum
func (m *RESTSigstoreVerifier) validateVerifierTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTSigstoreVerifierTypeVerifierTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTSigstoreVerifier) validateVerifierType(formats strfmt.Registry) error {

	if err := validate.Required("verifier_type", "body", m.VerifierType); err != nil {
		return err
	}

	// value enum
	if err := m.validateVerifierTypeEnum("verifier_type", "body", *m.VerifierType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t sigstore verifier based on context it is used
func (m *RESTSigstoreVerifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTSigstoreVerifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSigstoreVerifier) UnmarshalBinary(b []byte) error {
	var res RESTSigstoreVerifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
