// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RESTSigstoreRootOfTrustGet r e s t sigstore root of trust get
//
// swagger:model RESTSigstoreRootOfTrustGet
type RESTSigstoreRootOfTrustGet struct {

	// cfg type
	// Enum: [user_created ground federal]
	CfgType string `json:"cfg_type,omitempty"`

	// comment
	// Example: example comment
	Comment string `json:"comment,omitempty"`

	// is private
	// Example: true
	IsPrivate bool `json:"is_private,omitempty"`

	// name
	// Example: example_name
	Name string `json:"name,omitempty"`

	// rekor public key
	// Example: -----BEGIN PUBLIC KEY-----XXXXXXXXXX-----END PUBLIC KEY-----
	RekorPublicKey string `json:"rekor_public_key,omitempty"`

	// root cert
	// Example: -----BEGIN CERTIFICATE-----XXXXXXXXXX-----END CERTIFICATE-----
	RootCert string `json:"root_cert,omitempty"`

	// sct public key
	// Example: -----BEGIN PUBLIC KEY-----XXXXXXXXXX-----END PUBLIC KEY-----
	SctPublicKey string `json:"sct_public_key,omitempty"`

	// verifiers
	Verifiers []*RESTSigstoreVerifier `json:"verifiers"`
}

// Validate validates this r e s t sigstore root of trust get
func (m *RESTSigstoreRootOfTrustGet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfgType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifiers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rESTSigstoreRootOfTrustGetTypeCfgTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user_created","ground","federal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTSigstoreRootOfTrustGetTypeCfgTypePropEnum = append(rESTSigstoreRootOfTrustGetTypeCfgTypePropEnum, v)
	}
}

const (

	// RESTSigstoreRootOfTrustGetCfgTypeUserCreated captures enum value "user_created"
	RESTSigstoreRootOfTrustGetCfgTypeUserCreated string = "user_created"

	// RESTSigstoreRootOfTrustGetCfgTypeGround captures enum value "ground"
	RESTSigstoreRootOfTrustGetCfgTypeGround string = "ground"

	// RESTSigstoreRootOfTrustGetCfgTypeFederal captures enum value "federal"
	RESTSigstoreRootOfTrustGetCfgTypeFederal string = "federal"
)

// prop value enum
func (m *RESTSigstoreRootOfTrustGet) validateCfgTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTSigstoreRootOfTrustGetTypeCfgTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTSigstoreRootOfTrustGet) validateCfgType(formats strfmt.Registry) error {
	if swag.IsZero(m.CfgType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCfgTypeEnum("cfg_type", "body", m.CfgType); err != nil {
		return err
	}

	return nil
}

func (m *RESTSigstoreRootOfTrustGet) validateVerifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.Verifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Verifiers); i++ {
		if swag.IsZero(m.Verifiers[i]) { // not required
			continue
		}

		if m.Verifiers[i] != nil {
			if err := m.Verifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("verifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("verifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t sigstore root of trust get based on the context it is used
func (m *RESTSigstoreRootOfTrustGet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVerifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSigstoreRootOfTrustGet) contextValidateVerifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Verifiers); i++ {

		if m.Verifiers[i] != nil {

			if swag.IsZero(m.Verifiers[i]) { // not required
				return nil
			}

			if err := m.Verifiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("verifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("verifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTSigstoreRootOfTrustGet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSigstoreRootOfTrustGet) UnmarshalBinary(b []byte) error {
	var res RESTSigstoreRootOfTrustGet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
