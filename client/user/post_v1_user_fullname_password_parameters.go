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

// NewPostV1UserFullnamePasswordParams creates a new PostV1UserFullnamePasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1UserFullnamePasswordParams() *PostV1UserFullnamePasswordParams {
	return &PostV1UserFullnamePasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1UserFullnamePasswordParamsWithTimeout creates a new PostV1UserFullnamePasswordParams object
// with the ability to set a timeout on a request.
func NewPostV1UserFullnamePasswordParamsWithTimeout(timeout time.Duration) *PostV1UserFullnamePasswordParams {
	return &PostV1UserFullnamePasswordParams{
		timeout: timeout,
	}
}

// NewPostV1UserFullnamePasswordParamsWithContext creates a new PostV1UserFullnamePasswordParams object
// with the ability to set a context for a request.
func NewPostV1UserFullnamePasswordParamsWithContext(ctx context.Context) *PostV1UserFullnamePasswordParams {
	return &PostV1UserFullnamePasswordParams{
		Context: ctx,
	}
}

// NewPostV1UserFullnamePasswordParamsWithHTTPClient creates a new PostV1UserFullnamePasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1UserFullnamePasswordParamsWithHTTPClient(client *http.Client) *PostV1UserFullnamePasswordParams {
	return &PostV1UserFullnamePasswordParams{
		HTTPClient: client,
	}
}

/*
PostV1UserFullnamePasswordParams contains all the parameters to send to the API endpoint

	for the post v1 user fullname password operation.

	Typically these are written to a http.Request.
*/
type PostV1UserFullnamePasswordParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   User password configuration data
	*/
	Body *models.RESTUserPwdConfigData

	/* Fullname.

	   User name
	*/
	Fullname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 user fullname password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1UserFullnamePasswordParams) WithDefaults() *PostV1UserFullnamePasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 user fullname password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1UserFullnamePasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) WithTimeout(timeout time.Duration) *PostV1UserFullnamePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) WithContext(ctx context.Context) *PostV1UserFullnamePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) WithHTTPClient(client *http.Client) *PostV1UserFullnamePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) WithXAuthToken(xAuthToken string) *PostV1UserFullnamePasswordParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) WithBody(body *models.RESTUserPwdConfigData) *PostV1UserFullnamePasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) SetBody(body *models.RESTUserPwdConfigData) {
	o.Body = body
}

// WithFullname adds the fullname to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) WithFullname(fullname string) *PostV1UserFullnamePasswordParams {
	o.SetFullname(fullname)
	return o
}

// SetFullname adds the fullname to the post v1 user fullname password params
func (o *PostV1UserFullnamePasswordParams) SetFullname(fullname string) {
	o.Fullname = fullname
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1UserFullnamePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param fullname
	if err := r.SetPathParam("fullname", o.Fullname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}