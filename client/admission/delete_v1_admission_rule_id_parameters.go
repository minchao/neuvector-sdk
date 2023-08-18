// Code generated by go-swagger; DO NOT EDIT.

package admission

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

// NewDeleteV1AdmissionRuleIDParams creates a new DeleteV1AdmissionRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1AdmissionRuleIDParams() *DeleteV1AdmissionRuleIDParams {
	return &DeleteV1AdmissionRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1AdmissionRuleIDParamsWithTimeout creates a new DeleteV1AdmissionRuleIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1AdmissionRuleIDParamsWithTimeout(timeout time.Duration) *DeleteV1AdmissionRuleIDParams {
	return &DeleteV1AdmissionRuleIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1AdmissionRuleIDParamsWithContext creates a new DeleteV1AdmissionRuleIDParams object
// with the ability to set a context for a request.
func NewDeleteV1AdmissionRuleIDParamsWithContext(ctx context.Context) *DeleteV1AdmissionRuleIDParams {
	return &DeleteV1AdmissionRuleIDParams{
		Context: ctx,
	}
}

// NewDeleteV1AdmissionRuleIDParamsWithHTTPClient creates a new DeleteV1AdmissionRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1AdmissionRuleIDParamsWithHTTPClient(client *http.Client) *DeleteV1AdmissionRuleIDParams {
	return &DeleteV1AdmissionRuleIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1AdmissionRuleIDParams contains all the parameters to send to the API endpoint

	for the delete v1 admission rule ID operation.

	Typically these are written to a http.Request.
*/
type DeleteV1AdmissionRuleIDParams struct {

	// XAuthToken.
	XAuthToken string

	/* ID.

	   Rule ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 admission rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1AdmissionRuleIDParams) WithDefaults() *DeleteV1AdmissionRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 admission rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1AdmissionRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) WithTimeout(timeout time.Duration) *DeleteV1AdmissionRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) WithContext(ctx context.Context) *DeleteV1AdmissionRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) WithHTTPClient(client *http.Client) *DeleteV1AdmissionRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) WithXAuthToken(xAuthToken string) *DeleteV1AdmissionRuleIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithID adds the id to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) WithID(id string) *DeleteV1AdmissionRuleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 admission rule ID params
func (o *DeleteV1AdmissionRuleIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1AdmissionRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
