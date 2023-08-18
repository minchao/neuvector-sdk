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

// NewPostV1UserRoleParams creates a new PostV1UserRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1UserRoleParams() *PostV1UserRoleParams {
	return &PostV1UserRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1UserRoleParamsWithTimeout creates a new PostV1UserRoleParams object
// with the ability to set a timeout on a request.
func NewPostV1UserRoleParamsWithTimeout(timeout time.Duration) *PostV1UserRoleParams {
	return &PostV1UserRoleParams{
		timeout: timeout,
	}
}

// NewPostV1UserRoleParamsWithContext creates a new PostV1UserRoleParams object
// with the ability to set a context for a request.
func NewPostV1UserRoleParamsWithContext(ctx context.Context) *PostV1UserRoleParams {
	return &PostV1UserRoleParams{
		Context: ctx,
	}
}

// NewPostV1UserRoleParamsWithHTTPClient creates a new PostV1UserRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1UserRoleParamsWithHTTPClient(client *http.Client) *PostV1UserRoleParams {
	return &PostV1UserRoleParams{
		HTTPClient: client,
	}
}

/*
PostV1UserRoleParams contains all the parameters to send to the API endpoint

	for the post v1 user role operation.

	Typically these are written to a http.Request.
*/
type PostV1UserRoleParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Role information
	*/
	Body *models.RESTUserRoleConfigData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 user role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1UserRoleParams) WithDefaults() *PostV1UserRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 user role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1UserRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 user role params
func (o *PostV1UserRoleParams) WithTimeout(timeout time.Duration) *PostV1UserRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 user role params
func (o *PostV1UserRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 user role params
func (o *PostV1UserRoleParams) WithContext(ctx context.Context) *PostV1UserRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 user role params
func (o *PostV1UserRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 user role params
func (o *PostV1UserRoleParams) WithHTTPClient(client *http.Client) *PostV1UserRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 user role params
func (o *PostV1UserRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the post v1 user role params
func (o *PostV1UserRoleParams) WithXAuthToken(xAuthToken string) *PostV1UserRoleParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post v1 user role params
func (o *PostV1UserRoleParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the post v1 user role params
func (o *PostV1UserRoleParams) WithBody(body *models.RESTUserRoleConfigData) *PostV1UserRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 user role params
func (o *PostV1UserRoleParams) SetBody(body *models.RESTUserRoleConfigData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1UserRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}