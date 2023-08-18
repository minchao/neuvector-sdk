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

// RESTServerLDAP r e s t server l d a p
//
// swagger:model RESTServerLDAP
type RESTServerLDAP struct {

	// base dn
	// Example: dc=example,dc=org
	// Required: true
	BaseDn *string `json:"base_dn"`

	// bind dn
	// Example: cn=admin,dc=example,dc=org
	// Required: true
	BindDn *string `json:"bind_dn"`

	// bind password
	// Example: mypassword
	BindPassword string `json:"bind_password,omitempty"`

	// default role
	// Example: reader
	// Required: true
	DefaultRole *string `json:"default_role"`

	// directory
	// Required: true
	Directory *string `json:"directory"`

	// enable
	// Example: false
	// Required: true
	Enable *bool `json:"enable"`

	// group mapped roles
	GroupMappedRoles []*GroupRoleMapping `json:"group_mapped_roles"`

	// group member attr
	// Required: true
	GroupMemberAttr *string `json:"group_member_attr"`

	// hostname
	// Example: 172.17.0.3
	// Required: true
	Hostname *string `json:"hostname"`

	// port
	// Example: 389
	// Required: true
	Port *uint16 `json:"port"`

	// role groups
	RoleGroups *RESTServerLDAPRoleGroups `json:"role_groups,omitempty"`

	// ssl
	// Example: false
	// Required: true
	Ssl *bool `json:"ssl"`

	// username attr
	// Required: true
	UsernameAttr *string `json:"username_attr"`
}

// Validate validates this r e s t server l d a p
func (m *RESTServerLDAP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseDn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindDn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMappedRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMemberAttr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsernameAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTServerLDAP) validateBaseDn(formats strfmt.Registry) error {

	if err := validate.Required("base_dn", "body", m.BaseDn); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateBindDn(formats strfmt.Registry) error {

	if err := validate.Required("bind_dn", "body", m.BindDn); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateDefaultRole(formats strfmt.Registry) error {

	if err := validate.Required("default_role", "body", m.DefaultRole); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateDirectory(formats strfmt.Registry) error {

	if err := validate.Required("directory", "body", m.Directory); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateEnable(formats strfmt.Registry) error {

	if err := validate.Required("enable", "body", m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateGroupMappedRoles(formats strfmt.Registry) error {
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

func (m *RESTServerLDAP) validateGroupMemberAttr(formats strfmt.Registry) error {

	if err := validate.Required("group_member_attr", "body", m.GroupMemberAttr); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateRoleGroups(formats strfmt.Registry) error {
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

func (m *RESTServerLDAP) validateSsl(formats strfmt.Registry) error {

	if err := validate.Required("ssl", "body", m.Ssl); err != nil {
		return err
	}

	return nil
}

func (m *RESTServerLDAP) validateUsernameAttr(formats strfmt.Registry) error {

	if err := validate.Required("username_attr", "body", m.UsernameAttr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t server l d a p based on the context it is used
func (m *RESTServerLDAP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *RESTServerLDAP) contextValidateGroupMappedRoles(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RESTServerLDAP) contextValidateRoleGroups(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RESTServerLDAP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerLDAP) UnmarshalBinary(b []byte) error {
	var res RESTServerLDAP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTServerLDAPRoleGroups r e s t server l d a p role groups
//
// swagger:model RESTServerLDAPRoleGroups
type RESTServerLDAPRoleGroups struct {

	// groups
	// Example: ["admin1","admin2"]
	Groups []string `json:"groups"`

	// role
	// Example: admin
	Role string `json:"role,omitempty"`
}

// Validate validates this r e s t server l d a p role groups
func (m *RESTServerLDAPRoleGroups) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t server l d a p role groups based on context it is used
func (m *RESTServerLDAPRoleGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTServerLDAPRoleGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTServerLDAPRoleGroups) UnmarshalBinary(b []byte) error {
	var res RESTServerLDAPRoleGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
