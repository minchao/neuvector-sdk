// Code generated by go-swagger; DO NOT EDIT.

package process

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

// NewGetV1ProcessProfileParams creates a new GetV1ProcessProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ProcessProfileParams() *GetV1ProcessProfileParams {
	return &GetV1ProcessProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ProcessProfileParamsWithTimeout creates a new GetV1ProcessProfileParams object
// with the ability to set a timeout on a request.
func NewGetV1ProcessProfileParamsWithTimeout(timeout time.Duration) *GetV1ProcessProfileParams {
	return &GetV1ProcessProfileParams{
		timeout: timeout,
	}
}

// NewGetV1ProcessProfileParamsWithContext creates a new GetV1ProcessProfileParams object
// with the ability to set a context for a request.
func NewGetV1ProcessProfileParamsWithContext(ctx context.Context) *GetV1ProcessProfileParams {
	return &GetV1ProcessProfileParams{
		Context: ctx,
	}
}

// NewGetV1ProcessProfileParamsWithHTTPClient creates a new GetV1ProcessProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ProcessProfileParamsWithHTTPClient(client *http.Client) *GetV1ProcessProfileParams {
	return &GetV1ProcessProfileParams{
		HTTPClient: client,
	}
}

/*
GetV1ProcessProfileParams contains all the parameters to send to the API endpoint

	for the get v1 process profile operation.

	Typically these are written to a http.Request.
*/
type GetV1ProcessProfileParams struct {

	// XAuthToken.
	XAuthToken string

	/* Scope.

	   When set to fed, returned fed process profiles. When set to local, returned local process profiles. If there is no query string 'scope', all process profiles will be returned.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 process profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ProcessProfileParams) WithDefaults() *GetV1ProcessProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 process profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ProcessProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 process profile params
func (o *GetV1ProcessProfileParams) WithTimeout(timeout time.Duration) *GetV1ProcessProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 process profile params
func (o *GetV1ProcessProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 process profile params
func (o *GetV1ProcessProfileParams) WithContext(ctx context.Context) *GetV1ProcessProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 process profile params
func (o *GetV1ProcessProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 process profile params
func (o *GetV1ProcessProfileParams) WithHTTPClient(client *http.Client) *GetV1ProcessProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 process profile params
func (o *GetV1ProcessProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 process profile params
func (o *GetV1ProcessProfileParams) WithXAuthToken(xAuthToken string) *GetV1ProcessProfileParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 process profile params
func (o *GetV1ProcessProfileParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithScope adds the scope to the get v1 process profile params
func (o *GetV1ProcessProfileParams) WithScope(scope *string) *GetV1ProcessProfileParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get v1 process profile params
func (o *GetV1ProcessProfileParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ProcessProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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