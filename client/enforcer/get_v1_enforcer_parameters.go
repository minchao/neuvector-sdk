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

// NewGetV1EnforcerParams creates a new GetV1EnforcerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1EnforcerParams() *GetV1EnforcerParams {
	return &GetV1EnforcerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1EnforcerParamsWithTimeout creates a new GetV1EnforcerParams object
// with the ability to set a timeout on a request.
func NewGetV1EnforcerParamsWithTimeout(timeout time.Duration) *GetV1EnforcerParams {
	return &GetV1EnforcerParams{
		timeout: timeout,
	}
}

// NewGetV1EnforcerParamsWithContext creates a new GetV1EnforcerParams object
// with the ability to set a context for a request.
func NewGetV1EnforcerParamsWithContext(ctx context.Context) *GetV1EnforcerParams {
	return &GetV1EnforcerParams{
		Context: ctx,
	}
}

// NewGetV1EnforcerParamsWithHTTPClient creates a new GetV1EnforcerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1EnforcerParamsWithHTTPClient(client *http.Client) *GetV1EnforcerParams {
	return &GetV1EnforcerParams{
		HTTPClient: client,
	}
}

/*
GetV1EnforcerParams contains all the parameters to send to the API endpoint

	for the get v1 enforcer operation.

	Typically these are written to a http.Request.
*/
type GetV1EnforcerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 enforcer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1EnforcerParams) WithDefaults() *GetV1EnforcerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 enforcer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1EnforcerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 enforcer params
func (o *GetV1EnforcerParams) WithTimeout(timeout time.Duration) *GetV1EnforcerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 enforcer params
func (o *GetV1EnforcerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 enforcer params
func (o *GetV1EnforcerParams) WithContext(ctx context.Context) *GetV1EnforcerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 enforcer params
func (o *GetV1EnforcerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 enforcer params
func (o *GetV1EnforcerParams) WithHTTPClient(client *http.Client) *GetV1EnforcerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 enforcer params
func (o *GetV1EnforcerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1EnforcerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
