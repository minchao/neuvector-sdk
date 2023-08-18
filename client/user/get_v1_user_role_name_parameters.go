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
)

// NewGetV1UserRoleNameParams creates a new GetV1UserRoleNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1UserRoleNameParams() *GetV1UserRoleNameParams {
	return &GetV1UserRoleNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1UserRoleNameParamsWithTimeout creates a new GetV1UserRoleNameParams object
// with the ability to set a timeout on a request.
func NewGetV1UserRoleNameParamsWithTimeout(timeout time.Duration) *GetV1UserRoleNameParams {
	return &GetV1UserRoleNameParams{
		timeout: timeout,
	}
}

// NewGetV1UserRoleNameParamsWithContext creates a new GetV1UserRoleNameParams object
// with the ability to set a context for a request.
func NewGetV1UserRoleNameParamsWithContext(ctx context.Context) *GetV1UserRoleNameParams {
	return &GetV1UserRoleNameParams{
		Context: ctx,
	}
}

// NewGetV1UserRoleNameParamsWithHTTPClient creates a new GetV1UserRoleNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1UserRoleNameParamsWithHTTPClient(client *http.Client) *GetV1UserRoleNameParams {
	return &GetV1UserRoleNameParams{
		HTTPClient: client,
	}
}

/*
GetV1UserRoleNameParams contains all the parameters to send to the API endpoint

	for the get v1 user role name operation.

	Typically these are written to a http.Request.
*/
type GetV1UserRoleNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   User role name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 user role name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UserRoleNameParams) WithDefaults() *GetV1UserRoleNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 user role name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UserRoleNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 user role name params
func (o *GetV1UserRoleNameParams) WithTimeout(timeout time.Duration) *GetV1UserRoleNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 user role name params
func (o *GetV1UserRoleNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 user role name params
func (o *GetV1UserRoleNameParams) WithContext(ctx context.Context) *GetV1UserRoleNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 user role name params
func (o *GetV1UserRoleNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 user role name params
func (o *GetV1UserRoleNameParams) WithHTTPClient(client *http.Client) *GetV1UserRoleNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 user role name params
func (o *GetV1UserRoleNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 user role name params
func (o *GetV1UserRoleNameParams) WithXAuthToken(xAuthToken string) *GetV1UserRoleNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 user role name params
func (o *GetV1UserRoleNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the get v1 user role name params
func (o *GetV1UserRoleNameParams) WithName(name string) *GetV1UserRoleNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 user role name params
func (o *GetV1UserRoleNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1UserRoleNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
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