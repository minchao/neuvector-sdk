// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGetV1GroupNameParams creates a new GetV1GroupNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1GroupNameParams() *GetV1GroupNameParams {
	return &GetV1GroupNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1GroupNameParamsWithTimeout creates a new GetV1GroupNameParams object
// with the ability to set a timeout on a request.
func NewGetV1GroupNameParamsWithTimeout(timeout time.Duration) *GetV1GroupNameParams {
	return &GetV1GroupNameParams{
		timeout: timeout,
	}
}

// NewGetV1GroupNameParamsWithContext creates a new GetV1GroupNameParams object
// with the ability to set a context for a request.
func NewGetV1GroupNameParamsWithContext(ctx context.Context) *GetV1GroupNameParams {
	return &GetV1GroupNameParams{
		Context: ctx,
	}
}

// NewGetV1GroupNameParamsWithHTTPClient creates a new GetV1GroupNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1GroupNameParamsWithHTTPClient(client *http.Client) *GetV1GroupNameParams {
	return &GetV1GroupNameParams{
		HTTPClient: client,
	}
}

/*
GetV1GroupNameParams contains all the parameters to send to the API endpoint

	for the get v1 group name operation.

	Typically these are written to a http.Request.
*/
type GetV1GroupNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   Group name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1GroupNameParams) WithDefaults() *GetV1GroupNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1GroupNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 group name params
func (o *GetV1GroupNameParams) WithTimeout(timeout time.Duration) *GetV1GroupNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 group name params
func (o *GetV1GroupNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 group name params
func (o *GetV1GroupNameParams) WithContext(ctx context.Context) *GetV1GroupNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 group name params
func (o *GetV1GroupNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 group name params
func (o *GetV1GroupNameParams) WithHTTPClient(client *http.Client) *GetV1GroupNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 group name params
func (o *GetV1GroupNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 group name params
func (o *GetV1GroupNameParams) WithXAuthToken(xAuthToken string) *GetV1GroupNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 group name params
func (o *GetV1GroupNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the get v1 group name params
func (o *GetV1GroupNameParams) WithName(name string) *GetV1GroupNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 group name params
func (o *GetV1GroupNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1GroupNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
