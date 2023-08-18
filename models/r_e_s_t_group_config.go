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

// RESTGroupConfig r e s t group config
//
// swagger:model RESTGroupConfig
type RESTGroupConfig struct {

	// cfg type
	// Required: true
	// Enum: [learned user_created ground federal]
	CfgType *string `json:"cfg_type"`

	// criteria
	Criteria []*RESTCriteriaEntry `json:"criteria"`

	// name
	// Example: containerEQU
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this r e s t group config
func (m *RESTGroupConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfgType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rESTGroupConfigTypeCfgTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["learned","user_created","ground","federal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTGroupConfigTypeCfgTypePropEnum = append(rESTGroupConfigTypeCfgTypePropEnum, v)
	}
}

const (

	// RESTGroupConfigCfgTypeLearned captures enum value "learned"
	RESTGroupConfigCfgTypeLearned string = "learned"

	// RESTGroupConfigCfgTypeUserCreated captures enum value "user_created"
	RESTGroupConfigCfgTypeUserCreated string = "user_created"

	// RESTGroupConfigCfgTypeGround captures enum value "ground"
	RESTGroupConfigCfgTypeGround string = "ground"

	// RESTGroupConfigCfgTypeFederal captures enum value "federal"
	RESTGroupConfigCfgTypeFederal string = "federal"
)

// prop value enum
func (m *RESTGroupConfig) validateCfgTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTGroupConfigTypeCfgTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTGroupConfig) validateCfgType(formats strfmt.Registry) error {

	if err := validate.Required("cfg_type", "body", m.CfgType); err != nil {
		return err
	}

	// value enum
	if err := m.validateCfgTypeEnum("cfg_type", "body", *m.CfgType); err != nil {
		return err
	}

	return nil
}

func (m *RESTGroupConfig) validateCriteria(formats strfmt.Registry) error {
	if swag.IsZero(m.Criteria) { // not required
		return nil
	}

	for i := 0; i < len(m.Criteria); i++ {
		if swag.IsZero(m.Criteria[i]) { // not required
			continue
		}

		if m.Criteria[i] != nil {
			if err := m.Criteria[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RESTGroupConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t group config based on the context it is used
func (m *RESTGroupConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTGroupConfig) contextValidateCriteria(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Criteria); i++ {

		if m.Criteria[i] != nil {

			if swag.IsZero(m.Criteria[i]) { // not required
				return nil
			}

			if err := m.Criteria[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTGroupConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTGroupConfig) UnmarshalBinary(b []byte) error {
	var res RESTGroupConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
