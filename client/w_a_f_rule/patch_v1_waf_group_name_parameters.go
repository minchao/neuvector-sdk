// Code generated by go-swagger; DO NOT EDIT.

package w_a_f_rule

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

// NewPatchV1WafGroupNameParams creates a new PatchV1WafGroupNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1WafGroupNameParams() *PatchV1WafGroupNameParams {
	return &PatchV1WafGroupNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1WafGroupNameParamsWithTimeout creates a new PatchV1WafGroupNameParams object
// with the ability to set a timeout on a request.
func NewPatchV1WafGroupNameParamsWithTimeout(timeout time.Duration) *PatchV1WafGroupNameParams {
	return &PatchV1WafGroupNameParams{
		timeout: timeout,
	}
}

// NewPatchV1WafGroupNameParamsWithContext creates a new PatchV1WafGroupNameParams object
// with the ability to set a context for a request.
func NewPatchV1WafGroupNameParamsWithContext(ctx context.Context) *PatchV1WafGroupNameParams {
	return &PatchV1WafGroupNameParams{
		Context: ctx,
	}
}

// NewPatchV1WafGroupNameParamsWithHTTPClient creates a new PatchV1WafGroupNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1WafGroupNameParamsWithHTTPClient(client *http.Client) *PatchV1WafGroupNameParams {
	return &PatchV1WafGroupNameParams{
		HTTPClient: client,
	}
}

/*
PatchV1WafGroupNameParams contains all the parameters to send to the API endpoint

	for the patch v1 waf group name operation.

	Typically these are written to a http.Request.
*/
type PatchV1WafGroupNameParams struct {

	/* Body.

	   waf group data
	*/
	Body *models.RESTWafGroupConfigData

	/* Name.

	   waf group name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 waf group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1WafGroupNameParams) WithDefaults() *PatchV1WafGroupNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 waf group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1WafGroupNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) WithTimeout(timeout time.Duration) *PatchV1WafGroupNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) WithContext(ctx context.Context) *PatchV1WafGroupNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) WithHTTPClient(client *http.Client) *PatchV1WafGroupNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) WithBody(body *models.RESTWafGroupConfigData) *PatchV1WafGroupNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) SetBody(body *models.RESTWafGroupConfigData) {
	o.Body = body
}

// WithName adds the name to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) WithName(name string) *PatchV1WafGroupNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 waf group name params
func (o *PatchV1WafGroupNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1WafGroupNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
