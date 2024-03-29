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

// RESTAdmRuleTypeOptions r e s t adm rule type options
//
// swagger:model RESTAdmRuleTypeOptions
type RESTAdmRuleTypeOptions struct {

	// deny options
	// Required: true
	DenyOptions *RESTAdmCatOptions `json:"deny_options"`

	// exception options
	// Required: true
	ExceptionOptions *RESTAdmCatOptions `json:"exception_options"`

	// psp collection
	PspCollection []*RESTAdmRuleCriterion `json:"psp_collection"`

	// map key is domain(string type)
	// Example: {"baseline":["Sets HostNetwork, HostPID, or HostIPC to true.","Allows privileged container(s)."],"restricted":["Uses illegal volume type.","Allows running as root user."]}
	PssCollections map[string][]string `json:"pss_collections,omitempty"`

	// sigstore verifiers
	// Example: ["public/verifier1","private1/verifier1","private1/verifier2"]
	SigstoreVerifiers []string `json:"sigstore_verifiers"`
}

// Validate validates this r e s t adm rule type options
func (m *RESTAdmRuleTypeOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDenyOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExceptionOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePspCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdmRuleTypeOptions) validateDenyOptions(formats strfmt.Registry) error {

	if err := validate.Required("deny_options", "body", m.DenyOptions); err != nil {
		return err
	}

	if m.DenyOptions != nil {
		if err := m.DenyOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deny_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deny_options")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAdmRuleTypeOptions) validateExceptionOptions(formats strfmt.Registry) error {

	if err := validate.Required("exception_options", "body", m.ExceptionOptions); err != nil {
		return err
	}

	if m.ExceptionOptions != nil {
		if err := m.ExceptionOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exception_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exception_options")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAdmRuleTypeOptions) validatePspCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.PspCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.PspCollection); i++ {
		if swag.IsZero(m.PspCollection[i]) { // not required
			continue
		}

		if m.PspCollection[i] != nil {
			if err := m.PspCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("psp_collection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("psp_collection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this r e s t adm rule type options based on the context it is used
func (m *RESTAdmRuleTypeOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDenyOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExceptionOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePspCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdmRuleTypeOptions) contextValidateDenyOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DenyOptions != nil {

		if err := m.DenyOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deny_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deny_options")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAdmRuleTypeOptions) contextValidateExceptionOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ExceptionOptions != nil {

		if err := m.ExceptionOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exception_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exception_options")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAdmRuleTypeOptions) contextValidatePspCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PspCollection); i++ {

		if m.PspCollection[i] != nil {

			if swag.IsZero(m.PspCollection[i]) { // not required
				return nil
			}

			if err := m.PspCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("psp_collection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("psp_collection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdmRuleTypeOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdmRuleTypeOptions) UnmarshalBinary(b []byte) error {
	var res RESTAdmRuleTypeOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
