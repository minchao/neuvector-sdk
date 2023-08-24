// Code generated by go-swagger; DO NOT EDIT.

package d_l_p

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

// NewGetV1DlpSensorParams creates a new GetV1DlpSensorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1DlpSensorParams() *GetV1DlpSensorParams {
	return &GetV1DlpSensorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1DlpSensorParamsWithTimeout creates a new GetV1DlpSensorParams object
// with the ability to set a timeout on a request.
func NewGetV1DlpSensorParamsWithTimeout(timeout time.Duration) *GetV1DlpSensorParams {
	return &GetV1DlpSensorParams{
		timeout: timeout,
	}
}

// NewGetV1DlpSensorParamsWithContext creates a new GetV1DlpSensorParams object
// with the ability to set a context for a request.
func NewGetV1DlpSensorParamsWithContext(ctx context.Context) *GetV1DlpSensorParams {
	return &GetV1DlpSensorParams{
		Context: ctx,
	}
}

// NewGetV1DlpSensorParamsWithHTTPClient creates a new GetV1DlpSensorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1DlpSensorParamsWithHTTPClient(client *http.Client) *GetV1DlpSensorParams {
	return &GetV1DlpSensorParams{
		HTTPClient: client,
	}
}

/*
GetV1DlpSensorParams contains all the parameters to send to the API endpoint

	for the get v1 dlp sensor operation.

	Typically these are written to a http.Request.
*/
type GetV1DlpSensorParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 dlp sensor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DlpSensorParams) WithDefaults() *GetV1DlpSensorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 dlp sensor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DlpSensorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 dlp sensor params
func (o *GetV1DlpSensorParams) WithTimeout(timeout time.Duration) *GetV1DlpSensorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 dlp sensor params
func (o *GetV1DlpSensorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 dlp sensor params
func (o *GetV1DlpSensorParams) WithContext(ctx context.Context) *GetV1DlpSensorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 dlp sensor params
func (o *GetV1DlpSensorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 dlp sensor params
func (o *GetV1DlpSensorParams) WithHTTPClient(client *http.Client) *GetV1DlpSensorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 dlp sensor params
func (o *GetV1DlpSensorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1DlpSensorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
