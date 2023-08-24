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

// NewGetV1ResponseWorkloadRulesIDParams creates a new GetV1ResponseWorkloadRulesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ResponseWorkloadRulesIDParams() *GetV1ResponseWorkloadRulesIDParams {
	return &GetV1ResponseWorkloadRulesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ResponseWorkloadRulesIDParamsWithTimeout creates a new GetV1ResponseWorkloadRulesIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ResponseWorkloadRulesIDParamsWithTimeout(timeout time.Duration) *GetV1ResponseWorkloadRulesIDParams {
	return &GetV1ResponseWorkloadRulesIDParams{
		timeout: timeout,
	}
}

// NewGetV1ResponseWorkloadRulesIDParamsWithContext creates a new GetV1ResponseWorkloadRulesIDParams object
// with the ability to set a context for a request.
func NewGetV1ResponseWorkloadRulesIDParamsWithContext(ctx context.Context) *GetV1ResponseWorkloadRulesIDParams {
	return &GetV1ResponseWorkloadRulesIDParams{
		Context: ctx,
	}
}

// NewGetV1ResponseWorkloadRulesIDParamsWithHTTPClient creates a new GetV1ResponseWorkloadRulesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ResponseWorkloadRulesIDParamsWithHTTPClient(client *http.Client) *GetV1ResponseWorkloadRulesIDParams {
	return &GetV1ResponseWorkloadRulesIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ResponseWorkloadRulesIDParams contains all the parameters to send to the API endpoint

	for the get v1 response workload rules ID operation.

	Typically these are written to a http.Request.
*/
type GetV1ResponseWorkloadRulesIDParams struct {

	/* ID.

	   Workload rules ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 response workload rules ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ResponseWorkloadRulesIDParams) WithDefaults() *GetV1ResponseWorkloadRulesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 response workload rules ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ResponseWorkloadRulesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) WithTimeout(timeout time.Duration) *GetV1ResponseWorkloadRulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) WithContext(ctx context.Context) *GetV1ResponseWorkloadRulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) WithHTTPClient(client *http.Client) *GetV1ResponseWorkloadRulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) WithID(id string) *GetV1ResponseWorkloadRulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 response workload rules ID params
func (o *GetV1ResponseWorkloadRulesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ResponseWorkloadRulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
