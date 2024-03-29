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

// NewGetV1UserParams creates a new GetV1UserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1UserParams() *GetV1UserParams {
	return &GetV1UserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1UserParamsWithTimeout creates a new GetV1UserParams object
// with the ability to set a timeout on a request.
func NewGetV1UserParamsWithTimeout(timeout time.Duration) *GetV1UserParams {
	return &GetV1UserParams{
		timeout: timeout,
	}
}

// NewGetV1UserParamsWithContext creates a new GetV1UserParams object
// with the ability to set a context for a request.
func NewGetV1UserParamsWithContext(ctx context.Context) *GetV1UserParams {
	return &GetV1UserParams{
		Context: ctx,
	}
}

// NewGetV1UserParamsWithHTTPClient creates a new GetV1UserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1UserParamsWithHTTPClient(client *http.Client) *GetV1UserParams {
	return &GetV1UserParams{
		HTTPClient: client,
	}
}

/*
GetV1UserParams contains all the parameters to send to the API endpoint

	for the get v1 user operation.

	Typically these are written to a http.Request.
*/
type GetV1UserParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UserParams) WithDefaults() *GetV1UserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 user params
func (o *GetV1UserParams) WithTimeout(timeout time.Duration) *GetV1UserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 user params
func (o *GetV1UserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 user params
func (o *GetV1UserParams) WithContext(ctx context.Context) *GetV1UserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 user params
func (o *GetV1UserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 user params
func (o *GetV1UserParams) WithHTTPClient(client *http.Client) *GetV1UserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 user params
func (o *GetV1UserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1UserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
