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

// RESTWafGroup r e s t waf group
//
// swagger:model RESTWafGroup
type RESTWafGroup struct {

	// cfg type
	// Required: true
	// Enum: [user_created ground]
	CfgType *string `json:"cfg_type"`

	// name
	// Example: nv.kube-proxy.kube-system
	// Required: true
	Name *string `json:"name"`

	// sensors
	// Required: true
	Sensors []*RESTWafSetting `json:"sensors"`

	// status
	// Example: true
	// Required: true
	Status *bool `json:"status"`
}

// Validate validates this r e s t waf group
func (m *RESTWafGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfgType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rESTWafGroupTypeCfgTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user_created","ground"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTWafGroupTypeCfgTypePropEnum = append(rESTWafGroupTypeCfgTypePropEnum, v)
	}
}

const (

	// RESTWafGroupCfgTypeUserCreated captures enum value "user_created"
	RESTWafGroupCfgTypeUserCreated string = "user_created"

	// RESTWafGroupCfgTypeGround captures enum value "ground"
	RESTWafGroupCfgTypeGround string = "ground"
)

// prop value enum
func (m *RESTWafGroup) validateCfgTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTWafGroupTypeCfgTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTWafGroup) validateCfgType(formats strfmt.Registry) error {

	if err := validate.Required("cfg_type", "body", m.CfgType); err != nil {
		return err
	}

	// value enum
	if err := m.validateCfgTypeEnum("cfg_type", "body", *m.CfgType); err != nil {
		return err
	}

	return nil
}

func (m *RESTWafGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RESTWafGroup) validateSensors(formats strfmt.Registry) error {

	if err := validate.Required("sensors", "body", m.Sensors); err != nil {
		return err
	}

	for i := 0; i < len(m.Sensors); i++ {
		if swag.IsZero(m.Sensors[i]) { // not required
			continue
		}

		if m.Sensors[i] != nil {
			if err := m.Sensors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sensors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sensors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RESTWafGroup) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t waf group based on the context it is used
func (m *RESTWafGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSensors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTWafGroup) contextValidateSensors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sensors); i++ {

		if m.Sensors[i] != nil {

			if swag.IsZero(m.Sensors[i]) { // not required
				return nil
			}

			if err := m.Sensors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sensors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sensors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTWafGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTWafGroup) UnmarshalBinary(b []byte) error {
	var res RESTWafGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
