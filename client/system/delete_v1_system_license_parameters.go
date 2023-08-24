// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewDeleteV1SystemLicenseParams creates a new DeleteV1SystemLicenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1SystemLicenseParams() *DeleteV1SystemLicenseParams {
	return &DeleteV1SystemLicenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1SystemLicenseParamsWithTimeout creates a new DeleteV1SystemLicenseParams object
// with the ability to set a timeout on a request.
func NewDeleteV1SystemLicenseParamsWithTimeout(timeout time.Duration) *DeleteV1SystemLicenseParams {
	return &DeleteV1SystemLicenseParams{
		timeout: timeout,
	}
}

// NewDeleteV1SystemLicenseParamsWithContext creates a new DeleteV1SystemLicenseParams object
// with the ability to set a context for a request.
func NewDeleteV1SystemLicenseParamsWithContext(ctx context.Context) *DeleteV1SystemLicenseParams {
	return &DeleteV1SystemLicenseParams{
		Context: ctx,
	}
}

// NewDeleteV1SystemLicenseParamsWithHTTPClient creates a new DeleteV1SystemLicenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1SystemLicenseParamsWithHTTPClient(client *http.Client) *DeleteV1SystemLicenseParams {
	return &DeleteV1SystemLicenseParams{
		HTTPClient: client,
	}
}

/*
DeleteV1SystemLicenseParams contains all the parameters to send to the API endpoint

	for the delete v1 system license operation.

	Typically these are written to a http.Request.
*/
type DeleteV1SystemLicenseParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 system license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SystemLicenseParams) WithDefaults() *DeleteV1SystemLicenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 system license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SystemLicenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 system license params
func (o *DeleteV1SystemLicenseParams) WithTimeout(timeout time.Duration) *DeleteV1SystemLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 system license params
func (o *DeleteV1SystemLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 system license params
func (o *DeleteV1SystemLicenseParams) WithContext(ctx context.Context) *DeleteV1SystemLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 system license params
func (o *DeleteV1SystemLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 system license params
func (o *DeleteV1SystemLicenseParams) WithHTTPClient(client *http.Client) *DeleteV1SystemLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 system license params
func (o *DeleteV1SystemLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1SystemLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
