// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetV1ServerNameParams creates a new GetV1ServerNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ServerNameParams() *GetV1ServerNameParams {
	return &GetV1ServerNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ServerNameParamsWithTimeout creates a new GetV1ServerNameParams object
// with the ability to set a timeout on a request.
func NewGetV1ServerNameParamsWithTimeout(timeout time.Duration) *GetV1ServerNameParams {
	return &GetV1ServerNameParams{
		timeout: timeout,
	}
}

// NewGetV1ServerNameParamsWithContext creates a new GetV1ServerNameParams object
// with the ability to set a context for a request.
func NewGetV1ServerNameParamsWithContext(ctx context.Context) *GetV1ServerNameParams {
	return &GetV1ServerNameParams{
		Context: ctx,
	}
}

// NewGetV1ServerNameParamsWithHTTPClient creates a new GetV1ServerNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ServerNameParamsWithHTTPClient(client *http.Client) *GetV1ServerNameParams {
	return &GetV1ServerNameParams{
		HTTPClient: client,
	}
}

/*
GetV1ServerNameParams contains all the parameters to send to the API endpoint

	for the get v1 server name operation.

	Typically these are written to a http.Request.
*/
type GetV1ServerNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   Name of the server
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 server name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ServerNameParams) WithDefaults() *GetV1ServerNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 server name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ServerNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 server name params
func (o *GetV1ServerNameParams) WithTimeout(timeout time.Duration) *GetV1ServerNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 server name params
func (o *GetV1ServerNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 server name params
func (o *GetV1ServerNameParams) WithContext(ctx context.Context) *GetV1ServerNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 server name params
func (o *GetV1ServerNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 server name params
func (o *GetV1ServerNameParams) WithHTTPClient(client *http.Client) *GetV1ServerNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 server name params
func (o *GetV1ServerNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 server name params
func (o *GetV1ServerNameParams) WithXAuthToken(xAuthToken string) *GetV1ServerNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 server name params
func (o *GetV1ServerNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the get v1 server name params
func (o *GetV1ServerNameParams) WithName(name string) *GetV1ServerNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 server name params
func (o *GetV1ServerNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ServerNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}