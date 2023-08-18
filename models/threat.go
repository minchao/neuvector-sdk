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

// Threat threat
//
// swagger:model Threat
type Threat struct {

	// action
	// Example: monitor
	// Required: true
	Action *string `json:"action"`

	// application
	// Example: HTTP
	// Required: true
	Application *string `json:"application"`

	// cap len
	// Example: 32043
	CapLen uint16 `json:"cap_len,omitempty"`

	// client ip
	// Example: 172.17.0.4
	// Required: true
	ClientIP *string `json:"client_ip"`

	// client port
	// Example: 53825
	// Required: true
	ClientPort *uint16 `json:"client_port"`

	// client workload domain
	ClientWorkloadDomain string `json:"client_workload_domain,omitempty"`

	// client workload id
	// Example: 7e31a2a4d4074ba459f8c22ee90
	// Required: true
	ClientWorkloadID *string `json:"client_workload_id"`

	// client workload image
	ClientWorkloadImage string `json:"client_workload_image,omitempty"`

	// client workload name
	// Example: iperfclient
	// Required: true
	ClientWorkloadName *string `json:"client_workload_name"`

	// client workload service
	ClientWorkloadService string `json:"client_workload_service,omitempty"`

	// cluster name
	// Example: cluster1
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// count
	// Example: 1
	// Required: true
	Count *uint32 `json:"count"`

	// enforcer id
	// Example: 2fdc03e027d6483633376609
	// Required: true
	EnforcerID *string `json:"enforcer_id"`

	// enforcer name
	// Example: allinone
	// Required: true
	EnforcerName *string `json:"enforcer_name"`

	// ether type
	// Example: 2048
	// Required: true
	EtherType *uint16 `json:"ether_type"`

	// group
	// Required: true
	Group *string `json:"group"`

	// host id
	// Example: ubuntu64:ZYA3:NPGZ:IU64:XJ3P:SUU3:QYEF:SLVT
	// Required: true
	HostID *string `json:"host_id"`

	// host name
	// Example: ubuntu64
	// Required: true
	HostName *string `json:"host_name"`

	// icmp code
	// Example: 0
	// Required: true
	IcmpCode *uint8 `json:"icmp_code"`

	// icmp type
	// Example: 0
	// Required: true
	IcmpType *uint8 `json:"icmp_type"`

	// id
	// Example: 999-9c96-11e7-83c1-17580b
	// Required: true
	ID *string `json:"id"`

	// ip proto
	// Example: 6
	// Required: true
	IPProto *uint8 `json:"ip_proto"`

	// level
	// Example: CRIT
	// Required: true
	Level *string `json:"level"`

	// message
	// Example: Header duration=3s, threshold=3s
	// Required: true
	Message *string `json:"message"`

	// monitor
	// Example: true
	// Required: true
	Monitor *bool `json:"monitor"`

	// name
	// Example: HTTP.Request.Slowloris
	Name string `json:"name,omitempty"`

	// packet
	// Example: base64string
	Packet string `json:"packet,omitempty"`

	// reported at
	// Example: 2017-09-18T17:28:32Z
	// Required: true
	ReportedAt *string `json:"reported_at"`

	// reported timestamp
	// Example: 1505755712
	// Required: true
	ReportedTimestamp *int64 `json:"reported_timestamp"`

	// response rule id
	// Example: 10005
	ResponseRuleID int64 `json:"response_rule_id,omitempty"`

	// sensor
	// Required: true
	Sensor *string `json:"sensor"`

	// server conn port
	// Example: 5000
	// Required: true
	ServerConnPort *uint16 `json:"server_conn_port"`

	// server ip
	// Example: 172.17.0.3
	// Required: true
	ServerIP *string `json:"server_ip"`

	// server port
	// Example: 5000
	// Required: true
	ServerPort *uint16 `json:"server_port"`

	// server workload domain
	ServerWorkloadDomain string `json:"server_workload_domain,omitempty"`

	// server workload id
	// Example: 2cbab37b43efe049e583924a73a764b096ce8f15ea
	// Required: true
	ServerWorkloadID *string `json:"server_workload_id"`

	// server workload image
	ServerWorkloadImage string `json:"server_workload_image,omitempty"`

	// server workload name
	// Example: iperfserver
	// Required: true
	ServerWorkloadName *string `json:"server_workload_name"`

	// server workload service
	ServerWorkloadService string `json:"server_workload_service,omitempty"`

	// severity
	// Example: critical
	// Required: true
	Severity *string `json:"severity"`

	// target
	// Example: true
	// Required: true
	Target *string `json:"target"`

	// threat id
	// Example: 2017
	// Required: true
	ThreatID *uint32 `json:"threat_id"`
}

// Validate validates this threat
func (m *Threat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientWorkloadID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientWorkloadName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforcerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforcerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEtherType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmpCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmpType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPProto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerConnPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerWorkloadID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerWorkloadName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreatID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Threat) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateApplication(formats strfmt.Registry) error {

	if err := validate.Required("application", "body", m.Application); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateClientIP(formats strfmt.Registry) error {

	if err := validate.Required("client_ip", "body", m.ClientIP); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateClientPort(formats strfmt.Registry) error {

	if err := validate.Required("client_port", "body", m.ClientPort); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateClientWorkloadID(formats strfmt.Registry) error {

	if err := validate.Required("client_workload_id", "body", m.ClientWorkloadID); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateClientWorkloadName(formats strfmt.Registry) error {

	if err := validate.Required("client_workload_name", "body", m.ClientWorkloadName); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("cluster_name", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateEnforcerID(formats strfmt.Registry) error {

	if err := validate.Required("enforcer_id", "body", m.EnforcerID); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateEnforcerName(formats strfmt.Registry) error {

	if err := validate.Required("enforcer_name", "body", m.EnforcerName); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateEtherType(formats strfmt.Registry) error {

	if err := validate.Required("ether_type", "body", m.EtherType); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("host_name", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateIcmpCode(formats strfmt.Registry) error {

	if err := validate.Required("icmp_code", "body", m.IcmpCode); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateIcmpType(formats strfmt.Registry) error {

	if err := validate.Required("icmp_type", "body", m.IcmpType); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateIPProto(formats strfmt.Registry) error {

	if err := validate.Required("ip_proto", "body", m.IPProto); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateMonitor(formats strfmt.Registry) error {

	if err := validate.Required("monitor", "body", m.Monitor); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateReportedAt(formats strfmt.Registry) error {

	if err := validate.Required("reported_at", "body", m.ReportedAt); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateReportedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("reported_timestamp", "body", m.ReportedTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateSensor(formats strfmt.Registry) error {

	if err := validate.Required("sensor", "body", m.Sensor); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateServerConnPort(formats strfmt.Registry) error {

	if err := validate.Required("server_conn_port", "body", m.ServerConnPort); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateServerIP(formats strfmt.Registry) error {

	if err := validate.Required("server_ip", "body", m.ServerIP); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateServerPort(formats strfmt.Registry) error {

	if err := validate.Required("server_port", "body", m.ServerPort); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateServerWorkloadID(formats strfmt.Registry) error {

	if err := validate.Required("server_workload_id", "body", m.ServerWorkloadID); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateServerWorkloadName(formats strfmt.Registry) error {

	if err := validate.Required("server_workload_name", "body", m.ServerWorkloadName); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

func (m *Threat) validateThreatID(formats strfmt.Registry) error {

	if err := validate.Required("threat_id", "body", m.ThreatID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this threat based on context it is used
func (m *Threat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Threat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Threat) UnmarshalBinary(b []byte) error {
	var res Threat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
