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

	"github.com/minchao/neuvector-sdk/models"
)

// NewPostV1FileDlpParams creates a new PostV1FileDlpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1FileDlpParams() *PostV1FileDlpParams {
	return &PostV1FileDlpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1FileDlpParamsWithTimeout creates a new PostV1FileDlpParams object
// with the ability to set a timeout on a request.
func NewPostV1FileDlpParamsWithTimeout(timeout time.Duration) *PostV1FileDlpParams {
	return &PostV1FileDlpParams{
		timeout: timeout,
	}
}

// NewPostV1FileDlpParamsWithContext creates a new PostV1FileDlpParams object
// with the ability to set a context for a request.
func NewPostV1FileDlpParamsWithContext(ctx context.Context) *PostV1FileDlpParams {
	return &PostV1FileDlpParams{
		Context: ctx,
	}
}

// NewPostV1FileDlpParamsWithHTTPClient creates a new PostV1FileDlpParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1FileDlpParamsWithHTTPClient(client *http.Client) *PostV1FileDlpParams {
	return &PostV1FileDlpParams{
		HTTPClient: client,
	}
}

/*
PostV1FileDlpParams contains all the parameters to send to the API endpoint

	for the post v1 file dlp operation.

	Typically these are written to a http.Request.
*/
type PostV1FileDlpParams struct {

	/* Body.

	   Configuration data
	*/
	Body *models.RESTDlpSensorExport

	/* Scope.

	   It exports the DLP configurations when the scope is local.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 file dlp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FileDlpParams) WithDefaults() *PostV1FileDlpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 file dlp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FileDlpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 file dlp params
func (o *PostV1FileDlpParams) WithTimeout(timeout time.Duration) *PostV1FileDlpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 file dlp params
func (o *PostV1FileDlpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 file dlp params
func (o *PostV1FileDlpParams) WithContext(ctx context.Context) *PostV1FileDlpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 file dlp params
func (o *PostV1FileDlpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 file dlp params
func (o *PostV1FileDlpParams) WithHTTPClient(client *http.Client) *PostV1FileDlpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 file dlp params
func (o *PostV1FileDlpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post v1 file dlp params
func (o *PostV1FileDlpParams) WithBody(body *models.RESTDlpSensorExport) *PostV1FileDlpParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 file dlp params
func (o *PostV1FileDlpParams) SetBody(body *models.RESTDlpSensorExport) {
	o.Body = body
}

// WithScope adds the scope to the post v1 file dlp params
func (o *PostV1FileDlpParams) WithScope(scope *string) *PostV1FileDlpParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the post v1 file dlp params
func (o *PostV1FileDlpParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1FileDlpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
