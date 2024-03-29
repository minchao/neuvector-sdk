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

// RESTServerSAMLConfig r e s t server s a m l config
//
// swagger:model RESTServerSAMLConfig
type RESTServerSAMLConfig struct {

	// default role
	// Example: admin
	DefaultRole string `json:"default_role,omitempty"`

	// enable
	// Example: true
	Enable bool `json:"enable,omitempty"`

	// group claim
	// Required: true
	GroupClaim *string `json:"group_claim"`

	// group mapped roles
	GroupMappedRoles []*GroupRoleMapping `json:"group_mapped_roles"`

	// issuer
	// Example: http://www.okta.com/xkbjKIvo0h
	// Required: true
	Issuer *string `json:"issuer"`

	// role groups
	RoleGroups *RESTServerSAMLConfigRoleGroups `json:"role_groups,omitempty"`

	// sso url
	// Example: https://dev-258.oktapreview.com/app/88_examplesamlapp_1/exYKIvqo0h7/sso/saml
	// Required: true
	SsoURL *string `json:"sso_url"`

	// x509 cert
	// Example: E7B0OS/N3KMVCL6KNMZ2+LOV90S7854NSD84P0BF
	X509Cert string `json:"x509_cert,omitempty"`

	// x509 cert extra
	// Example: ["E7B0OS/N3KMVCL6KNMZ2+LOV90S7854NSD84P0BF","E7B0OS/N3KMVCL6KNMZ2+LOV90S7854NSD84P0BF"]
	X509CertExtra []string `json:"x509_cert_extra"`
}

// Validate validates this r e s t server s a m l config
func (m *RESTServerSAMLConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupClaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMappedRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsoURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTServerSAMLConfig) validateGroupClaim(formats strfmt.Registry) error {

	if err := validate.Required("group_claim", "body", m.GroupClaim); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerSAMLConfig) validateGroupMappedRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupMappedRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupMappedRoles); i++ {
		if swag.IsZero(m.GroupMappedRoles[i]) { // not required
			continue
		}

		if m.GroupMappedRoles[i] != nil {
			if err := m.GroupMappedRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("group_mapped_roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("group_mapped_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RESTServerSAMLConfig) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerSAMLConfig) validateRoleGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleGroups) { // not required
		return nil
	}

	if m.RoleGroups != nil {
		if err := m.RoleGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_groups")
			}
			return err
		}
	}

	return nil
}

func (m *RESTServerSAMLConfig) validateSsoURL(formats strfmt.Registry) error {

	if err := validate.Required("sso_url", "body", m.SsoURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t server s a m l config based on the context it is used
func (m *RESTServerSAMLConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupMappedRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTServerSAMLConfig) contextValidateGroupMappedRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupMappedRoles); i++ {

		if m.GroupMappedRoles[i] != nil {

			if swag.IsZero(m.GroupMappedRoles[i]) { // not required
				return nil
			}

			if err := m.GroupMappedRoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("group_mapped_roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("group_mapped_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RESTServerSAMLConfig) contextValidateRoleGroups(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleGroups != nil {

		if swag.IsZero(m.RoleGroups) { // not required
			return nil
		}

		if err := m.RoleGroups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_groups")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTServerSAMLConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerSAMLConfig) UnmarshalBinary(b []byte) error {
	var res RESTServerSAMLConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTServerSAMLConfigRoleGroups r e s t server s a m l config role groups
//
// swagger:model RESTServerSAMLConfigRoleGroups
type RESTServerSAMLConfigRoleGroups struct {

	// groups
	// Example: ["admin1","admin2"]
	Groups []string `json:"groups"`

	// role
	// Example: admin
	Role string `json:"role,omitempty"`
}

// Validate validates this r e s t server s a m l config role groups
func (m *RESTServerSAMLConfigRoleGroups) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t server s a m l config role groups based on context it is used
func (m *RESTServerSAMLConfigRoleGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTServerSAMLConfigRoleGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerSAMLConfigRoleGroups) UnmarshalBinary(b []byte) error {
	var res RESTServerSAMLConfigRoleGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
