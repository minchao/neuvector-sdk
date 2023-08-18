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

// NewGetV1LogEventParams creates a new GetV1LogEventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1LogEventParams() *GetV1LogEventParams {
	return &GetV1LogEventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1LogEventParamsWithTimeout creates a new GetV1LogEventParams object
// with the ability to set a timeout on a request.
func NewGetV1LogEventParamsWithTimeout(timeout time.Duration) *GetV1LogEventParams {
	return &GetV1LogEventParams{
		timeout: timeout,
	}
}

// NewGetV1LogEventParamsWithContext creates a new GetV1LogEventParams object
// with the ability to set a context for a request.
func NewGetV1LogEventParamsWithContext(ctx context.Context) *GetV1LogEventParams {
	return &GetV1LogEventParams{
		Context: ctx,
	}
}

// NewGetV1LogEventParamsWithHTTPClient creates a new GetV1LogEventParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1LogEventParamsWithHTTPClient(client *http.Client) *GetV1LogEventParams {
	return &GetV1LogEventParams{
		HTTPClient: client,
	}
}

/*
GetV1LogEventParams contains all the parameters to send to the API endpoint

	for the get v1 log event operation.

	Typically these are written to a http.Request.
*/
type GetV1LogEventParams struct {

	// XAuthToken.
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 log event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogEventParams) WithDefaults() *GetV1LogEventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 log event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1LogEventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 log event params
func (o *GetV1LogEventParams) WithTimeout(timeout time.Duration) *GetV1LogEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 log event params
func (o *GetV1LogEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 log event params
func (o *GetV1LogEventParams) WithContext(ctx context.Context) *GetV1LogEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 log event params
func (o *GetV1LogEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 log event params
func (o *GetV1LogEventParams) WithHTTPClient(client *http.Client) *GetV1LogEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 log event params
func (o *GetV1LogEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 log event params
func (o *GetV1LogEventParams) WithXAuthToken(xAuthToken string) *GetV1LogEventParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 log event params
func (o *GetV1LogEventParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1LogEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
