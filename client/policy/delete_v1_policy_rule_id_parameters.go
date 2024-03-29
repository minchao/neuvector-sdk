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
)

// NewDeleteV1PolicyRuleIDParams creates a new DeleteV1PolicyRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1PolicyRuleIDParams() *DeleteV1PolicyRuleIDParams {
	return &DeleteV1PolicyRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1PolicyRuleIDParamsWithTimeout creates a new DeleteV1PolicyRuleIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1PolicyRuleIDParamsWithTimeout(timeout time.Duration) *DeleteV1PolicyRuleIDParams {
	return &DeleteV1PolicyRuleIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1PolicyRuleIDParamsWithContext creates a new DeleteV1PolicyRuleIDParams object
// with the ability to set a context for a request.
func NewDeleteV1PolicyRuleIDParamsWithContext(ctx context.Context) *DeleteV1PolicyRuleIDParams {
	return &DeleteV1PolicyRuleIDParams{
		Context: ctx,
	}
}

// NewDeleteV1PolicyRuleIDParamsWithHTTPClient creates a new DeleteV1PolicyRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1PolicyRuleIDParamsWithHTTPClient(client *http.Client) *DeleteV1PolicyRuleIDParams {
	return &DeleteV1PolicyRuleIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1PolicyRuleIDParams contains all the parameters to send to the API endpoint

	for the delete v1 policy rule ID operation.

	Typically these are written to a http.Request.
*/
type DeleteV1PolicyRuleIDParams struct {

	/* ID.

	   Rule ID

	   Format: uint32
	*/
	ID uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 policy rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PolicyRuleIDParams) WithDefaults() *DeleteV1PolicyRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 policy rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PolicyRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) WithTimeout(timeout time.Duration) *DeleteV1PolicyRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) WithContext(ctx context.Context) *DeleteV1PolicyRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) WithHTTPClient(client *http.Client) *DeleteV1PolicyRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) WithID(id uint32) *DeleteV1PolicyRuleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 policy rule ID params
func (o *DeleteV1PolicyRuleIDParams) SetID(id uint32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1PolicyRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatUint32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
