// Code generated by go-swagger; DO NOT EDIT.

package sniffer

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

// NewGetV1SnifferParams creates a new GetV1SnifferParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SnifferParams() *GetV1SnifferParams {
	return &GetV1SnifferParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SnifferParamsWithTimeout creates a new GetV1SnifferParams object
// with the ability to set a timeout on a request.
func NewGetV1SnifferParamsWithTimeout(timeout time.Duration) *GetV1SnifferParams {
	return &GetV1SnifferParams{
		timeout: timeout,
	}
}

// NewGetV1SnifferParamsWithContext creates a new GetV1SnifferParams object
// with the ability to set a context for a request.
func NewGetV1SnifferParamsWithContext(ctx context.Context) *GetV1SnifferParams {
	return &GetV1SnifferParams{
		Context: ctx,
	}
}

// NewGetV1SnifferParamsWithHTTPClient creates a new GetV1SnifferParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SnifferParamsWithHTTPClient(client *http.Client) *GetV1SnifferParams {
	return &GetV1SnifferParams{
		HTTPClient: client,
	}
}

/*
GetV1SnifferParams contains all the parameters to send to the API endpoint

	for the get v1 sniffer operation.

	Typically these are written to a http.Request.
*/
type GetV1SnifferParams struct {

	/* FWorkload.

	   Workload ID
	*/
	FWorkload string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 sniffer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SnifferParams) WithDefaults() *GetV1SnifferParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 sniffer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SnifferParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 sniffer params
func (o *GetV1SnifferParams) WithTimeout(timeout time.Duration) *GetV1SnifferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 sniffer params
func (o *GetV1SnifferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 sniffer params
func (o *GetV1SnifferParams) WithContext(ctx context.Context) *GetV1SnifferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 sniffer params
func (o *GetV1SnifferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 sniffer params
func (o *GetV1SnifferParams) WithHTTPClient(client *http.Client) *GetV1SnifferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 sniffer params
func (o *GetV1SnifferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFWorkload adds the fWorkload to the get v1 sniffer params
func (o *GetV1SnifferParams) WithFWorkload(fWorkload string) *GetV1SnifferParams {
	o.SetFWorkload(fWorkload)
	return o
}

// SetFWorkload adds the fWorkload to the get v1 sniffer params
func (o *GetV1SnifferParams) SetFWorkload(fWorkload string) {
	o.FWorkload = fWorkload
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SnifferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param f_workload
	qrFWorkload := o.FWorkload
	qFWorkload := qrFWorkload
	if qFWorkload != "" {

		if err := r.SetQueryParam("f_workload", qFWorkload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
