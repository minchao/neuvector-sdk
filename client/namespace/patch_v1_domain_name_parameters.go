// Code generated by go-swagger; DO NOT EDIT.

package namespace

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

// NewPatchV1DomainNameParams creates a new PatchV1DomainNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1DomainNameParams() *PatchV1DomainNameParams {
	return &PatchV1DomainNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1DomainNameParamsWithTimeout creates a new PatchV1DomainNameParams object
// with the ability to set a timeout on a request.
func NewPatchV1DomainNameParamsWithTimeout(timeout time.Duration) *PatchV1DomainNameParams {
	return &PatchV1DomainNameParams{
		timeout: timeout,
	}
}

// NewPatchV1DomainNameParamsWithContext creates a new PatchV1DomainNameParams object
// with the ability to set a context for a request.
func NewPatchV1DomainNameParamsWithContext(ctx context.Context) *PatchV1DomainNameParams {
	return &PatchV1DomainNameParams{
		Context: ctx,
	}
}

// NewPatchV1DomainNameParamsWithHTTPClient creates a new PatchV1DomainNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1DomainNameParamsWithHTTPClient(client *http.Client) *PatchV1DomainNameParams {
	return &PatchV1DomainNameParams{
		HTTPClient: client,
	}
}

/*
PatchV1DomainNameParams contains all the parameters to send to the API endpoint

	for the patch v1 domain name operation.

	Typically these are written to a http.Request.
*/
type PatchV1DomainNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Namespace update data
	*/
	Body *models.RESTDomainEntryConfigData

	/* Name.

	   namespace name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 domain name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1DomainNameParams) WithDefaults() *PatchV1DomainNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 domain name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1DomainNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 domain name params
func (o *PatchV1DomainNameParams) WithTimeout(timeout time.Duration) *PatchV1DomainNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 domain name params
func (o *PatchV1DomainNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 domain name params
func (o *PatchV1DomainNameParams) WithContext(ctx context.Context) *PatchV1DomainNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 domain name params
func (o *PatchV1DomainNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 domain name params
func (o *PatchV1DomainNameParams) WithHTTPClient(client *http.Client) *PatchV1DomainNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 domain name params
func (o *PatchV1DomainNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the patch v1 domain name params
func (o *PatchV1DomainNameParams) WithXAuthToken(xAuthToken string) *PatchV1DomainNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the patch v1 domain name params
func (o *PatchV1DomainNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the patch v1 domain name params
func (o *PatchV1DomainNameParams) WithBody(body *models.RESTDomainEntryConfigData) *PatchV1DomainNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 domain name params
func (o *PatchV1DomainNameParams) SetBody(body *models.RESTDomainEntryConfigData) {
	o.Body = body
}

// WithName adds the name to the patch v1 domain name params
func (o *PatchV1DomainNameParams) WithName(name string) *PatchV1DomainNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 domain name params
func (o *PatchV1DomainNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1DomainNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
