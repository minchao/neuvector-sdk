// Code generated by go-swagger; DO NOT EDIT.

package file

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

// NewPostV1FileDlpConfigParams creates a new PostV1FileDlpConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1FileDlpConfigParams() *PostV1FileDlpConfigParams {
	return &PostV1FileDlpConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1FileDlpConfigParamsWithTimeout creates a new PostV1FileDlpConfigParams object
// with the ability to set a timeout on a request.
func NewPostV1FileDlpConfigParamsWithTimeout(timeout time.Duration) *PostV1FileDlpConfigParams {
	return &PostV1FileDlpConfigParams{
		timeout: timeout,
	}
}

// NewPostV1FileDlpConfigParamsWithContext creates a new PostV1FileDlpConfigParams object
// with the ability to set a context for a request.
func NewPostV1FileDlpConfigParamsWithContext(ctx context.Context) *PostV1FileDlpConfigParams {
	return &PostV1FileDlpConfigParams{
		Context: ctx,
	}
}

// NewPostV1FileDlpConfigParamsWithHTTPClient creates a new PostV1FileDlpConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1FileDlpConfigParamsWithHTTPClient(client *http.Client) *PostV1FileDlpConfigParams {
	return &PostV1FileDlpConfigParams{
		HTTPClient: client,
	}
}

/*
PostV1FileDlpConfigParams contains all the parameters to send to the API endpoint

	for the post v1 file dlp config operation.

	Typically these are written to a http.Request.
*/
type PostV1FileDlpConfigParams struct {

	// XTransactionID.
	XTransactionID *string

	/* Body.

	   DLP config yaml file
	*/
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 file dlp config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FileDlpConfigParams) WithDefaults() *PostV1FileDlpConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 file dlp config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FileDlpConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) WithTimeout(timeout time.Duration) *PostV1FileDlpConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) WithContext(ctx context.Context) *PostV1FileDlpConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) WithHTTPClient(client *http.Client) *PostV1FileDlpConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXTransactionID adds the xTransactionID to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) WithXTransactionID(xTransactionID *string) *PostV1FileDlpConfigParams {
	o.SetXTransactionID(xTransactionID)
	return o
}

// SetXTransactionID adds the xTransactionId to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) SetXTransactionID(xTransactionID *string) {
	o.XTransactionID = xTransactionID
}

// WithBody adds the body to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) WithBody(body string) *PostV1FileDlpConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 file dlp config params
func (o *PostV1FileDlpConfigParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1FileDlpConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XTransactionID != nil {

		// header param X-Transaction-ID
		if err := r.SetHeaderParam("X-Transaction-ID", *o.XTransactionID); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
