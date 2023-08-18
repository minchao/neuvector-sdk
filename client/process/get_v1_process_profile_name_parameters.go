// Code generated by go-swagger; DO NOT EDIT.

package process

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

// NewGetV1ProcessProfileNameParams creates a new GetV1ProcessProfileNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ProcessProfileNameParams() *GetV1ProcessProfileNameParams {
	return &GetV1ProcessProfileNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ProcessProfileNameParamsWithTimeout creates a new GetV1ProcessProfileNameParams object
// with the ability to set a timeout on a request.
func NewGetV1ProcessProfileNameParamsWithTimeout(timeout time.Duration) *GetV1ProcessProfileNameParams {
	return &GetV1ProcessProfileNameParams{
		timeout: timeout,
	}
}

// NewGetV1ProcessProfileNameParamsWithContext creates a new GetV1ProcessProfileNameParams object
// with the ability to set a context for a request.
func NewGetV1ProcessProfileNameParamsWithContext(ctx context.Context) *GetV1ProcessProfileNameParams {
	return &GetV1ProcessProfileNameParams{
		Context: ctx,
	}
}

// NewGetV1ProcessProfileNameParamsWithHTTPClient creates a new GetV1ProcessProfileNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ProcessProfileNameParamsWithHTTPClient(client *http.Client) *GetV1ProcessProfileNameParams {
	return &GetV1ProcessProfileNameParams{
		HTTPClient: client,
	}
}

/*
GetV1ProcessProfileNameParams contains all the parameters to send to the API endpoint

	for the get v1 process profile name operation.

	Typically these are written to a http.Request.
*/
type GetV1ProcessProfileNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   Process profile name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 process profile name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ProcessProfileNameParams) WithDefaults() *GetV1ProcessProfileNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 process profile name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ProcessProfileNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) WithTimeout(timeout time.Duration) *GetV1ProcessProfileNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) WithContext(ctx context.Context) *GetV1ProcessProfileNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) WithHTTPClient(client *http.Client) *GetV1ProcessProfileNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) WithXAuthToken(xAuthToken string) *GetV1ProcessProfileNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) WithName(name string) *GetV1ProcessProfileNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 process profile name params
func (o *GetV1ProcessProfileNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ProcessProfileNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
