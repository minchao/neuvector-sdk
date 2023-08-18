// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event event
//
// swagger:model Event
type Event struct {

	// category
	// Example: WORKLOAD
	// Required: true
	Category *string `json:"category"`

	// cluster name
	// Example: cluster1
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// controller id
	// Required: true
	ControllerID *string `json:"controller_id"`

	// controller name
	// Required: true
	ControllerName *string `json:"controller_name"`

	// enforcer id
	// Example: bced57295eefbd3f3bd0cf798b6184fa789
	// Required: true
	EnforcerID *string `json:"enforcer_id"`

	// enforcer limit
	// Example: 0
	EnforcerLimit int64 `json:"enforcer_limit,omitempty"`

	// enforcer name
	// Example: allinone
	// Required: true
	EnforcerName *string `json:"enforcer_name"`

	// host id
	// Example: QK56:LFQP:IJSV:FXBN:QJV7:5MBB:6PL5
	// Required: true
	HostID *string `json:"host_id"`

	// host name
	// Example: ubuntu64
	// Required: true
	HostName *string `json:"host_name"`

	// level
	// Example: INFO
	// Required: true
	Level *string `json:"level"`

	// license expire
	LicenseExpire string `json:"license_expire,omitempty"`

	// message
	// Example: Start managing container kafkadocker_zookeeper_1
	// Required: true
	Message *string `json:"message"`

	// name
	// Example: Container.Managed
	Name string `json:"name,omitempty"`

	// reported at
	// Example: 2018-01-24T22:24:30Z
	// Required: true
	// Format: date-time
	ReportedAt *strfmt.DateTime `json:"reported_at"`

	// reported timestamp
	// Example: 1516832670
	// Required: true
	ReportedTimestamp *int64 `json:"reported_timestamp"`

	// response rule id
	// Example: 10009
	ResponseRuleID int64 `json:"response_rule_id,omitempty"`

	// rest body
	RestBody string `json:"rest_body,omitempty"`

	// rest method
	RestMethod string `json:"rest_method,omitempty"`

	// rest request
	RestRequest string `json:"rest_request,omitempty"`

	// user
	// Example: user
	// Required: true
	User *string `json:"user"`

	// user addr
	// Required: true
	UserAddr *string `json:"user_addr"`

	// map key is domain(string type)
	// Example: {"domain1":"admin","domain2":"reader"}
	// Required: true
	UserRoles map[string]string `json:"user_roles"`

	// user session
	// Required: true
	UserSession *string `json:"user_session"`

	// workload domain
	// Required: true
	WorkloadDomain *string `json:"workload_domain"`

	// workload id
	// Example: 7df6a19648e2860c89fe12c8d5b1c52079a
	// Required: true
	WorkloadID *string `json:"workload_id"`

	// workload image
	// Example: zookeeper
	// Required: true
	WorkloadImage *string `json:"workload_image"`

	// workload name
	// Example: kafkadocker_zookeeper_1
	// Required: true
	WorkloadName *string `json:"workload_name"`

	// workload service
	// Required: true
	WorkloadService *string `json:"workload_service"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforcerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforcerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSession(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("cluster_name", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateControllerID(formats strfmt.Registry) error {

	if err := validate.Required("controller_id", "body", m.ControllerID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateControllerName(formats strfmt.Registry) error {

	if err := validate.Required("controller_name", "body", m.ControllerName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateEnforcerID(formats strfmt.Registry) error {

	if err := validate.Required("enforcer_id", "body", m.EnforcerID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateEnforcerName(formats strfmt.Registry) error {

	if err := validate.Required("enforcer_name", "body", m.EnforcerName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("host_name", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateReportedAt(formats strfmt.Registry) error {

	if err := validate.Required("reported_at", "body", m.ReportedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("reported_at", "body", "date-time", m.ReportedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateReportedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("reported_timestamp", "body", m.ReportedTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateUserAddr(formats strfmt.Registry) error {

	if err := validate.Required("user_addr", "body", m.UserAddr); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateUserRoles(formats strfmt.Registry) error {

	if err := validate.Required("user_roles", "body", m.UserRoles); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateUserSession(formats strfmt.Registry) error {

	if err := validate.Required("user_session", "body", m.UserSession); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateWorkloadDomain(formats strfmt.Registry) error {

	if err := validate.Required("workload_domain", "body", m.WorkloadDomain); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateWorkloadID(formats strfmt.Registry) error {

	if err := validate.Required("workload_id", "body", m.WorkloadID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateWorkloadImage(formats strfmt.Registry) error {

	if err := validate.Required("workload_image", "body", m.WorkloadImage); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateWorkloadName(formats strfmt.Registry) error {

	if err := validate.Required("workload_name", "body", m.WorkloadName); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateWorkloadService(formats strfmt.Registry) error {

	if err := validate.Required("workload_service", "body", m.WorkloadService); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event based on context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
