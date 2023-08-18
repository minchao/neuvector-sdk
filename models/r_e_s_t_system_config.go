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

// RESTSystemConfig r e s t system config
//
// swagger:model RESTSystemConfig
type RESTSystemConfig struct {

	// auth by platform
	// Example: true
	// Required: true
	AuthByPlatform *bool `json:"auth_by_platform"`

	// auth order
	// Example: ["local","ldap"]
	// Required: true
	AuthOrder []string `json:"auth_order"`

	// cluster name
	// Example: cluster1
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// configured internal subnets
	// Example: ["69.89.0.0/16","172.217.5.0/23"]
	ConfiguredInternalSubnets []string `json:"configured_internal_subnets"`

	// controller debug
	// Example: ["cpath","scan","k8s_monitor"]
	// Required: true
	ControllerDebug []string `json:"controller_debug"`

	// csp type
	// Example: aws
	CspType string `json:"csp_type,omitempty"`

	// ibmsa ep connected at
	// Required: true
	IbmsaEpConnectedAt *string `json:"ibmsa_ep_connected_at"`

	// ibmsa ep dashboard url
	// Required: true
	IbmsaEpDashboardURL *string `json:"ibmsa_ep_dashboard_url"`

	// ibmsa ep enabled
	// Example: false
	// Required: true
	IbmsaEpEnabled *bool `json:"ibmsa_ep_enabled"`

	// ibmsa ep start
	// Example: 1
	// Required: true
	IbmsaEpStart *uint32 `json:"ibmsa_ep_start"`

	// mode auto d2m
	// Example: false
	// Required: true
	ModeAutoD2m *bool `json:"mode_auto_d2m"`

	// mode auto d2m duration
	// Example: 1505755716
	// Required: true
	ModeAutoD2mDuration *int64 `json:"mode_auto_d2m_duration"`

	// mode auto m2p
	// Example: false
	// Required: true
	ModeAutoM2p *bool `json:"mode_auto_m2p"`

	// mode auto m2p duration
	// Example: 1505755716
	// Required: true
	ModeAutoM2pDuration *int64 `json:"mode_auto_m2p_duration"`

	// monitor service mesh
	// Example: true
	// Required: true
	MonitorServiceMesh *bool `json:"monitor_service_mesh"`

	// net service policy mode
	// Example: Monitor
	// Required: true
	NetServicePolicyMode *string `json:"net_service_policy_mode"`

	// net service status
	// Example: false
	// Required: true
	NetServiceStatus *bool `json:"net_service_status"`

	// new service policy mode
	// Example: Discover
	// Required: true
	NewServicePolicyMode *string `json:"new_service_policy_mode"`

	// new service profile baseline
	// Example: zero-drift
	// Required: true
	NewServiceProfileBaseline *string `json:"new_service_profile_baseline"`

	// no telemetry report
	// Example: false
	// Required: true
	NoTelemetryReport *bool `json:"no_telemetry_report"`

	// rancher ep
	// Required: true
	RancherEp *string `json:"rancher_ep"`

	// registry http proxy
	// Required: true
	RegistryHTTPProxy *RESTProxy `json:"registry_http_proxy"`

	// registry http proxy status
	// Example: true
	// Required: true
	RegistryHTTPProxyStatus *bool `json:"registry_http_proxy_status"`

	// registry https proxy
	// Required: true
	RegistryHTTPSProxy *RESTProxy `json:"registry_https_proxy"`

	// registry https proxy status
	// Example: false
	// Required: true
	RegistryHTTPSProxyStatus *bool `json:"registry_https_proxy_status"`

	// scanner autoscale
	// Required: true
	ScannerAutoscale *RESTSystemConfigAutoscale `json:"scanner_autoscale"`

	// single cve per syslog
	// Example: false
	// Required: true
	SingleCvePerSyslog *bool `json:"single_cve_per_syslog"`

	// syslog categories
	// Example: ["event","violation","threat","incident"]
	// Required: true
	SyslogCategories []string `json:"syslog_categories"`

	// syslog cve in layers
	// Example: false
	// Required: true
	SyslogCveInLayers *bool `json:"syslog_cve_in_layers"`

	// syslog in json
	// Example: true
	// Required: true
	SyslogInJSON *bool `json:"syslog_in_json"`

	// syslog ip
	// Example: 10.1.0.14
	// Required: true
	SyslogIP *string `json:"syslog_ip"`

	// syslog ip proto
	// Example: 6
	// Required: true
	SyslogIPProto *uint8 `json:"syslog_ip_proto"`

	// syslog level
	// Example: INFO
	// Required: true
	SyslogLevel *string `json:"syslog_level"`

	// syslog port
	// Example: 514
	// Required: true
	SyslogPort *uint16 `json:"syslog_port"`

	// syslog server cert
	// Example: E7B0OS/N3KMVCL6KNMZ2+LOV90S7854NSD84P0BF
	SyslogServerCert string `json:"syslog_server_cert,omitempty"`

	// syslog status
	// Example: false
	// Required: true
	SyslogStatus *bool `json:"syslog_status"`

	// unused group aging
	// Example: 123
	// Required: true
	UnusedGroupAging *uint8 `json:"unused_group_aging"`

	// webhooks
	// Required: true
	Webhooks []*RESTWebhook `json:"webhooks"`

	// xff enabled
	// Example: false
	// Required: true
	XffEnabled *bool `json:"xff_enabled"`
}

// Validate validates this r e s t system config
func (m *RESTSystemConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthByPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerDebug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIbmsaEpConnectedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIbmsaEpDashboardURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIbmsaEpEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIbmsaEpStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeAutoD2m(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeAutoD2mDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeAutoM2p(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeAutoM2pDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitorServiceMesh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetServicePolicyMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetServiceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewServicePolicyMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewServiceProfileBaseline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoTelemetryReport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRancherEp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryHTTPProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryHTTPProxyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryHTTPSProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryHTTPSProxyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScannerAutoscale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingleCvePerSyslog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogCveInLayers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogInJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogIPProto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslogStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnusedGroupAging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhooks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXffEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfig) validateAuthByPlatform(formats strfmt.Registry) error {

	if err := validate.Required("auth_by_platform", "body", m.AuthByPlatform); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateAuthOrder(formats strfmt.Registry) error {

	if err := validate.Required("auth_order", "body", m.AuthOrder); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("cluster_name", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

var rESTSystemConfigControllerDebugItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cpath","conn","mutex","scan","cluster","k8s_monitor"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rESTSystemConfigControllerDebugItemsEnum = append(rESTSystemConfigControllerDebugItemsEnum, v)
	}
}

func (m *RESTSystemConfig) validateControllerDebugItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rESTSystemConfigControllerDebugItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RESTSystemConfig) validateControllerDebug(formats strfmt.Registry) error {

	if err := validate.Required("controller_debug", "body", m.ControllerDebug); err != nil {
		return err
	}

	for i := 0; i < len(m.ControllerDebug); i++ {

		// value enum
		if err := m.validateControllerDebugItemsEnum("controller_debug"+"."+strconv.Itoa(i), "body", m.ControllerDebug[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *RESTSystemConfig) validateIbmsaEpConnectedAt(formats strfmt.Registry) error {

	if err := validate.Required("ibmsa_ep_connected_at", "body", m.IbmsaEpConnectedAt); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateIbmsaEpDashboardURL(formats strfmt.Registry) error {

	if err := validate.Required("ibmsa_ep_dashboard_url", "body", m.IbmsaEpDashboardURL); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateIbmsaEpEnabled(formats strfmt.Registry) error {

	if err := validate.Required("ibmsa_ep_enabled", "body", m.IbmsaEpEnabled); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateIbmsaEpStart(formats strfmt.Registry) error {

	if err := validate.Required("ibmsa_ep_start", "body", m.IbmsaEpStart); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateModeAutoD2m(formats strfmt.Registry) error {

	if err := validate.Required("mode_auto_d2m", "body", m.ModeAutoD2m); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateModeAutoD2mDuration(formats strfmt.Registry) error {

	if err := validate.Required("mode_auto_d2m_duration", "body", m.ModeAutoD2mDuration); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateModeAutoM2p(formats strfmt.Registry) error {

	if err := validate.Required("mode_auto_m2p", "body", m.ModeAutoM2p); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateModeAutoM2pDuration(formats strfmt.Registry) error {

	if err := validate.Required("mode_auto_m2p_duration", "body", m.ModeAutoM2pDuration); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateMonitorServiceMesh(formats strfmt.Registry) error {

	if err := validate.Required("monitor_service_mesh", "body", m.MonitorServiceMesh); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateNetServicePolicyMode(formats strfmt.Registry) error {

	if err := validate.Required("net_service_policy_mode", "body", m.NetServicePolicyMode); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateNetServiceStatus(formats strfmt.Registry) error {

	if err := validate.Required("net_service_status", "body", m.NetServiceStatus); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateNewServicePolicyMode(formats strfmt.Registry) error {

	if err := validate.Required("new_service_policy_mode", "body", m.NewServicePolicyMode); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateNewServiceProfileBaseline(formats strfmt.Registry) error {

	if err := validate.Required("new_service_profile_baseline", "body", m.NewServiceProfileBaseline); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateNoTelemetryReport(formats strfmt.Registry) error {

	if err := validate.Required("no_telemetry_report", "body", m.NoTelemetryReport); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateRancherEp(formats strfmt.Registry) error {

	if err := validate.Required("rancher_ep", "body", m.RancherEp); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateRegistryHTTPProxy(formats strfmt.Registry) error {

	if err := validate.Required("registry_http_proxy", "body", m.RegistryHTTPProxy); err != nil {
		return err
	}

	if m.RegistryHTTPProxy != nil {
		if err := m.RegistryHTTPProxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_http_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_http_proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfig) validateRegistryHTTPProxyStatus(formats strfmt.Registry) error {

	if err := validate.Required("registry_http_proxy_status", "body", m.RegistryHTTPProxyStatus); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateRegistryHTTPSProxy(formats strfmt.Registry) error {

	if err := validate.Required("registry_https_proxy", "body", m.RegistryHTTPSProxy); err != nil {
		return err
	}

	if m.RegistryHTTPSProxy != nil {
		if err := m.RegistryHTTPSProxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_https_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_https_proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfig) validateRegistryHTTPSProxyStatus(formats strfmt.Registry) error {

	if err := validate.Required("registry_https_proxy_status", "body", m.RegistryHTTPSProxyStatus); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateScannerAutoscale(formats strfmt.Registry) error {

	if err := validate.Required("scanner_autoscale", "body", m.ScannerAutoscale); err != nil {
		return err
	}

	if m.ScannerAutoscale != nil {
		if err := m.ScannerAutoscale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner_autoscale")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanner_autoscale")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfig) validateSingleCvePerSyslog(formats strfmt.Registry) error {

	if err := validate.Required("single_cve_per_syslog", "body", m.SingleCvePerSyslog); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogCategories(formats strfmt.Registry) error {

	if err := validate.Required("syslog_categories", "body", m.SyslogCategories); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogCveInLayers(formats strfmt.Registry) error {

	if err := validate.Required("syslog_cve_in_layers", "body", m.SyslogCveInLayers); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogInJSON(formats strfmt.Registry) error {

	if err := validate.Required("syslog_in_json", "body", m.SyslogInJSON); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogIP(formats strfmt.Registry) error {

	if err := validate.Required("syslog_ip", "body", m.SyslogIP); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogIPProto(formats strfmt.Registry) error {

	if err := validate.Required("syslog_ip_proto", "body", m.SyslogIPProto); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogLevel(formats strfmt.Registry) error {

	if err := validate.Required("syslog_level", "body", m.SyslogLevel); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogPort(formats strfmt.Registry) error {

	if err := validate.Required("syslog_port", "body", m.SyslogPort); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateSyslogStatus(formats strfmt.Registry) error {

	if err := validate.Required("syslog_status", "body", m.SyslogStatus); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateUnusedGroupAging(formats strfmt.Registry) error {

	if err := validate.Required("unused_group_aging", "body", m.UnusedGroupAging); err != nil {
		return err
	}

	return nil
}

func (m *RESTSystemConfig) validateWebhooks(formats strfmt.Registry) error {

	if err := validate.Required("webhooks", "body", m.Webhooks); err != nil {
		return err
	}

	for i := 0; i < len(m.Webhooks); i++ {
		if swag.IsZero(m.Webhooks[i]) { // not required
			continue
		}

		if m.Webhooks[i] != nil {
			if err := m.Webhooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RESTSystemConfig) validateXffEnabled(formats strfmt.Registry) error {

	if err := validate.Required("xff_enabled", "body", m.XffEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this r e s t system config based on the context it is used
func (m *RESTSystemConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistryHTTPProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistryHTTPSProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScannerAutoscale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfig) contextValidateRegistryHTTPProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistryHTTPProxy != nil {

		if err := m.RegistryHTTPProxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_http_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_http_proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfig) contextValidateRegistryHTTPSProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistryHTTPSProxy != nil {

		if err := m.RegistryHTTPSProxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_https_proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_https_proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfig) contextValidateScannerAutoscale(ctx context.Context, formats strfmt.Registry) error {

	if m.ScannerAutoscale != nil {

		if err := m.ScannerAutoscale.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner_autoscale")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanner_autoscale")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfig) contextValidateWebhooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Webhooks); i++ {

		if m.Webhooks[i] != nil {

			if swag.IsZero(m.Webhooks[i]) { // not required
				return nil
			}

			if err := m.Webhooks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemConfig) UnmarshalBinary(b []byte) error {
	var res RESTSystemConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}