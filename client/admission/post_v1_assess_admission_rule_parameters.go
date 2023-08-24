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

// NewPostV1AssessAdmissionRuleParams creates a new PostV1AssessAdmissionRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1AssessAdmissionRuleParams() *PostV1AssessAdmissionRuleParams {
	return &PostV1AssessAdmissionRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1AssessAdmissionRuleParamsWithTimeout creates a new PostV1AssessAdmissionRuleParams object
// with the ability to set a timeout on a request.
func NewPostV1AssessAdmissionRuleParamsWithTimeout(timeout time.Duration) *PostV1AssessAdmissionRuleParams {
	return &PostV1AssessAdmissionRuleParams{
		timeout: timeout,
	}
}

// NewPostV1AssessAdmissionRuleParamsWithContext creates a new PostV1AssessAdmissionRuleParams object
// with the ability to set a context for a request.
func NewPostV1AssessAdmissionRuleParamsWithContext(ctx context.Context) *PostV1AssessAdmissionRuleParams {
	return &PostV1AssessAdmissionRuleParams{
		Context: ctx,
	}
}

// NewPostV1AssessAdmissionRuleParamsWithHTTPClient creates a new PostV1AssessAdmissionRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1AssessAdmissionRuleParamsWithHTTPClient(client *http.Client) *PostV1AssessAdmissionRuleParams {
	return &PostV1AssessAdmissionRuleParams{
		HTTPClient: client,
	}
}

/*
PostV1AssessAdmissionRuleParams contains all the parameters to send to the API endpoint

	for the post v1 assess admission rule operation.

	Typically these are written to a http.Request.
*/
type PostV1AssessAdmissionRuleParams struct {

	/* Body.

	   Admission rule data
	*/
	Body *models.RESTAdmissionRuleConfigData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 assess admission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AssessAdmissionRuleParams) WithDefaults() *PostV1AssessAdmissionRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 assess admission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AssessAdmissionRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) WithTimeout(timeout time.Duration) *PostV1AssessAdmissionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) WithContext(ctx context.Context) *PostV1AssessAdmissionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) WithHTTPClient(client *http.Client) *PostV1AssessAdmissionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) WithBody(body *models.RESTAdmissionRuleConfigData) *PostV1AssessAdmissionRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 assess admission rule params
func (o *PostV1AssessAdmissionRuleParams) SetBody(body *models.RESTAdmissionRuleConfigData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1AssessAdmissionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
