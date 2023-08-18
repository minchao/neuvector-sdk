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

// RESTServerOIDCConfig r e s t server o ID c config
//
// swagger:model RESTServerOIDCConfig
type RESTServerOIDCConfig struct {

	// client id
	// Example: 0oai4gcl8xXh2itGi1h7
	// Required: true
	ClientID *string `json:"client_id"`

	// client secret
	// Example: QJju4mL1VLU0CAcD05WJ83K0D_e0gQEydowOvVqv
	ClientSecret string `json:"client_secret,omitempty"`

	// default role
	// Example: admin
	// Required: true
	DefaultRole *string `json:"default_role"`

	// enable
	// Example: true
	// Required: true
	Enable *bool `json:"enable"`

	// group claim
	// Required: true
	GroupClaim *string `json:"group_claim"`

	// group mapped roles
	GroupMappedRoles []*GroupRoleMapping `json:"group_mapped_roles"`

	// issuer
	// Example: https://dev-256438.oktapreview.com
	// Required: true
	Issuer *string `json:"issuer"`

	// role groups
	RoleGroups *RESTServerOIDCConfigRoleGroups `json:"role_groups,omitempty"`

	// scopes
	Scopes []string `json:"scopes"`
}

// Validate validates this r e s t server o ID c config
func (m *RESTServerOIDCConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnable(formats); err != nil {
		res = append(res, err)
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTServerOIDCConfig) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerOIDCConfig) validateDefaultRole(formats strfmt.Registry) error {

	if err := validate.Required("default_role", "body", m.DefaultRole); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerOIDCConfig) validateEnable(formats strfmt.Registry) error {

	if err := validate.Required("enable", "body", m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerOIDCConfig) validateGroupClaim(formats strfmt.Registry) error {

	if err := validate.Required("group_claim", "body", m.GroupClaim); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerOIDCConfig) validateGroupMappedRoles(formats strfmt.Registry) error {
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

func (m *RESTServerOIDCConfig) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerOIDCConfig) validateRoleGroups(formats strfmt.Registry) error {
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

// ContextValidate validate this r e s t server o ID c config based on the context it is used
func (m *RESTServerOIDCConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *RESTServerOIDCConfig) contextValidateGroupMappedRoles(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RESTServerOIDCConfig) contextValidateRoleGroups(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RESTServerOIDCConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerOIDCConfig) UnmarshalBinary(b []byte) error {
	var res RESTServerOIDCConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTServerOIDCConfigRoleGroups r e s t server o ID c config role groups
//
// swagger:model RESTServerOIDCConfigRoleGroups
type RESTServerOIDCConfigRoleGroups struct {

	// groups
	// Example: ["admin1","admin2"]
	Groups []string `json:"groups"`

	// role
	// Example: admin
	Role string `json:"role,omitempty"`
}

// Validate validates this r e s t server o ID c config role groups
func (m *RESTServerOIDCConfigRoleGroups) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t server o ID c config role groups based on context it is used
func (m *RESTServerOIDCConfigRoleGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTServerOIDCConfigRoleGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerOIDCConfigRoleGroups) UnmarshalBinary(b []byte) error {
	var res RESTServerOIDCConfigRoleGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
