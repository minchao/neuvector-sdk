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

// NewGetV1LogActivityParams creates a new GetV1LogActivityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1LogActivityParams() *GetV1LogActivityParams {
	return &GetV1LogActivityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1LogActivityParamsWithTimeout creates a new GetV1LogActivityParams object
// with the ability to set a timeout on a request.
func NewGetV1LogActivityParamsWithTimeout(timeout time.Duration) *GetV1LogActivityParams {
	return &GetV1LogActivityParams{
		timeout: timeout,
	}
}

// NewGetV1LogActivityParamsWithContext creates a new GetV1LogActivityParams object
// with the ability to set a context for a request.
func NewGetV1LogActivityParamsWithContext(ctx context.Context) *GetV1LogActivityParams {
	return &GetV1LogActivityParams{
		Context: ctx,
	}
}

// NewGetV1LogActivityParamsWithHTTPClient creates a new GetV1LogActivityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1LogActivityParamsWithHTTPClient(client *http.Client) *GetV1LogActivityParams {
	return &GetV1LogActivityParams{
		HTTPClient: client,
	}
}

/*
GetV1LogActivityParams contains all the parameters to send to the API endpoint

	for the get v1 log activity operation.

	Typically these are written to a http.Request.
*/
type GetV1LogActivityParams struct {

	// XAuthToken.
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 log activity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogActivityParams) WithDefaults() *GetV1LogActivityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 log activity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogActivityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 log activity params
func (o *GetV1LogActivityParams) WithTimeout(timeout time.Duration) *GetV1LogActivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 log activity params
func (o *GetV1LogActivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 log activity params
func (o *GetV1LogActivityParams) WithContext(ctx context.Context) *GetV1LogActivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 log activity params
func (o *GetV1LogActivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 log activity params
func (o *GetV1LogActivityParams) WithHTTPClient(client *http.Client) *GetV1LogActivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 log activity params
func (o *GetV1LogActivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 log activity params
func (o *GetV1LogActivityParams) WithXAuthToken(xAuthToken string) *GetV1LogActivityParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 log activity params
func (o *GetV1LogActivityParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1LogActivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}