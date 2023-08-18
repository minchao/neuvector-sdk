// Code generated by go-swagger; DO NOT EDIT.

package enforcer

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

// NewGetV1EnforcerIDConfigParams creates a new GetV1EnforcerIDConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1EnforcerIDConfigParams() *GetV1EnforcerIDConfigParams {
	return &GetV1EnforcerIDConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1EnforcerIDConfigParamsWithTimeout creates a new GetV1EnforcerIDConfigParams object
// with the ability to set a timeout on a request.
func NewGetV1EnforcerIDConfigParamsWithTimeout(timeout time.Duration) *GetV1EnforcerIDConfigParams {
	return &GetV1EnforcerIDConfigParams{
		timeout: timeout,
	}
}

// NewGetV1EnforcerIDConfigParamsWithContext creates a new GetV1EnforcerIDConfigParams object
// with the ability to set a context for a request.
func NewGetV1EnforcerIDConfigParamsWithContext(ctx context.Context) *GetV1EnforcerIDConfigParams {
	return &GetV1EnforcerIDConfigParams{
		Context: ctx,
	}
}

// NewGetV1EnforcerIDConfigParamsWithHTTPClient creates a new GetV1EnforcerIDConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1EnforcerIDConfigParamsWithHTTPClient(client *http.Client) *GetV1EnforcerIDConfigParams {
	return &GetV1EnforcerIDConfigParams{
		HTTPClient: client,
	}
}

/*
GetV1EnforcerIDConfigParams contains all the parameters to send to the API endpoint

	for the get v1 enforcer ID config operation.

	Typically these are written to a http.Request.
*/
type GetV1EnforcerIDConfigParams struct {

	// XAuthToken.
	XAuthToken string

	/* ID.

	   Enforcer ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 enforcer ID config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1EnforcerIDConfigParams) WithDefaults() *GetV1EnforcerIDConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 enforcer ID config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1EnforcerIDConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) WithTimeout(timeout time.Duration) *GetV1EnforcerIDConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) WithContext(ctx context.Context) *GetV1EnforcerIDConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) WithHTTPClient(client *http.Client) *GetV1EnforcerIDConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) WithXAuthToken(xAuthToken string) *GetV1EnforcerIDConfigParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithID adds the id to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) WithID(id string) *GetV1EnforcerIDConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 enforcer ID config params
func (o *GetV1EnforcerIDConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1EnforcerIDConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
