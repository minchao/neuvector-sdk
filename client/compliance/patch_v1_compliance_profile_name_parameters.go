// Code generated by go-swagger; DO NOT EDIT.

package compliance

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

// NewPatchV1ComplianceProfileNameParams creates a new PatchV1ComplianceProfileNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ComplianceProfileNameParams() *PatchV1ComplianceProfileNameParams {
	return &PatchV1ComplianceProfileNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ComplianceProfileNameParamsWithTimeout creates a new PatchV1ComplianceProfileNameParams object
// with the ability to set a timeout on a request.
func NewPatchV1ComplianceProfileNameParamsWithTimeout(timeout time.Duration) *PatchV1ComplianceProfileNameParams {
	return &PatchV1ComplianceProfileNameParams{
		timeout: timeout,
	}
}

// NewPatchV1ComplianceProfileNameParamsWithContext creates a new PatchV1ComplianceProfileNameParams object
// with the ability to set a context for a request.
func NewPatchV1ComplianceProfileNameParamsWithContext(ctx context.Context) *PatchV1ComplianceProfileNameParams {
	return &PatchV1ComplianceProfileNameParams{
		Context: ctx,
	}
}

// NewPatchV1ComplianceProfileNameParamsWithHTTPClient creates a new PatchV1ComplianceProfileNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ComplianceProfileNameParamsWithHTTPClient(client *http.Client) *PatchV1ComplianceProfileNameParams {
	return &PatchV1ComplianceProfileNameParams{
		HTTPClient: client,
	}
}

/*
PatchV1ComplianceProfileNameParams contains all the parameters to send to the API endpoint

	for the patch v1 compliance profile name operation.

	Typically these are written to a http.Request.
*/
type PatchV1ComplianceProfileNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Compliance profile config data
	*/
	Body *models.RESTComplianceProfileConfigData

	/* Name.

	   Compliance profile name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 compliance profile name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ComplianceProfileNameParams) WithDefaults() *PatchV1ComplianceProfileNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 compliance profile name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ComplianceProfileNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) WithTimeout(timeout time.Duration) *PatchV1ComplianceProfileNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) WithContext(ctx context.Context) *PatchV1ComplianceProfileNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) WithHTTPClient(client *http.Client) *PatchV1ComplianceProfileNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) WithXAuthToken(xAuthToken string) *PatchV1ComplianceProfileNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) WithBody(body *models.RESTComplianceProfileConfigData) *PatchV1ComplianceProfileNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) SetBody(body *models.RESTComplianceProfileConfigData) {
	o.Body = body
}

// WithName adds the name to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) WithName(name string) *PatchV1ComplianceProfileNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 compliance profile name params
func (o *PatchV1ComplianceProfileNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ComplianceProfileNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
