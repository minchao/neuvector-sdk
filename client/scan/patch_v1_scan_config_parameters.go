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

// NewPatchV1ScanConfigParams creates a new PatchV1ScanConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ScanConfigParams() *PatchV1ScanConfigParams {
	return &PatchV1ScanConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ScanConfigParamsWithTimeout creates a new PatchV1ScanConfigParams object
// with the ability to set a timeout on a request.
func NewPatchV1ScanConfigParamsWithTimeout(timeout time.Duration) *PatchV1ScanConfigParams {
	return &PatchV1ScanConfigParams{
		timeout: timeout,
	}
}

// NewPatchV1ScanConfigParamsWithContext creates a new PatchV1ScanConfigParams object
// with the ability to set a context for a request.
func NewPatchV1ScanConfigParamsWithContext(ctx context.Context) *PatchV1ScanConfigParams {
	return &PatchV1ScanConfigParams{
		Context: ctx,
	}
}

// NewPatchV1ScanConfigParamsWithHTTPClient creates a new PatchV1ScanConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ScanConfigParamsWithHTTPClient(client *http.Client) *PatchV1ScanConfigParams {
	return &PatchV1ScanConfigParams{
		HTTPClient: client,
	}
}

/*
PatchV1ScanConfigParams contains all the parameters to send to the API endpoint

	for the patch v1 scan config operation.

	Typically these are written to a http.Request.
*/
type PatchV1ScanConfigParams struct {

	/* Body.

	   Scan configure data
	*/
	Body *models.RESTScanConfigData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 scan config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScanConfigParams) WithDefaults() *PatchV1ScanConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 scan config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScanConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) WithTimeout(timeout time.Duration) *PatchV1ScanConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) WithContext(ctx context.Context) *PatchV1ScanConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) WithHTTPClient(client *http.Client) *PatchV1ScanConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) WithBody(body *models.RESTScanConfigData) *PatchV1ScanConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 scan config params
func (o *PatchV1ScanConfigParams) SetBody(body *models.RESTScanConfigData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ScanConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
