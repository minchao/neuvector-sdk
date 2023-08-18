// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewPatchV1AuthParams creates a new PatchV1AuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1AuthParams() *PatchV1AuthParams {
	return &PatchV1AuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1AuthParamsWithTimeout creates a new PatchV1AuthParams object
// with the ability to set a timeout on a request.
func NewPatchV1AuthParamsWithTimeout(timeout time.Duration) *PatchV1AuthParams {
	return &PatchV1AuthParams{
		timeout: timeout,
	}
}

// NewPatchV1AuthParamsWithContext creates a new PatchV1AuthParams object
// with the ability to set a context for a request.
func NewPatchV1AuthParamsWithContext(ctx context.Context) *PatchV1AuthParams {
	return &PatchV1AuthParams{
		Context: ctx,
	}
}

// NewPatchV1AuthParamsWithHTTPClient creates a new PatchV1AuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1AuthParamsWithHTTPClient(client *http.Client) *PatchV1AuthParams {
	return &PatchV1AuthParams{
		HTTPClient: client,
	}
}

/*
PatchV1AuthParams contains all the parameters to send to the API endpoint

	for the patch v1 auth operation.

	Typically these are written to a http.Request.
*/
type PatchV1AuthParams struct {

	// XAuthToken.
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1AuthParams) WithDefaults() *PatchV1AuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1AuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 auth params
func (o *PatchV1AuthParams) WithTimeout(timeout time.Duration) *PatchV1AuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 auth params
func (o *PatchV1AuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 auth params
func (o *PatchV1AuthParams) WithContext(ctx context.Context) *PatchV1AuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 auth params
func (o *PatchV1AuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 auth params
func (o *PatchV1AuthParams) WithHTTPClient(client *http.Client) *PatchV1AuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 auth params
func (o *PatchV1AuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the patch v1 auth params
func (o *PatchV1AuthParams) WithXAuthToken(xAuthToken string) *PatchV1AuthParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the patch v1 auth params
func (o *PatchV1AuthParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1AuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
