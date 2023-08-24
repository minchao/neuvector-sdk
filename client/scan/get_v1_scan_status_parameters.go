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

// NewGetV1ScanStatusParams creates a new GetV1ScanStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ScanStatusParams() *GetV1ScanStatusParams {
	return &GetV1ScanStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ScanStatusParamsWithTimeout creates a new GetV1ScanStatusParams object
// with the ability to set a timeout on a request.
func NewGetV1ScanStatusParamsWithTimeout(timeout time.Duration) *GetV1ScanStatusParams {
	return &GetV1ScanStatusParams{
		timeout: timeout,
	}
}

// NewGetV1ScanStatusParamsWithContext creates a new GetV1ScanStatusParams object
// with the ability to set a context for a request.
func NewGetV1ScanStatusParamsWithContext(ctx context.Context) *GetV1ScanStatusParams {
	return &GetV1ScanStatusParams{
		Context: ctx,
	}
}

// NewGetV1ScanStatusParamsWithHTTPClient creates a new GetV1ScanStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ScanStatusParamsWithHTTPClient(client *http.Client) *GetV1ScanStatusParams {
	return &GetV1ScanStatusParams{
		HTTPClient: client,
	}
}

/*
GetV1ScanStatusParams contains all the parameters to send to the API endpoint

	for the get v1 scan status operation.

	Typically these are written to a http.Request.
*/
type GetV1ScanStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 scan status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanStatusParams) WithDefaults() *GetV1ScanStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 scan status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 scan status params
func (o *GetV1ScanStatusParams) WithTimeout(timeout time.Duration) *GetV1ScanStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 scan status params
func (o *GetV1ScanStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 scan status params
func (o *GetV1ScanStatusParams) WithContext(ctx context.Context) *GetV1ScanStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 scan status params
func (o *GetV1ScanStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 scan status params
func (o *GetV1ScanStatusParams) WithHTTPClient(client *http.Client) *GetV1ScanStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 scan status params
func (o *GetV1ScanStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ScanStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
