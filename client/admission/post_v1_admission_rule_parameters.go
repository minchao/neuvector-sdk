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

// NewPostV1AdmissionRuleParams creates a new PostV1AdmissionRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1AdmissionRuleParams() *PostV1AdmissionRuleParams {
	return &PostV1AdmissionRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1AdmissionRuleParamsWithTimeout creates a new PostV1AdmissionRuleParams object
// with the ability to set a timeout on a request.
func NewPostV1AdmissionRuleParamsWithTimeout(timeout time.Duration) *PostV1AdmissionRuleParams {
	return &PostV1AdmissionRuleParams{
		timeout: timeout,
	}
}

// NewPostV1AdmissionRuleParamsWithContext creates a new PostV1AdmissionRuleParams object
// with the ability to set a context for a request.
func NewPostV1AdmissionRuleParamsWithContext(ctx context.Context) *PostV1AdmissionRuleParams {
	return &PostV1AdmissionRuleParams{
		Context: ctx,
	}
}

// NewPostV1AdmissionRuleParamsWithHTTPClient creates a new PostV1AdmissionRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1AdmissionRuleParamsWithHTTPClient(client *http.Client) *PostV1AdmissionRuleParams {
	return &PostV1AdmissionRuleParams{
		HTTPClient: client,
	}
}

/*
PostV1AdmissionRuleParams contains all the parameters to send to the API endpoint

	for the post v1 admission rule operation.

	Typically these are written to a http.Request.
*/
type PostV1AdmissionRuleParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Admission rule data
	*/
	Body *models.RESTAdmissionRuleConfigData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 admission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AdmissionRuleParams) WithDefaults() *PostV1AdmissionRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 admission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AdmissionRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) WithTimeout(timeout time.Duration) *PostV1AdmissionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) WithContext(ctx context.Context) *PostV1AdmissionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) WithHTTPClient(client *http.Client) *PostV1AdmissionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) WithXAuthToken(xAuthToken string) *PostV1AdmissionRuleParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) WithBody(body *models.RESTAdmissionRuleConfigData) *PostV1AdmissionRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 admission rule params
func (o *PostV1AdmissionRuleParams) SetBody(body *models.RESTAdmissionRuleConfigData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1AdmissionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}
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
