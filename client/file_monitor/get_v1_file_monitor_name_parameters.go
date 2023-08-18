// Code generated by go-swagger; DO NOT EDIT.

package file_monitor

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

// NewGetV1FileMonitorNameParams creates a new GetV1FileMonitorNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1FileMonitorNameParams() *GetV1FileMonitorNameParams {
	return &GetV1FileMonitorNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1FileMonitorNameParamsWithTimeout creates a new GetV1FileMonitorNameParams object
// with the ability to set a timeout on a request.
func NewGetV1FileMonitorNameParamsWithTimeout(timeout time.Duration) *GetV1FileMonitorNameParams {
	return &GetV1FileMonitorNameParams{
		timeout: timeout,
	}
}

// NewGetV1FileMonitorNameParamsWithContext creates a new GetV1FileMonitorNameParams object
// with the ability to set a context for a request.
func NewGetV1FileMonitorNameParamsWithContext(ctx context.Context) *GetV1FileMonitorNameParams {
	return &GetV1FileMonitorNameParams{
		Context: ctx,
	}
}

// NewGetV1FileMonitorNameParamsWithHTTPClient creates a new GetV1FileMonitorNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1FileMonitorNameParamsWithHTTPClient(client *http.Client) *GetV1FileMonitorNameParams {
	return &GetV1FileMonitorNameParams{
		HTTPClient: client,
	}
}

/*
GetV1FileMonitorNameParams contains all the parameters to send to the API endpoint

	for the get v1 file monitor name operation.

	Typically these are written to a http.Request.
*/
type GetV1FileMonitorNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   File monitor name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 file monitor name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FileMonitorNameParams) WithDefaults() *GetV1FileMonitorNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 file monitor name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FileMonitorNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) WithTimeout(timeout time.Duration) *GetV1FileMonitorNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) WithContext(ctx context.Context) *GetV1FileMonitorNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) WithHTTPClient(client *http.Client) *GetV1FileMonitorNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) WithXAuthToken(xAuthToken string) *GetV1FileMonitorNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) WithName(name string) *GetV1FileMonitorNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 file monitor name params
func (o *GetV1FileMonitorNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1FileMonitorNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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