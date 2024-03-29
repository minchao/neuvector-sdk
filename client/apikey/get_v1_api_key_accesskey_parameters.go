// Code generated by go-swagger; DO NOT EDIT.

package apikey

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

// NewGetV1APIKeyAccesskeyParams creates a new GetV1APIKeyAccesskeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1APIKeyAccesskeyParams() *GetV1APIKeyAccesskeyParams {
	return &GetV1APIKeyAccesskeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1APIKeyAccesskeyParamsWithTimeout creates a new GetV1APIKeyAccesskeyParams object
// with the ability to set a timeout on a request.
func NewGetV1APIKeyAccesskeyParamsWithTimeout(timeout time.Duration) *GetV1APIKeyAccesskeyParams {
	return &GetV1APIKeyAccesskeyParams{
		timeout: timeout,
	}
}

// NewGetV1APIKeyAccesskeyParamsWithContext creates a new GetV1APIKeyAccesskeyParams object
// with the ability to set a context for a request.
func NewGetV1APIKeyAccesskeyParamsWithContext(ctx context.Context) *GetV1APIKeyAccesskeyParams {
	return &GetV1APIKeyAccesskeyParams{
		Context: ctx,
	}
}

// NewGetV1APIKeyAccesskeyParamsWithHTTPClient creates a new GetV1APIKeyAccesskeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1APIKeyAccesskeyParamsWithHTTPClient(client *http.Client) *GetV1APIKeyAccesskeyParams {
	return &GetV1APIKeyAccesskeyParams{
		HTTPClient: client,
	}
}

/*
GetV1APIKeyAccesskeyParams contains all the parameters to send to the API endpoint

	for the get v1 API key accesskey operation.

	Typically these are written to a http.Request.
*/
type GetV1APIKeyAccesskeyParams struct {

	/* Accesskey.

	   Apikey access key
	*/
	Accesskey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 API key accesskey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1APIKeyAccesskeyParams) WithDefaults() *GetV1APIKeyAccesskeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 API key accesskey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1APIKeyAccesskeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) WithTimeout(timeout time.Duration) *GetV1APIKeyAccesskeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) WithContext(ctx context.Context) *GetV1APIKeyAccesskeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) WithHTTPClient(client *http.Client) *GetV1APIKeyAccesskeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccesskey adds the accesskey to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) WithAccesskey(accesskey string) *GetV1APIKeyAccesskeyParams {
	o.SetAccesskey(accesskey)
	return o
}

// SetAccesskey adds the accesskey to the get v1 API key accesskey params
func (o *GetV1APIKeyAccesskeyParams) SetAccesskey(accesskey string) {
	o.Accesskey = accesskey
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1APIKeyAccesskeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accesskey
	if err := r.SetPathParam("accesskey", o.Accesskey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
