// Code generated by go-swagger; DO NOT EDIT.

package scan

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

// NewGetV1ScanHostIDParams creates a new GetV1ScanHostIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ScanHostIDParams() *GetV1ScanHostIDParams {
	return &GetV1ScanHostIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ScanHostIDParamsWithTimeout creates a new GetV1ScanHostIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ScanHostIDParamsWithTimeout(timeout time.Duration) *GetV1ScanHostIDParams {
	return &GetV1ScanHostIDParams{
		timeout: timeout,
	}
}

// NewGetV1ScanHostIDParamsWithContext creates a new GetV1ScanHostIDParams object
// with the ability to set a context for a request.
func NewGetV1ScanHostIDParamsWithContext(ctx context.Context) *GetV1ScanHostIDParams {
	return &GetV1ScanHostIDParams{
		Context: ctx,
	}
}

// NewGetV1ScanHostIDParamsWithHTTPClient creates a new GetV1ScanHostIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ScanHostIDParamsWithHTTPClient(client *http.Client) *GetV1ScanHostIDParams {
	return &GetV1ScanHostIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ScanHostIDParams contains all the parameters to send to the API endpoint

	for the get v1 scan host ID operation.

	Typically these are written to a http.Request.
*/
type GetV1ScanHostIDParams struct {

	// XAuthToken.
	XAuthToken string

	/* ID.

	   Host ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 scan host ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanHostIDParams) WithDefaults() *GetV1ScanHostIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 scan host ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanHostIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) WithTimeout(timeout time.Duration) *GetV1ScanHostIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) WithContext(ctx context.Context) *GetV1ScanHostIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) WithHTTPClient(client *http.Client) *GetV1ScanHostIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) WithXAuthToken(xAuthToken string) *GetV1ScanHostIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithID adds the id to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) WithID(id string) *GetV1ScanHostIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 scan host ID params
func (o *GetV1ScanHostIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ScanHostIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}