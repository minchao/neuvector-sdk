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

	"github.com/minchao/neuvector-sdk/models"
)

// NewPostV1ScanSigstoreRootOfTrustParams creates a new PostV1ScanSigstoreRootOfTrustParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ScanSigstoreRootOfTrustParams() *PostV1ScanSigstoreRootOfTrustParams {
	return &PostV1ScanSigstoreRootOfTrustParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ScanSigstoreRootOfTrustParamsWithTimeout creates a new PostV1ScanSigstoreRootOfTrustParams object
// with the ability to set a timeout on a request.
func NewPostV1ScanSigstoreRootOfTrustParamsWithTimeout(timeout time.Duration) *PostV1ScanSigstoreRootOfTrustParams {
	return &PostV1ScanSigstoreRootOfTrustParams{
		timeout: timeout,
	}
}

// NewPostV1ScanSigstoreRootOfTrustParamsWithContext creates a new PostV1ScanSigstoreRootOfTrustParams object
// with the ability to set a context for a request.
func NewPostV1ScanSigstoreRootOfTrustParamsWithContext(ctx context.Context) *PostV1ScanSigstoreRootOfTrustParams {
	return &PostV1ScanSigstoreRootOfTrustParams{
		Context: ctx,
	}
}

// NewPostV1ScanSigstoreRootOfTrustParamsWithHTTPClient creates a new PostV1ScanSigstoreRootOfTrustParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ScanSigstoreRootOfTrustParamsWithHTTPClient(client *http.Client) *PostV1ScanSigstoreRootOfTrustParams {
	return &PostV1ScanSigstoreRootOfTrustParams{
		HTTPClient: client,
	}
}

/*
PostV1ScanSigstoreRootOfTrustParams contains all the parameters to send to the API endpoint

	for the post v1 scan sigstore root of trust operation.

	Typically these are written to a http.Request.
*/
type PostV1ScanSigstoreRootOfTrustParams struct {

	/* Body.

	   Root of Trust Data
	*/
	Body *models.RESTSigstoreRootOfTrustPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 scan sigstore root of trust params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ScanSigstoreRootOfTrustParams) WithDefaults() *PostV1ScanSigstoreRootOfTrustParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 scan sigstore root of trust params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ScanSigstoreRootOfTrustParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) WithTimeout(timeout time.Duration) *PostV1ScanSigstoreRootOfTrustParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) WithContext(ctx context.Context) *PostV1ScanSigstoreRootOfTrustParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) WithHTTPClient(client *http.Client) *PostV1ScanSigstoreRootOfTrustParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) WithBody(body *models.RESTSigstoreRootOfTrustPost) *PostV1ScanSigstoreRootOfTrustParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 scan sigstore root of trust params
func (o *PostV1ScanSigstoreRootOfTrustParams) SetBody(body *models.RESTSigstoreRootOfTrustPost) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ScanSigstoreRootOfTrustParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
