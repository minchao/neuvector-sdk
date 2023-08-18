// Code generated by go-swagger; DO NOT EDIT.

package response_rule

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

// NewDeleteV1ResponseRuleIDParams creates a new DeleteV1ResponseRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ResponseRuleIDParams() *DeleteV1ResponseRuleIDParams {
	return &DeleteV1ResponseRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ResponseRuleIDParamsWithTimeout creates a new DeleteV1ResponseRuleIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ResponseRuleIDParamsWithTimeout(timeout time.Duration) *DeleteV1ResponseRuleIDParams {
	return &DeleteV1ResponseRuleIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1ResponseRuleIDParamsWithContext creates a new DeleteV1ResponseRuleIDParams object
// with the ability to set a context for a request.
func NewDeleteV1ResponseRuleIDParamsWithContext(ctx context.Context) *DeleteV1ResponseRuleIDParams {
	return &DeleteV1ResponseRuleIDParams{
		Context: ctx,
	}
}

// NewDeleteV1ResponseRuleIDParamsWithHTTPClient creates a new DeleteV1ResponseRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ResponseRuleIDParamsWithHTTPClient(client *http.Client) *DeleteV1ResponseRuleIDParams {
	return &DeleteV1ResponseRuleIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ResponseRuleIDParams contains all the parameters to send to the API endpoint

	for the delete v1 response rule ID operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ResponseRuleIDParams struct {

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

// WithDefaults hydrates default values in the delete v1 response rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ResponseRuleIDParams) WithDefaults() *DeleteV1ResponseRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 response rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ResponseRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) WithTimeout(timeout time.Duration) *DeleteV1ResponseRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) WithContext(ctx context.Context) *DeleteV1ResponseRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) WithHTTPClient(client *http.Client) *DeleteV1ResponseRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) WithXAuthToken(xAuthToken string) *DeleteV1ResponseRuleIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithID adds the id to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) WithID(id string) *DeleteV1ResponseRuleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 response rule ID params
func (o *DeleteV1ResponseRuleIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ResponseRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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