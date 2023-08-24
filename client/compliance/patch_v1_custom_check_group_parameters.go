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

// NewPatchV1CustomCheckGroupParams creates a new PatchV1CustomCheckGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1CustomCheckGroupParams() *PatchV1CustomCheckGroupParams {
	return &PatchV1CustomCheckGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1CustomCheckGroupParamsWithTimeout creates a new PatchV1CustomCheckGroupParams object
// with the ability to set a timeout on a request.
func NewPatchV1CustomCheckGroupParamsWithTimeout(timeout time.Duration) *PatchV1CustomCheckGroupParams {
	return &PatchV1CustomCheckGroupParams{
		timeout: timeout,
	}
}

// NewPatchV1CustomCheckGroupParamsWithContext creates a new PatchV1CustomCheckGroupParams object
// with the ability to set a context for a request.
func NewPatchV1CustomCheckGroupParamsWithContext(ctx context.Context) *PatchV1CustomCheckGroupParams {
	return &PatchV1CustomCheckGroupParams{
		Context: ctx,
	}
}

// NewPatchV1CustomCheckGroupParamsWithHTTPClient creates a new PatchV1CustomCheckGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1CustomCheckGroupParamsWithHTTPClient(client *http.Client) *PatchV1CustomCheckGroupParams {
	return &PatchV1CustomCheckGroupParams{
		HTTPClient: client,
	}
}

/*
PatchV1CustomCheckGroupParams contains all the parameters to send to the API endpoint

	for the patch v1 custom check group operation.

	Typically these are written to a http.Request.
*/
type PatchV1CustomCheckGroupParams struct {

	/* Body.

	   Script config data
	*/
	Body *models.RESTCustomCheckConfigData

	/* Group.

	   Script config name
	*/
	Group string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 custom check group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1CustomCheckGroupParams) WithDefaults() *PatchV1CustomCheckGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 custom check group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1CustomCheckGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) WithTimeout(timeout time.Duration) *PatchV1CustomCheckGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) WithContext(ctx context.Context) *PatchV1CustomCheckGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) WithHTTPClient(client *http.Client) *PatchV1CustomCheckGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) WithBody(body *models.RESTCustomCheckConfigData) *PatchV1CustomCheckGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) SetBody(body *models.RESTCustomCheckConfigData) {
	o.Body = body
}

// WithGroup adds the group to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) WithGroup(group string) *PatchV1CustomCheckGroupParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the patch v1 custom check group params
func (o *PatchV1CustomCheckGroupParams) SetGroup(group string) {
	o.Group = group
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1CustomCheckGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param group
	if err := r.SetPathParam("group", o.Group); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
