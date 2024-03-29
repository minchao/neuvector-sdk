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

	"github.com/minchao/neuvector-sdk/models"
)

// NewPostV1DlpSensorParams creates a new PostV1DlpSensorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1DlpSensorParams() *PostV1DlpSensorParams {
	return &PostV1DlpSensorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1DlpSensorParamsWithTimeout creates a new PostV1DlpSensorParams object
// with the ability to set a timeout on a request.
func NewPostV1DlpSensorParamsWithTimeout(timeout time.Duration) *PostV1DlpSensorParams {
	return &PostV1DlpSensorParams{
		timeout: timeout,
	}
}

// NewPostV1DlpSensorParamsWithContext creates a new PostV1DlpSensorParams object
// with the ability to set a context for a request.
func NewPostV1DlpSensorParamsWithContext(ctx context.Context) *PostV1DlpSensorParams {
	return &PostV1DlpSensorParams{
		Context: ctx,
	}
}

// NewPostV1DlpSensorParamsWithHTTPClient creates a new PostV1DlpSensorParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1DlpSensorParamsWithHTTPClient(client *http.Client) *PostV1DlpSensorParams {
	return &PostV1DlpSensorParams{
		HTTPClient: client,
	}
}

/*
PostV1DlpSensorParams contains all the parameters to send to the API endpoint

	for the post v1 dlp sensor operation.

	Typically these are written to a http.Request.
*/
type PostV1DlpSensorParams struct {

	/* Body.

	   Sensor data
	*/
	Body *models.RESTDlpSensorConfigData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 dlp sensor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1DlpSensorParams) WithDefaults() *PostV1DlpSensorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 dlp sensor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1DlpSensorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) WithTimeout(timeout time.Duration) *PostV1DlpSensorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) WithContext(ctx context.Context) *PostV1DlpSensorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) WithHTTPClient(client *http.Client) *PostV1DlpSensorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) WithBody(body *models.RESTDlpSensorConfigData) *PostV1DlpSensorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 dlp sensor params
func (o *PostV1DlpSensorParams) SetBody(body *models.RESTDlpSensorConfigData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1DlpSensorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
