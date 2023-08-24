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

// NewGetV1LogViolationParams creates a new GetV1LogViolationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1LogViolationParams() *GetV1LogViolationParams {
	return &GetV1LogViolationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1LogViolationParamsWithTimeout creates a new GetV1LogViolationParams object
// with the ability to set a timeout on a request.
func NewGetV1LogViolationParamsWithTimeout(timeout time.Duration) *GetV1LogViolationParams {
	return &GetV1LogViolationParams{
		timeout: timeout,
	}
}

// NewGetV1LogViolationParamsWithContext creates a new GetV1LogViolationParams object
// with the ability to set a context for a request.
func NewGetV1LogViolationParamsWithContext(ctx context.Context) *GetV1LogViolationParams {
	return &GetV1LogViolationParams{
		Context: ctx,
	}
}

// NewGetV1LogViolationParamsWithHTTPClient creates a new GetV1LogViolationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1LogViolationParamsWithHTTPClient(client *http.Client) *GetV1LogViolationParams {
	return &GetV1LogViolationParams{
		HTTPClient: client,
	}
}

/*
GetV1LogViolationParams contains all the parameters to send to the API endpoint

	for the get v1 log violation operation.

	Typically these are written to a http.Request.
*/
type GetV1LogViolationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 log violation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogViolationParams) WithDefaults() *GetV1LogViolationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 log violation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogViolationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 log violation params
func (o *GetV1LogViolationParams) WithTimeout(timeout time.Duration) *GetV1LogViolationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 log violation params
func (o *GetV1LogViolationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 log violation params
func (o *GetV1LogViolationParams) WithContext(ctx context.Context) *GetV1LogViolationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 log violation params
func (o *GetV1LogViolationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 log violation params
func (o *GetV1LogViolationParams) WithHTTPClient(client *http.Client) *GetV1LogViolationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 log violation params
func (o *GetV1LogViolationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1LogViolationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
