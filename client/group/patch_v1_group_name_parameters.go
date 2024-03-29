// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewPatchV1GroupNameParams creates a new PatchV1GroupNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1GroupNameParams() *PatchV1GroupNameParams {
	return &PatchV1GroupNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1GroupNameParamsWithTimeout creates a new PatchV1GroupNameParams object
// with the ability to set a timeout on a request.
func NewPatchV1GroupNameParamsWithTimeout(timeout time.Duration) *PatchV1GroupNameParams {
	return &PatchV1GroupNameParams{
		timeout: timeout,
	}
}

// NewPatchV1GroupNameParamsWithContext creates a new PatchV1GroupNameParams object
// with the ability to set a context for a request.
func NewPatchV1GroupNameParamsWithContext(ctx context.Context) *PatchV1GroupNameParams {
	return &PatchV1GroupNameParams{
		Context: ctx,
	}
}

// NewPatchV1GroupNameParamsWithHTTPClient creates a new PatchV1GroupNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1GroupNameParamsWithHTTPClient(client *http.Client) *PatchV1GroupNameParams {
	return &PatchV1GroupNameParams{
		HTTPClient: client,
	}
}

/*
PatchV1GroupNameParams contains all the parameters to send to the API endpoint

	for the patch v1 group name operation.

	Typically these are written to a http.Request.
*/
type PatchV1GroupNameParams struct {

	/* Body.

	   Group update data
	*/
	Body *models.RESTGroupConfigData

	/* Name.

	   Group name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1GroupNameParams) WithDefaults() *PatchV1GroupNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1GroupNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 group name params
func (o *PatchV1GroupNameParams) WithTimeout(timeout time.Duration) *PatchV1GroupNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 group name params
func (o *PatchV1GroupNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 group name params
func (o *PatchV1GroupNameParams) WithContext(ctx context.Context) *PatchV1GroupNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 group name params
func (o *PatchV1GroupNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 group name params
func (o *PatchV1GroupNameParams) WithHTTPClient(client *http.Client) *PatchV1GroupNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 group name params
func (o *PatchV1GroupNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch v1 group name params
func (o *PatchV1GroupNameParams) WithBody(body *models.RESTGroupConfigData) *PatchV1GroupNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 group name params
func (o *PatchV1GroupNameParams) SetBody(body *models.RESTGroupConfigData) {
	o.Body = body
}

// WithName adds the name to the patch v1 group name params
func (o *PatchV1GroupNameParams) WithName(name string) *PatchV1GroupNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 group name params
func (o *PatchV1GroupNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1GroupNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
