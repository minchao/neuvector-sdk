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

// RESTSystemConfigMiscCfgV2 r e s t system config misc cfg v2
//
// swagger:model RESTSystemConfigMiscCfgV2
type RESTSystemConfigMiscCfgV2 struct {

	// cluster name
	// Example: cluster1
	ClusterName string `json:"cluster_name,omitempty"`

	// controller debug
	// Example: ["scan","k8s_monitor"]
	ControllerDebug []string `json:"controller_debug"`

	// monitor service mesh
	// Example: true
	MonitorServiceMesh bool `json:"monitor_service_mesh,omitempty"`

	// no telemetry report
	// Example: false
	NoTelemetryReport bool `json:"no_telemetry_report,omitempty"`

	// unused group aging
	// Example: 30
	UnusedGroupAging uint8 `json:"unused_group_aging,omitempty"`

	// xff enabled
	// Example: false
	XffEnabled bool `json:"xff_enabled,omitempty"`
}

// Validate validates this r e s t system config misc cfg v2
func (m *RESTSystemConfigMiscCfgV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControllerDebug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rESTSystemConfigMiscCfgV2ControllerDebugItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cpath","conn","mutex","scan","cluster","k8s_monitor"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTSystemConfigMiscCfgV2ControllerDebugItemsEnum = append(rESTSystemConfigMiscCfgV2ControllerDebugItemsEnum, v)
	}
}

func (m *RESTSystemConfigMiscCfgV2) validateControllerDebugItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTSystemConfigMiscCfgV2ControllerDebugItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTSystemConfigMiscCfgV2) validateControllerDebug(formats strfmt.Registry) error {
	if swag.IsZero(m.ControllerDebug) { // not required
		return nil
	}

	for i := 0; i < len(m.ControllerDebug); i++ {

		// value enum
		if err := m.validateControllerDebugItemsEnum("controller_debug"+"."+strconv.Itoa(i), "body", m.ControllerDebug[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this r e s t system config misc cfg v2 based on context it is used
func (m *RESTSystemConfigMiscCfgV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemConfigMiscCfgV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemConfigMiscCfgV2) UnmarshalBinary(b []byte) error {
	var res RESTSystemConfigMiscCfgV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
