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

// RESTWafRule r e s t waf rule
//
// swagger:model RESTWafRule
type RESTWafRule struct {

	// cfg type
	// Required: true
	// Enum: [user_created ground]
	CfgType *string `json:"cfg_type"`

	// id
	// Example: 40003
	// Required: true
	ID *uint32 `json:"id"`

	// name
	// Example: test
	// Required: true
	Name *string `json:"name"`

	// patterns
	// Required: true
	Patterns []*RESTWafCriteriaEntry `json:"patterns"`
}

// Validate validates this r e s t waf rule
func (m *RESTWafRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfgType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatterns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rESTWafRuleTypeCfgTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user_created","ground"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTWafRuleTypeCfgTypePropEnum = append(rESTWafRuleTypeCfgTypePropEnum, v)
	}
}

const (

	// RESTWafRuleCfgTypeUserCreated captures enum value "user_created"
	RESTWafRuleCfgTypeUserCreated string = "user_created"

	// RESTWafRuleCfgTypeGround captures enum value "ground"
	RESTWafRuleCfgTypeGround string = "ground"
)

// prop value enum
func (m *RESTWafRule) validateCfgTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTWafRuleTypeCfgTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTWafRule) validateCfgType(formats strfmt.Registry) error {

	if err := validate.Required("cfg_type", "body", m.CfgType); err != nil {
		return err
	}

	// value enum
	if err := m.validateCfgTypeEnum("cfg_type", "body", *m.CfgType); err != nil {
		return err
	}

	return nil
}

func (m *RESTWafRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RESTWafRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RESTWafRule) validatePatterns(formats strfmt.Registry) error {

	if err := validate.Required("patterns", "body", m.Patterns); err != nil {
		return err
	}

	for i := 0; i < len(m.Patterns); i++ {
		if swag.IsZero(m.Patterns[i]) { // not required
			continue
		}

		if m.Patterns[i] != nil {
			if err := m.Patterns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t waf rule based on the context it is used
func (m *RESTWafRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatterns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWafRule) contextValidatePatterns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Patterns); i++ {

		if m.Patterns[i] != nil {

			if swag.IsZero(m.Patterns[i]) { // not required
				return nil
			}

			if err := m.Patterns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTWafRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTWafRule) UnmarshalBinary(b []byte) error {
	var res RESTWafRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
