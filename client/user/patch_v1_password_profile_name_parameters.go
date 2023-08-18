// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewPatchV1PasswordProfileNameParams creates a new PatchV1PasswordProfileNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1PasswordProfileNameParams() *PatchV1PasswordProfileNameParams {
	return &PatchV1PasswordProfileNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PasswordProfileNameParamsWithTimeout creates a new PatchV1PasswordProfileNameParams object
// with the ability to set a timeout on a request.
func NewPatchV1PasswordProfileNameParamsWithTimeout(timeout time.Duration) *PatchV1PasswordProfileNameParams {
	return &PatchV1PasswordProfileNameParams{
		timeout: timeout,
	}
}

// NewPatchV1PasswordProfileNameParamsWithContext creates a new PatchV1PasswordProfileNameParams object
// with the ability to set a context for a request.
func NewPatchV1PasswordProfileNameParamsWithContext(ctx context.Context) *PatchV1PasswordProfileNameParams {
	return &PatchV1PasswordProfileNameParams{
		Context: ctx,
	}
}

// NewPatchV1PasswordProfileNameParamsWithHTTPClient creates a new PatchV1PasswordProfileNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1PasswordProfileNameParamsWithHTTPClient(client *http.Client) *PatchV1PasswordProfileNameParams {
	return &PatchV1PasswordProfileNameParams{
		HTTPClient: client,
	}
}

/*
PatchV1PasswordProfileNameParams contains all the parameters to send to the API endpoint

	for the patch v1 password profile name operation.

	Typically these are written to a http.Request.
*/
type PatchV1PasswordProfileNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Password profile data
	*/
	Body *models.RESTPwdProfileConfigData

	/* Name.

	   Password profile name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 password profile name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PasswordProfileNameParams) WithDefaults() *PatchV1PasswordProfileNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 password profile name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PasswordProfileNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) WithTimeout(timeout time.Duration) *PatchV1PasswordProfileNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) WithContext(ctx context.Context) *PatchV1PasswordProfileNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) WithHTTPClient(client *http.Client) *PatchV1PasswordProfileNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) WithXAuthToken(xAuthToken string) *PatchV1PasswordProfileNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) WithBody(body *models.RESTPwdProfileConfigData) *PatchV1PasswordProfileNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) SetBody(body *models.RESTPwdProfileConfigData) {
	o.Body = body
}

// WithName adds the name to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) WithName(name string) *PatchV1PasswordProfileNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 password profile name params
func (o *PatchV1PasswordProfileNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PasswordProfileNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
