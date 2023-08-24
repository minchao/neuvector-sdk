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
)

// NewGetV1AdmissionRulesParams creates a new GetV1AdmissionRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AdmissionRulesParams() *GetV1AdmissionRulesParams {
	return &GetV1AdmissionRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AdmissionRulesParamsWithTimeout creates a new GetV1AdmissionRulesParams object
// with the ability to set a timeout on a request.
func NewGetV1AdmissionRulesParamsWithTimeout(timeout time.Duration) *GetV1AdmissionRulesParams {
	return &GetV1AdmissionRulesParams{
		timeout: timeout,
	}
}

// NewGetV1AdmissionRulesParamsWithContext creates a new GetV1AdmissionRulesParams object
// with the ability to set a context for a request.
func NewGetV1AdmissionRulesParamsWithContext(ctx context.Context) *GetV1AdmissionRulesParams {
	return &GetV1AdmissionRulesParams{
		Context: ctx,
	}
}

// NewGetV1AdmissionRulesParamsWithHTTPClient creates a new GetV1AdmissionRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AdmissionRulesParamsWithHTTPClient(client *http.Client) *GetV1AdmissionRulesParams {
	return &GetV1AdmissionRulesParams{
		HTTPClient: client,
	}
}

/*
GetV1AdmissionRulesParams contains all the parameters to send to the API endpoint

	for the get v1 admission rules operation.

	Typically these are written to a http.Request.
*/
type GetV1AdmissionRulesParams struct {

	/* Scope.

	   When set to fed, returned fed admission rules. When set to local, returned local admission rules. If there is no query string 'scope', all admission rules will be returned.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 admission rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AdmissionRulesParams) WithDefaults() *GetV1AdmissionRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 admission rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AdmissionRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) WithTimeout(timeout time.Duration) *GetV1AdmissionRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) WithContext(ctx context.Context) *GetV1AdmissionRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) WithHTTPClient(client *http.Client) *GetV1AdmissionRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) WithScope(scope *string) *GetV1AdmissionRulesParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get v1 admission rules params
func (o *GetV1AdmissionRulesParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AdmissionRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
