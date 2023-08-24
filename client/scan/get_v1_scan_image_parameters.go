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

// NewGetV1ScanImageParams creates a new GetV1ScanImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ScanImageParams() *GetV1ScanImageParams {
	return &GetV1ScanImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ScanImageParamsWithTimeout creates a new GetV1ScanImageParams object
// with the ability to set a timeout on a request.
func NewGetV1ScanImageParamsWithTimeout(timeout time.Duration) *GetV1ScanImageParams {
	return &GetV1ScanImageParams{
		timeout: timeout,
	}
}

// NewGetV1ScanImageParamsWithContext creates a new GetV1ScanImageParams object
// with the ability to set a context for a request.
func NewGetV1ScanImageParamsWithContext(ctx context.Context) *GetV1ScanImageParams {
	return &GetV1ScanImageParams{
		Context: ctx,
	}
}

// NewGetV1ScanImageParamsWithHTTPClient creates a new GetV1ScanImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ScanImageParamsWithHTTPClient(client *http.Client) *GetV1ScanImageParams {
	return &GetV1ScanImageParams{
		HTTPClient: client,
	}
}

/*
GetV1ScanImageParams contains all the parameters to send to the API endpoint

	for the get v1 scan image operation.

	Typically these are written to a http.Request.
*/
type GetV1ScanImageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 scan image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanImageParams) WithDefaults() *GetV1ScanImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 scan image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 scan image params
func (o *GetV1ScanImageParams) WithTimeout(timeout time.Duration) *GetV1ScanImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 scan image params
func (o *GetV1ScanImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 scan image params
func (o *GetV1ScanImageParams) WithContext(ctx context.Context) *GetV1ScanImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 scan image params
func (o *GetV1ScanImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 scan image params
func (o *GetV1ScanImageParams) WithHTTPClient(client *http.Client) *GetV1ScanImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 scan image params
func (o *GetV1ScanImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ScanImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
