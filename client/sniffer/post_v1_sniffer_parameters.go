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

	"github.com/minchao/neuvector-sdk/models"
)

// NewPostV1SnifferParams creates a new PostV1SnifferParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1SnifferParams() *PostV1SnifferParams {
	return &PostV1SnifferParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1SnifferParamsWithTimeout creates a new PostV1SnifferParams object
// with the ability to set a timeout on a request.
func NewPostV1SnifferParamsWithTimeout(timeout time.Duration) *PostV1SnifferParams {
	return &PostV1SnifferParams{
		timeout: timeout,
	}
}

// NewPostV1SnifferParamsWithContext creates a new PostV1SnifferParams object
// with the ability to set a context for a request.
func NewPostV1SnifferParamsWithContext(ctx context.Context) *PostV1SnifferParams {
	return &PostV1SnifferParams{
		Context: ctx,
	}
}

// NewPostV1SnifferParamsWithHTTPClient creates a new PostV1SnifferParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1SnifferParamsWithHTTPClient(client *http.Client) *PostV1SnifferParams {
	return &PostV1SnifferParams{
		HTTPClient: client,
	}
}

/*
PostV1SnifferParams contains all the parameters to send to the API endpoint

	for the post v1 sniffer operation.

	Typically these are written to a http.Request.
*/
type PostV1SnifferParams struct {

	/* Body.

	   Sniffer args data
	*/
	Body *models.RESTSnifferArgsData

	/* FWorkload.

	   Workload ID
	*/
	FWorkload string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 sniffer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SnifferParams) WithDefaults() *PostV1SnifferParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 sniffer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SnifferParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 sniffer params
func (o *PostV1SnifferParams) WithTimeout(timeout time.Duration) *PostV1SnifferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 sniffer params
func (o *PostV1SnifferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 sniffer params
func (o *PostV1SnifferParams) WithContext(ctx context.Context) *PostV1SnifferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 sniffer params
func (o *PostV1SnifferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 sniffer params
func (o *PostV1SnifferParams) WithHTTPClient(client *http.Client) *PostV1SnifferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 sniffer params
func (o *PostV1SnifferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post v1 sniffer params
func (o *PostV1SnifferParams) WithBody(body *models.RESTSnifferArgsData) *PostV1SnifferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 sniffer params
func (o *PostV1SnifferParams) SetBody(body *models.RESTSnifferArgsData) {
	o.Body = body
}

// WithFWorkload adds the fWorkload to the post v1 sniffer params
func (o *PostV1SnifferParams) WithFWorkload(fWorkload string) *PostV1SnifferParams {
	o.SetFWorkload(fWorkload)
	return o
}

// SetFWorkload adds the fWorkload to the post v1 sniffer params
func (o *PostV1SnifferParams) SetFWorkload(fWorkload string) {
	o.FWorkload = fWorkload
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1SnifferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
