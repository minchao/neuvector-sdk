// Code generated by go-swagger; DO NOT EDIT.

package log

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

// NewGetV1LogSecurityParams creates a new GetV1LogSecurityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1LogSecurityParams() *GetV1LogSecurityParams {
	return &GetV1LogSecurityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1LogSecurityParamsWithTimeout creates a new GetV1LogSecurityParams object
// with the ability to set a timeout on a request.
func NewGetV1LogSecurityParamsWithTimeout(timeout time.Duration) *GetV1LogSecurityParams {
	return &GetV1LogSecurityParams{
		timeout: timeout,
	}
}

// NewGetV1LogSecurityParamsWithContext creates a new GetV1LogSecurityParams object
// with the ability to set a context for a request.
func NewGetV1LogSecurityParamsWithContext(ctx context.Context) *GetV1LogSecurityParams {
	return &GetV1LogSecurityParams{
		Context: ctx,
	}
}

// NewGetV1LogSecurityParamsWithHTTPClient creates a new GetV1LogSecurityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1LogSecurityParamsWithHTTPClient(client *http.Client) *GetV1LogSecurityParams {
	return &GetV1LogSecurityParams{
		HTTPClient: client,
	}
}

/*
GetV1LogSecurityParams contains all the parameters to send to the API endpoint

	for the get v1 log security operation.

	Typically these are written to a http.Request.
*/
type GetV1LogSecurityParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 log security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogSecurityParams) WithDefaults() *GetV1LogSecurityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 log security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogSecurityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 log security params
func (o *GetV1LogSecurityParams) WithTimeout(timeout time.Duration) *GetV1LogSecurityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 log security params
func (o *GetV1LogSecurityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 log security params
func (o *GetV1LogSecurityParams) WithContext(ctx context.Context) *GetV1LogSecurityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 log security params
func (o *GetV1LogSecurityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 log security params
func (o *GetV1LogSecurityParams) WithHTTPClient(client *http.Client) *GetV1LogSecurityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 log security params
func (o *GetV1LogSecurityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1LogSecurityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
