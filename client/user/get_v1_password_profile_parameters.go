// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetV1PasswordProfileParams creates a new GetV1PasswordProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1PasswordProfileParams() *GetV1PasswordProfileParams {
	return &GetV1PasswordProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PasswordProfileParamsWithTimeout creates a new GetV1PasswordProfileParams object
// with the ability to set a timeout on a request.
func NewGetV1PasswordProfileParamsWithTimeout(timeout time.Duration) *GetV1PasswordProfileParams {
	return &GetV1PasswordProfileParams{
		timeout: timeout,
	}
}

// NewGetV1PasswordProfileParamsWithContext creates a new GetV1PasswordProfileParams object
// with the ability to set a context for a request.
func NewGetV1PasswordProfileParamsWithContext(ctx context.Context) *GetV1PasswordProfileParams {
	return &GetV1PasswordProfileParams{
		Context: ctx,
	}
}

// NewGetV1PasswordProfileParamsWithHTTPClient creates a new GetV1PasswordProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1PasswordProfileParamsWithHTTPClient(client *http.Client) *GetV1PasswordProfileParams {
	return &GetV1PasswordProfileParams{
		HTTPClient: client,
	}
}

/*
GetV1PasswordProfileParams contains all the parameters to send to the API endpoint

	for the get v1 password profile operation.

	Typically these are written to a http.Request.
*/
type GetV1PasswordProfileParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 password profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PasswordProfileParams) WithDefaults() *GetV1PasswordProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 password profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PasswordProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 password profile params
func (o *GetV1PasswordProfileParams) WithTimeout(timeout time.Duration) *GetV1PasswordProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 password profile params
func (o *GetV1PasswordProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 password profile params
func (o *GetV1PasswordProfileParams) WithContext(ctx context.Context) *GetV1PasswordProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 password profile params
func (o *GetV1PasswordProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 password profile params
func (o *GetV1PasswordProfileParams) WithHTTPClient(client *http.Client) *GetV1PasswordProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 password profile params
func (o *GetV1PasswordProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PasswordProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
