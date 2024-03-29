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

// Incident incident
//
// swagger:model Incident
type Incident struct {

	// action
	// Example: violate
	// Required: true
	Action *string `json:"action"`

	// aggregation from
	// Example: 1515020888
	// Required: true
	AggregationFrom *int64 `json:"aggregation_from"`

	// client ip
	// Example: 192.168.1.62
	// Required: true
	ClientIP *string `json:"client_ip"`

	// client port
	// Example: 56564
	// Required: true
	ClientPort *uint16 `json:"client_port"`

	// cluster name
	// Example: cluster1
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// conn ingress
	// Example: false
	// Required: true
	ConnIngress *bool `json:"conn_ingress"`

	// count
	// Example: 1
	// Required: true
	Count *int64 `json:"count"`

	// enforcer id
	// Example: a928be54f34fbb696426890a7249c067
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

	// file name
	// Required: true
	FileName []string `json:"file_name"`

	// file path
	// Required: true
	FilePath *string `json:"file_path"`

	// group
	// Example: nv.iperfserver
	// Required: true
	Group *string `json:"group"`

	// host id
	// Example: ubuntu64:ZYA3:UZL5:2WOT:QYEF:SLVT:LIBD
	// Required: true
	HostID *string `json:"host_id"`

	// host name
	// Example: ubuntu64
	// Required: true
	HostName *string `json:"host_name"`

	// id
	// Example: e6e19591-75a0-43e9-bebb-145d588e6718
	// Required: true
	ID *string `json:"id"`

	// ip proto
	// Example: 6
	// Required: true
	IPProto *uint8 `json:"ip_proto"`

	// level
	// Example: WARNING
	// Required: true
	Level *string `json:"level"`

	// message
	// Example: dns tunneling
	// Required: true
	Message *string `json:"message"`

	// name
	// Example: Host.Suspicious.Process
	// Required: true
	Name *string `json:"name"`

	// proc cmd
	// Example: ./dns_tunneling/iodine/bin/iodine -f -r 172.17.0.3 -P
	// Required: true
	ProcCmd *string `json:"proc_cmd"`

	// proc effective uid
	// Example: 1000
	// Required: true
	ProcEffectiveUID *int64 `json:"proc_effective_uid"`

	// proc effective user
	// Example: root
	// Required: true
	ProcEffectiveUser *string `json:"proc_effective_user"`

	// proc name
	// Example: nc
	// Required: true
	ProcName *string `json:"proc_name"`

	// proc parent name
	// Example: sh
	// Required: true
	ProcParentName *string `json:"proc_parent_name"`

	// proc parent path
	// Example: /bin/dash
	// Required: true
	ProcParentPath *string `json:"proc_parent_path"`

	// proc path
	// Example: /bin/nc.traditional
	// Required: true
	ProcPath *string `json:"proc_path"`

	// proc real uid
	// Example: 1000
	// Required: true
	ProcRealUID *int64 `json:"proc_real_uid"`

	// proc real user
	// Example: test
	// Required: true
	ProcRealUser *string `json:"proc_real_user"`

	// remote workload domain
	// Required: true
	RemoteWorkloadDomain *string `json:"remote_workload_domain"`

	// remote workload id
	// Example: external
	// Required: true
	RemoteWorkloadID *string `json:"remote_workload_id"`

	// remote workload image
	// Example: iperfclient
	// Required: true
	RemoteWorkloadImage *string `json:"remote_workload_image"`

	// remote workload name
	// Example: iperfclient
	// Required: true
	RemoteWorkloadName *string `json:"remote_workload_name"`

	// remote workload service
	// Example: iperfclient
	// Required: true
	RemoteWorkloadService *string `json:"remote_workload_service"`

	// reported at
	// Example: 2018-01-03T23:08:08Z
	// Required: true
	// Format: date-time
	ReportedAt *strfmt.DateTime `json:"reported_at"`

	// reported timestamp
	// Example: 1515020888
	// Required: true
	ReportedTimestamp *int64 `json:"reported_timestamp"`

	// response rule id
	// Example: 10006
	// Required: true
	ResponseRuleID *int64 `json:"response_rule_id"`

	// rule id
	// Example: 00000000-0000-0000-0000-000000000001
	// Required: true
	RuleID *string `json:"rule_id"`

	// server conn port
	// Example: 80
	// Required: true
	ServerConnPort *uint16 `json:"server_conn_port"`

	// server ip
	// Example: 10.1.4.3
	// Required: true
	ServerIP *string `json:"server_ip"`

	// server port
	// Example: 80
	// Required: true
	ServerPort *uint16 `json:"server_port"`

	// workload domain
	// Required: true
	WorkloadDomain *string `json:"workload_domain"`

	// workload id
	// Example: 83e76eabd68494649440fa0a35451315289c70eb3094454e419952dffaa7715a
	// Required: true
	WorkloadID *string `json:"workload_id"`

	// workload image
	// Example: iperfserver
	// Required: true
	WorkloadImage *string `json:"workload_image"`

	// workload name
	// Example: iperfserver
	// Required: true
	WorkloadName *string `json:"workload_name"`

	// workload service
	// Example: iperfserver
	// Required: true
	WorkloadService *string `json:"workload_service"`
}

// Validate validates this incident
func (m *Incident) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregationFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnIngress(formats); err != nil {
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

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilePath(formats); err != nil {
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

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcCmd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcEffectiveUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcEffectiveUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcParentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcParentPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcRealUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcRealUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteWorkloadDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteWorkloadID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteWorkloadImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteWorkloadName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteWorkloadService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseRuleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleID(formats); err != nil {
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

func (m *Incident) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateAggregationFrom(formats strfmt.Registry) error {

	if err := validate.Required("aggregation_from", "body", m.AggregationFrom); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateClientIP(formats strfmt.Registry) error {

	if err := validate.Required("client_ip", "body", m.ClientIP); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateClientPort(formats strfmt.Registry) error {

	if err := validate.Required("client_port", "body", m.ClientPort); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("cluster_name", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateConnIngress(formats strfmt.Registry) error {

	if err := validate.Required("conn_ingress", "body", m.ConnIngress); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateEnforcerID(formats strfmt.Registry) error {

	if err := validate.Required("enforcer_id", "body", m.EnforcerID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateEnforcerName(formats strfmt.Registry) error {

	if err := validate.Required("enforcer_name", "body", m.EnforcerName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateEtherType(formats strfmt.Registry) error {

	if err := validate.Required("ether_type", "body", m.EtherType); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("file_name", "body", m.FileName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateFilePath(formats strfmt.Registry) error {

	if err := validate.Required("file_path", "body", m.FilePath); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("host_name", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateIPProto(formats strfmt.Registry) error {

	if err := validate.Required("ip_proto", "body", m.IPProto); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcCmd(formats strfmt.Registry) error {

	if err := validate.Required("proc_cmd", "body", m.ProcCmd); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcEffectiveUID(formats strfmt.Registry) error {

	if err := validate.Required("proc_effective_uid", "body", m.ProcEffectiveUID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcEffectiveUser(formats strfmt.Registry) error {

	if err := validate.Required("proc_effective_user", "body", m.ProcEffectiveUser); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcName(formats strfmt.Registry) error {

	if err := validate.Required("proc_name", "body", m.ProcName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcParentName(formats strfmt.Registry) error {

	if err := validate.Required("proc_parent_name", "body", m.ProcParentName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcParentPath(formats strfmt.Registry) error {

	if err := validate.Required("proc_parent_path", "body", m.ProcParentPath); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcPath(formats strfmt.Registry) error {

	if err := validate.Required("proc_path", "body", m.ProcPath); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcRealUID(formats strfmt.Registry) error {

	if err := validate.Required("proc_real_uid", "body", m.ProcRealUID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateProcRealUser(formats strfmt.Registry) error {

	if err := validate.Required("proc_real_user", "body", m.ProcRealUser); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateRemoteWorkloadDomain(formats strfmt.Registry) error {

	if err := validate.Required("remote_workload_domain", "body", m.RemoteWorkloadDomain); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateRemoteWorkloadID(formats strfmt.Registry) error {

	if err := validate.Required("remote_workload_id", "body", m.RemoteWorkloadID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateRemoteWorkloadImage(formats strfmt.Registry) error {

	if err := validate.Required("remote_workload_image", "body", m.RemoteWorkloadImage); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateRemoteWorkloadName(formats strfmt.Registry) error {

	if err := validate.Required("remote_workload_name", "body", m.RemoteWorkloadName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateRemoteWorkloadService(formats strfmt.Registry) error {

	if err := validate.Required("remote_workload_service", "body", m.RemoteWorkloadService); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateReportedAt(formats strfmt.Registry) error {

	if err := validate.Required("reported_at", "body", m.ReportedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("reported_at", "body", "date-time", m.ReportedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateReportedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("reported_timestamp", "body", m.ReportedTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateResponseRuleID(formats strfmt.Registry) error {

	if err := validate.Required("response_rule_id", "body", m.ResponseRuleID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateRuleID(formats strfmt.Registry) error {

	if err := validate.Required("rule_id", "body", m.RuleID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateServerConnPort(formats strfmt.Registry) error {

	if err := validate.Required("server_conn_port", "body", m.ServerConnPort); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateServerIP(formats strfmt.Registry) error {

	if err := validate.Required("server_ip", "body", m.ServerIP); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateServerPort(formats strfmt.Registry) error {

	if err := validate.Required("server_port", "body", m.ServerPort); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateWorkloadDomain(formats strfmt.Registry) error {

	if err := validate.Required("workload_domain", "body", m.WorkloadDomain); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateWorkloadID(formats strfmt.Registry) error {

	if err := validate.Required("workload_id", "body", m.WorkloadID); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateWorkloadImage(formats strfmt.Registry) error {

	if err := validate.Required("workload_image", "body", m.WorkloadImage); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateWorkloadName(formats strfmt.Registry) error {

	if err := validate.Required("workload_name", "body", m.WorkloadName); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateWorkloadService(formats strfmt.Registry) error {

	if err := validate.Required("workload_service", "body", m.WorkloadService); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this incident based on context it is used
func (m *Incident) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Incident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Incident) UnmarshalBinary(b []byte) error {
	var res Incident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
