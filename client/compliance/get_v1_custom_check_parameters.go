// Code generated by go-swagger; DO NOT EDIT.

package compliance

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

// NewGetV1CustomCheckParams creates a new GetV1CustomCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1CustomCheckParams() *GetV1CustomCheckParams {
	return &GetV1CustomCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1CustomCheckParamsWithTimeout creates a new GetV1CustomCheckParams object
// with the ability to set a timeout on a request.
func NewGetV1CustomCheckParamsWithTimeout(timeout time.Duration) *GetV1CustomCheckParams {
	return &GetV1CustomCheckParams{
		timeout: timeout,
	}
}

// NewGetV1CustomCheckParamsWithContext creates a new GetV1CustomCheckParams object
// with the ability to set a context for a request.
func NewGetV1CustomCheckParamsWithContext(ctx context.Context) *GetV1CustomCheckParams {
	return &GetV1CustomCheckParams{
		Context: ctx,
	}
}

// NewGetV1CustomCheckParamsWithHTTPClient creates a new GetV1CustomCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1CustomCheckParamsWithHTTPClient(client *http.Client) *GetV1CustomCheckParams {
	return &GetV1CustomCheckParams{
		HTTPClient: client,
	}
}

/*
GetV1CustomCheckParams contains all the parameters to send to the API endpoint

	for the get v1 custom check operation.

	Typically these are written to a http.Request.
*/
type GetV1CustomCheckParams struct {

	// XAuthToken.
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 custom check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1CustomCheckParams) WithDefaults() *GetV1CustomCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 custom check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1CustomCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 custom check params
func (o *GetV1CustomCheckParams) WithTimeout(timeout time.Duration) *GetV1CustomCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 custom check params
func (o *GetV1CustomCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 custom check params
func (o *GetV1CustomCheckParams) WithContext(ctx context.Context) *GetV1CustomCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 custom check params
func (o *GetV1CustomCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 custom check params
func (o *GetV1CustomCheckParams) WithHTTPClient(client *http.Client) *GetV1CustomCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 custom check params
func (o *GetV1CustomCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 custom check params
func (o *GetV1CustomCheckParams) WithXAuthToken(xAuthToken string) *GetV1CustomCheckParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 custom check params
func (o *GetV1CustomCheckParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1CustomCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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