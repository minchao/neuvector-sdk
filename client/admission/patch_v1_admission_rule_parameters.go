// Code generated by go-swagger; DO NOT EDIT.

package admission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// NewPatchV1AdmissionRuleParams creates a new PatchV1AdmissionRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1AdmissionRuleParams() *PatchV1AdmissionRuleParams {
	return &PatchV1AdmissionRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1AdmissionRuleParamsWithTimeout creates a new PatchV1AdmissionRuleParams object
// with the ability to set a timeout on a request.
func NewPatchV1AdmissionRuleParamsWithTimeout(timeout time.Duration) *PatchV1AdmissionRuleParams {
	return &PatchV1AdmissionRuleParams{
		timeout: timeout,
	}
}

// NewPatchV1AdmissionRuleParamsWithContext creates a new PatchV1AdmissionRuleParams object
// with the ability to set a context for a request.
func NewPatchV1AdmissionRuleParamsWithContext(ctx context.Context) *PatchV1AdmissionRuleParams {
	return &PatchV1AdmissionRuleParams{
		Context: ctx,
	}
}

// NewPatchV1AdmissionRuleParamsWithHTTPClient creates a new PatchV1AdmissionRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1AdmissionRuleParamsWithHTTPClient(client *http.Client) *PatchV1AdmissionRuleParams {
	return &PatchV1AdmissionRuleParams{
		HTTPClient: client,
	}
}

/*
PatchV1AdmissionRuleParams contains all the parameters to send to the API endpoint

	for the patch v1 admission rule operation.

	Typically these are written to a http.Request.
*/
type PatchV1AdmissionRuleParams struct {

	/* Body.

	   Admission rule data
	*/
	Body *models.RESTAdmissionRuleConfigData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 admission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1AdmissionRuleParams) WithDefaults() *PatchV1AdmissionRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 admission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1AdmissionRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) WithTimeout(timeout time.Duration) *PatchV1AdmissionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) WithContext(ctx context.Context) *PatchV1AdmissionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) WithHTTPClient(client *http.Client) *PatchV1AdmissionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) WithBody(body *models.RESTAdmissionRuleConfigData) *PatchV1AdmissionRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 admission rule params
func (o *PatchV1AdmissionRuleParams) SetBody(body *models.RESTAdmissionRuleConfigData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1AdmissionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
