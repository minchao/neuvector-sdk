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

// RESTAdmissionStats r e s t admission stats
//
// swagger:model RESTAdmissionStats
type RESTAdmissionStats struct {

	// jenkins allowed requests
	// Example: 3
	// Required: true
	JenkinsAllowedRequests *int64 `json:"jenkins_allowed_requests"`

	// jenkins denied requests
	// Example: 1
	// Required: true
	JenkinsDeniedRequests *int64 `json:"jenkins_denied_requests"`

	// jenkins erroneous requests
	// Example: 1
	// Required: true
	JenkinsErroneousRequests *int64 `json:"jenkins_erroneous_requests"`

	// k8s allowed requests
	// Example: 2
	// Required: true
	K8sAllowedRequests *int64 `json:"k8s_allowed_requests"`

	// k8s denied requests
	// Example: 1
	// Required: true
	K8sDeniedRequests *int64 `json:"k8s_denied_requests"`

	// k8s erroneous requests
	// Example: 1
	// Required: true
	K8sErroneousRequests *int64 `json:"k8s_erroneous_requests"`

	// k8s ignored requests
	// Example: 1
	// Required: true
	K8sIgnoredRequests *int64 `json:"k8s_ignored_requests"`
}

// Validate validates this r e s t admission stats
func (m *RESTAdmissionStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJenkinsAllowedRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJenkinsDeniedRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJenkinsErroneousRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateK8sAllowedRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateK8sDeniedRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateK8sErroneousRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateK8sIgnoredRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAdmissionStats) validateJenkinsAllowedRequests(formats strfmt.Registry) error {

	if err := validate.Required("jenkins_allowed_requests", "body", m.JenkinsAllowedRequests); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmissionStats) validateJenkinsDeniedRequests(formats strfmt.Registry) error {

	if err := validate.Required("jenkins_denied_requests", "body", m.JenkinsDeniedRequests); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmissionStats) validateJenkinsErroneousRequests(formats strfmt.Registry) error {

	if err := validate.Required("jenkins_erroneous_requests", "body", m.JenkinsErroneousRequests); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmissionStats) validateK8sAllowedRequests(formats strfmt.Registry) error {

	if err := validate.Required("k8s_allowed_requests", "body", m.K8sAllowedRequests); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmissionStats) validateK8sDeniedRequests(formats strfmt.Registry) error {

	if err := validate.Required("k8s_denied_requests", "body", m.K8sDeniedRequests); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmissionStats) validateK8sErroneousRequests(formats strfmt.Registry) error {

	if err := validate.Required("k8s_erroneous_requests", "body", m.K8sErroneousRequests); err != nil {
		return err
	}

	return nil
}

func (m *RESTAdmissionStats) validateK8sIgnoredRequests(formats strfmt.Registry) error {

	if err := validate.Required("k8s_ignored_requests", "body", m.K8sIgnoredRequests); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this r e s t admission stats based on context it is used
func (m *RESTAdmissionStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTAdmissionStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAdmissionStats) UnmarshalBinary(b []byte) error {
	var res RESTAdmissionStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}