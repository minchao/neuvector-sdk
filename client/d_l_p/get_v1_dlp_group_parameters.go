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

// NewGetV1DlpGroupParams creates a new GetV1DlpGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1DlpGroupParams() *GetV1DlpGroupParams {
	return &GetV1DlpGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1DlpGroupParamsWithTimeout creates a new GetV1DlpGroupParams object
// with the ability to set a timeout on a request.
func NewGetV1DlpGroupParamsWithTimeout(timeout time.Duration) *GetV1DlpGroupParams {
	return &GetV1DlpGroupParams{
		timeout: timeout,
	}
}

// NewGetV1DlpGroupParamsWithContext creates a new GetV1DlpGroupParams object
// with the ability to set a context for a request.
func NewGetV1DlpGroupParamsWithContext(ctx context.Context) *GetV1DlpGroupParams {
	return &GetV1DlpGroupParams{
		Context: ctx,
	}
}

// NewGetV1DlpGroupParamsWithHTTPClient creates a new GetV1DlpGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1DlpGroupParamsWithHTTPClient(client *http.Client) *GetV1DlpGroupParams {
	return &GetV1DlpGroupParams{
		HTTPClient: client,
	}
}

/*
GetV1DlpGroupParams contains all the parameters to send to the API endpoint

	for the get v1 dlp group operation.

	Typically these are written to a http.Request.
*/
type GetV1DlpGroupParams struct {

	// XAuthToken.
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 dlp group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DlpGroupParams) WithDefaults() *GetV1DlpGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 dlp group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DlpGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 dlp group params
func (o *GetV1DlpGroupParams) WithTimeout(timeout time.Duration) *GetV1DlpGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 dlp group params
func (o *GetV1DlpGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 dlp group params
func (o *GetV1DlpGroupParams) WithContext(ctx context.Context) *GetV1DlpGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 dlp group params
func (o *GetV1DlpGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 dlp group params
func (o *GetV1DlpGroupParams) WithHTTPClient(client *http.Client) *GetV1DlpGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 dlp group params
func (o *GetV1DlpGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 dlp group params
func (o *GetV1DlpGroupParams) WithXAuthToken(xAuthToken string) *GetV1DlpGroupParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 dlp group params
func (o *GetV1DlpGroupParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1DlpGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
