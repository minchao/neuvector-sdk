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

// NewGetV1ServerNameUserParams creates a new GetV1ServerNameUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ServerNameUserParams() *GetV1ServerNameUserParams {
	return &GetV1ServerNameUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ServerNameUserParamsWithTimeout creates a new GetV1ServerNameUserParams object
// with the ability to set a timeout on a request.
func NewGetV1ServerNameUserParamsWithTimeout(timeout time.Duration) *GetV1ServerNameUserParams {
	return &GetV1ServerNameUserParams{
		timeout: timeout,
	}
}

// NewGetV1ServerNameUserParamsWithContext creates a new GetV1ServerNameUserParams object
// with the ability to set a context for a request.
func NewGetV1ServerNameUserParamsWithContext(ctx context.Context) *GetV1ServerNameUserParams {
	return &GetV1ServerNameUserParams{
		Context: ctx,
	}
}

// NewGetV1ServerNameUserParamsWithHTTPClient creates a new GetV1ServerNameUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ServerNameUserParamsWithHTTPClient(client *http.Client) *GetV1ServerNameUserParams {
	return &GetV1ServerNameUserParams{
		HTTPClient: client,
	}
}

/*
GetV1ServerNameUserParams contains all the parameters to send to the API endpoint

	for the get v1 server name user operation.

	Typically these are written to a http.Request.
*/
type GetV1ServerNameUserParams struct {

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

// WithDefaults hydrates default values in the get v1 server name user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ServerNameUserParams) WithDefaults() *GetV1ServerNameUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 server name user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ServerNameUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 server name user params
func (o *GetV1ServerNameUserParams) WithTimeout(timeout time.Duration) *GetV1ServerNameUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 server name user params
func (o *GetV1ServerNameUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 server name user params
func (o *GetV1ServerNameUserParams) WithContext(ctx context.Context) *GetV1ServerNameUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 server name user params
func (o *GetV1ServerNameUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 server name user params
func (o *GetV1ServerNameUserParams) WithHTTPClient(client *http.Client) *GetV1ServerNameUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 server name user params
func (o *GetV1ServerNameUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 server name user params
func (o *GetV1ServerNameUserParams) WithXAuthToken(xAuthToken string) *GetV1ServerNameUserParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 server name user params
func (o *GetV1ServerNameUserParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the get v1 server name user params
func (o *GetV1ServerNameUserParams) WithName(name string) *GetV1ServerNameUserParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 server name user params
func (o *GetV1ServerNameUserParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ServerNameUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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