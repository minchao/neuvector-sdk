// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetV1SystemConfigParams creates a new GetV1SystemConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SystemConfigParams() *GetV1SystemConfigParams {
	return &GetV1SystemConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SystemConfigParamsWithTimeout creates a new GetV1SystemConfigParams object
// with the ability to set a timeout on a request.
func NewGetV1SystemConfigParamsWithTimeout(timeout time.Duration) *GetV1SystemConfigParams {
	return &GetV1SystemConfigParams{
		timeout: timeout,
	}
}

// NewGetV1SystemConfigParamsWithContext creates a new GetV1SystemConfigParams object
// with the ability to set a context for a request.
func NewGetV1SystemConfigParamsWithContext(ctx context.Context) *GetV1SystemConfigParams {
	return &GetV1SystemConfigParams{
		Context: ctx,
	}
}

// NewGetV1SystemConfigParamsWithHTTPClient creates a new GetV1SystemConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SystemConfigParamsWithHTTPClient(client *http.Client) *GetV1SystemConfigParams {
	return &GetV1SystemConfigParams{
		HTTPClient: client,
	}
}

/*
GetV1SystemConfigParams contains all the parameters to send to the API endpoint

	for the get v1 system config operation.

	Typically these are written to a http.Request.
*/
type GetV1SystemConfigParams struct {

	// XAuthToken.
	XAuthToken string

	/* Scope.

	   When set the scope to be fed, it will return the fed system configures. When set the scope to be local, it will return the local system configures. If there is no query string 'scope', it will return all system configures.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 system config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SystemConfigParams) WithDefaults() *GetV1SystemConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 system config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SystemConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 system config params
func (o *GetV1SystemConfigParams) WithTimeout(timeout time.Duration) *GetV1SystemConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 system config params
func (o *GetV1SystemConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 system config params
func (o *GetV1SystemConfigParams) WithContext(ctx context.Context) *GetV1SystemConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 system config params
func (o *GetV1SystemConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 system config params
func (o *GetV1SystemConfigParams) WithHTTPClient(client *http.Client) *GetV1SystemConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 system config params
func (o *GetV1SystemConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 system config params
func (o *GetV1SystemConfigParams) WithXAuthToken(xAuthToken string) *GetV1SystemConfigParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 system config params
func (o *GetV1SystemConfigParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithScope adds the scope to the get v1 system config params
func (o *GetV1SystemConfigParams) WithScope(scope *string) *GetV1SystemConfigParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get v1 system config params
func (o *GetV1SystemConfigParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SystemConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
