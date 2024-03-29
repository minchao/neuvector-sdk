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

// RESTSystemConfigV2 r e s t system config v2
//
// swagger:model RESTSystemConfigV2
type RESTSystemConfigV2 struct {

	// auth
	// Required: true
	Auth *RESTSystemConfigAuthV2 `json:"auth"`

	// ibmsa
	// Required: true
	Ibmsa *RESTSystemConfigIBMSAV2 `json:"ibmsa"`

	// misc
	// Required: true
	Misc *RESTSystemConfigMiscV2 `json:"misc"`

	// mode auto
	// Required: true
	ModeAuto *RESTSystemConfigModeAutoV2 `json:"mode_auto"`

	// net svc
	// Required: true
	NetSvc *RESTSystemConfigNetSvcV2 `json:"net_svc"`

	// new svc
	// Required: true
	NewSvc *RESTSystemConfigNewSvcV2 `json:"new_svc"`

	// proxy
	// Required: true
	Proxy *RESTSystemConfigProxyV2 `json:"proxy"`

	// scanner autoscale
	// Required: true
	ScannerAutoscale *RESTSystemConfigAutoscale `json:"scanner_autoscale"`

	// syslog
	// Required: true
	Syslog *RESTSystemConfigSyslogV2 `json:"syslog"`

	// webhooks
	// Required: true
	Webhooks []*RESTWebhook `json:"webhooks"`
}

// Validate validates this r e s t system config v2
func (m *RESTSystemConfigV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIbmsa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMisc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeAuto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetSvc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewSvc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScannerAutoscale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyslog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemConfigV2) validateAuth(formats strfmt.Registry) error {

	if err := validate.Required("auth", "body", m.Auth); err != nil {
		return err
	}

	if m.Auth != nil {
		if err := m.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateIbmsa(formats strfmt.Registry) error {

	if err := validate.Required("ibmsa", "body", m.Ibmsa); err != nil {
		return err
	}

	if m.Ibmsa != nil {
		if err := m.Ibmsa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ibmsa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ibmsa")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateMisc(formats strfmt.Registry) error {

	if err := validate.Required("misc", "body", m.Misc); err != nil {
		return err
	}

	if m.Misc != nil {
		if err := m.Misc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("misc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("misc")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateModeAuto(formats strfmt.Registry) error {

	if err := validate.Required("mode_auto", "body", m.ModeAuto); err != nil {
		return err
	}

	if m.ModeAuto != nil {
		if err := m.ModeAuto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_auto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_auto")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateNetSvc(formats strfmt.Registry) error {

	if err := validate.Required("net_svc", "body", m.NetSvc); err != nil {
		return err
	}

	if m.NetSvc != nil {
		if err := m.NetSvc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net_svc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("net_svc")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateNewSvc(formats strfmt.Registry) error {

	if err := validate.Required("new_svc", "body", m.NewSvc); err != nil {
		return err
	}

	if m.NewSvc != nil {
		if err := m.NewSvc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("new_svc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("new_svc")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateProxy(formats strfmt.Registry) error {

	if err := validate.Required("proxy", "body", m.Proxy); err != nil {
		return err
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateScannerAutoscale(formats strfmt.Registry) error {

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

func (m *RESTSystemConfigV2) validateSyslog(formats strfmt.Registry) error {

	if err := validate.Required("syslog", "body", m.Syslog); err != nil {
		return err
	}

	if m.Syslog != nil {
		if err := m.Syslog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syslog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syslog")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) validateWebhooks(formats strfmt.Registry) error {

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

// ContextValidate validate this r e s t system config v2 based on the context it is used
func (m *RESTSystemConfigV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIbmsa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMisc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeAuto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetSvc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewSvc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScannerAutoscale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyslog(ctx, formats); err != nil {
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

func (m *RESTSystemConfigV2) contextValidateAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.Auth != nil {

		if err := m.Auth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateIbmsa(ctx context.Context, formats strfmt.Registry) error {

	if m.Ibmsa != nil {

		if err := m.Ibmsa.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ibmsa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ibmsa")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateMisc(ctx context.Context, formats strfmt.Registry) error {

	if m.Misc != nil {

		if err := m.Misc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("misc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("misc")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateModeAuto(ctx context.Context, formats strfmt.Registry) error {

	if m.ModeAuto != nil {

		if err := m.ModeAuto.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_auto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_auto")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateNetSvc(ctx context.Context, formats strfmt.Registry) error {

	if m.NetSvc != nil {

		if err := m.NetSvc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net_svc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("net_svc")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateNewSvc(ctx context.Context, formats strfmt.Registry) error {

	if m.NewSvc != nil {

		if err := m.NewSvc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("new_svc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("new_svc")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxy != nil {

		if err := m.Proxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateScannerAutoscale(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RESTSystemConfigV2) contextValidateSyslog(ctx context.Context, formats strfmt.Registry) error {

	if m.Syslog != nil {

		if err := m.Syslog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syslog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syslog")
			}
			return err
		}
	}

	return nil
}

func (m *RESTSystemConfigV2) contextValidateWebhooks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RESTSystemConfigV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemConfigV2) UnmarshalBinary(b []byte) error {
	var res RESTSystemConfigV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
