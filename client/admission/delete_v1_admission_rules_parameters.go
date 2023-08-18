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

// NewDeleteV1AdmissionRulesParams creates a new DeleteV1AdmissionRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1AdmissionRulesParams() *DeleteV1AdmissionRulesParams {
	return &DeleteV1AdmissionRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1AdmissionRulesParamsWithTimeout creates a new DeleteV1AdmissionRulesParams object
// with the ability to set a timeout on a request.
func NewDeleteV1AdmissionRulesParamsWithTimeout(timeout time.Duration) *DeleteV1AdmissionRulesParams {
	return &DeleteV1AdmissionRulesParams{
		timeout: timeout,
	}
}

// NewDeleteV1AdmissionRulesParamsWithContext creates a new DeleteV1AdmissionRulesParams object
// with the ability to set a context for a request.
func NewDeleteV1AdmissionRulesParamsWithContext(ctx context.Context) *DeleteV1AdmissionRulesParams {
	return &DeleteV1AdmissionRulesParams{
		Context: ctx,
	}
}

// NewDeleteV1AdmissionRulesParamsWithHTTPClient creates a new DeleteV1AdmissionRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1AdmissionRulesParamsWithHTTPClient(client *http.Client) *DeleteV1AdmissionRulesParams {
	return &DeleteV1AdmissionRulesParams{
		HTTPClient: client,
	}
}

/*
DeleteV1AdmissionRulesParams contains all the parameters to send to the API endpoint

	for the delete v1 admission rules operation.

	Typically these are written to a http.Request.
*/
type DeleteV1AdmissionRulesParams struct {

	// XAuthToken.
	XAuthToken string

	/* Scope.

	   When set to fed, fed admission rules get removed. When set to local or no query string, local admission rules will be removed.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 admission rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1AdmissionRulesParams) WithDefaults() *DeleteV1AdmissionRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 admission rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1AdmissionRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) WithTimeout(timeout time.Duration) *DeleteV1AdmissionRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) WithContext(ctx context.Context) *DeleteV1AdmissionRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) WithHTTPClient(client *http.Client) *DeleteV1AdmissionRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) WithXAuthToken(xAuthToken string) *DeleteV1AdmissionRulesParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithScope adds the scope to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) WithScope(scope *string) *DeleteV1AdmissionRulesParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete v1 admission rules params
func (o *DeleteV1AdmissionRulesParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1AdmissionRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

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