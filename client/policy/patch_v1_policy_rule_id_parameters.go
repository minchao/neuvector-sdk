// Code generated by go-swagger; DO NOT EDIT.

package policy

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
	"github.com/go-openapi/swag"

	"github.com/minchao/neuvector-sdk/models"
)

// NewPatchV1PolicyRuleIDParams creates a new PatchV1PolicyRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1PolicyRuleIDParams() *PatchV1PolicyRuleIDParams {
	return &PatchV1PolicyRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PolicyRuleIDParamsWithTimeout creates a new PatchV1PolicyRuleIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1PolicyRuleIDParamsWithTimeout(timeout time.Duration) *PatchV1PolicyRuleIDParams {
	return &PatchV1PolicyRuleIDParams{
		timeout: timeout,
	}
}

// NewPatchV1PolicyRuleIDParamsWithContext creates a new PatchV1PolicyRuleIDParams object
// with the ability to set a context for a request.
func NewPatchV1PolicyRuleIDParamsWithContext(ctx context.Context) *PatchV1PolicyRuleIDParams {
	return &PatchV1PolicyRuleIDParams{
		Context: ctx,
	}
}

// NewPatchV1PolicyRuleIDParamsWithHTTPClient creates a new PatchV1PolicyRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1PolicyRuleIDParamsWithHTTPClient(client *http.Client) *PatchV1PolicyRuleIDParams {
	return &PatchV1PolicyRuleIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1PolicyRuleIDParams contains all the parameters to send to the API endpoint

	for the patch v1 policy rule ID operation.

	Typically these are written to a http.Request.
*/
type PatchV1PolicyRuleIDParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Policy rule update data
	*/
	Body *models.RESTPolicyRuleConfigData

	/* ID.

	   Rule ID

	   Format: uint32
	*/
	ID uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 policy rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PolicyRuleIDParams) WithDefaults() *PatchV1PolicyRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 policy rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PolicyRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) WithTimeout(timeout time.Duration) *PatchV1PolicyRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) WithContext(ctx context.Context) *PatchV1PolicyRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) WithHTTPClient(client *http.Client) *PatchV1PolicyRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) WithXAuthToken(xAuthToken string) *PatchV1PolicyRuleIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) WithBody(body *models.RESTPolicyRuleConfigData) *PatchV1PolicyRuleIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) SetBody(body *models.RESTPolicyRuleConfigData) {
	o.Body = body
}

// WithID adds the id to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) WithID(id uint32) *PatchV1PolicyRuleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch v1 policy rule ID params
func (o *PatchV1PolicyRuleIDParams) SetID(id uint32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PolicyRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatUint32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
