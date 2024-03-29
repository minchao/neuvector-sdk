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

// RESTSigstoreVerifierPatch r e s t sigstore verifier patch
//
// swagger:model RESTSigstoreVerifierPatch
type RESTSigstoreVerifierPatch struct {

	// cert issuer
	// Example: https://github.com/login/oauth
	CertIssuer string `json:"cert_issuer,omitempty"`

	// cert subject
	// Example: cert.subject@example.com
	CertSubject string `json:"cert_subject,omitempty"`

	// comment
	// Example: example comment
	Comment string `json:"comment,omitempty"`

	// public key
	// Example: -----BEGIN PUBLIC KEY-----XXXXXXXXXX-----END PUBLIC KEY-----
	PublicKey string `json:"public_key,omitempty"`

	// verifier type
	// Enum: [keypair keyless]
	VerifierType string `json:"verifier_type,omitempty"`
}

// Validate validates this r e s t sigstore verifier patch
func (m *RESTSigstoreVerifierPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerifierType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rESTSigstoreVerifierPatchTypeVerifierTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["keypair","keyless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTSigstoreVerifierPatchTypeVerifierTypePropEnum = append(rESTSigstoreVerifierPatchTypeVerifierTypePropEnum, v)
	}
}

const (

	// RESTSigstoreVerifierPatchVerifierTypeKeypair captures enum value "keypair"
	RESTSigstoreVerifierPatchVerifierTypeKeypair string = "keypair"

	// RESTSigstoreVerifierPatchVerifierTypeKeyless captures enum value "keyless"
	RESTSigstoreVerifierPatchVerifierTypeKeyless string = "keyless"
)

// prop value enum
func (m *RESTSigstoreVerifierPatch) validateVerifierTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTSigstoreVerifierPatchTypeVerifierTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTSigstoreVerifierPatch) validateVerifierType(formats strfmt.Registry) error {
	if swag.IsZero(m.VerifierType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVerifierTypeEnum("verifier_type", "body", m.VerifierType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t sigstore verifier patch based on context it is used
func (m *RESTSigstoreVerifierPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTSigstoreVerifierPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSigstoreVerifierPatch) UnmarshalBinary(b []byte) error {
	var res RESTSigstoreVerifierPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
