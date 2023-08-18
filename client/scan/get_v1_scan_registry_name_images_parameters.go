// Code generated by go-swagger; DO NOT EDIT.

package scan

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

// NewGetV1ScanRegistryNameImagesParams creates a new GetV1ScanRegistryNameImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ScanRegistryNameImagesParams() *GetV1ScanRegistryNameImagesParams {
	return &GetV1ScanRegistryNameImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ScanRegistryNameImagesParamsWithTimeout creates a new GetV1ScanRegistryNameImagesParams object
// with the ability to set a timeout on a request.
func NewGetV1ScanRegistryNameImagesParamsWithTimeout(timeout time.Duration) *GetV1ScanRegistryNameImagesParams {
	return &GetV1ScanRegistryNameImagesParams{
		timeout: timeout,
	}
}

// NewGetV1ScanRegistryNameImagesParamsWithContext creates a new GetV1ScanRegistryNameImagesParams object
// with the ability to set a context for a request.
func NewGetV1ScanRegistryNameImagesParamsWithContext(ctx context.Context) *GetV1ScanRegistryNameImagesParams {
	return &GetV1ScanRegistryNameImagesParams{
		Context: ctx,
	}
}

// NewGetV1ScanRegistryNameImagesParamsWithHTTPClient creates a new GetV1ScanRegistryNameImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ScanRegistryNameImagesParamsWithHTTPClient(client *http.Client) *GetV1ScanRegistryNameImagesParams {
	return &GetV1ScanRegistryNameImagesParams{
		HTTPClient: client,
	}
}

/*
GetV1ScanRegistryNameImagesParams contains all the parameters to send to the API endpoint

	for the get v1 scan registry name images operation.

	Typically these are written to a http.Request.
*/
type GetV1ScanRegistryNameImagesParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   Name of the registry
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 scan registry name images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanRegistryNameImagesParams) WithDefaults() *GetV1ScanRegistryNameImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 scan registry name images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanRegistryNameImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) WithTimeout(timeout time.Duration) *GetV1ScanRegistryNameImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) WithContext(ctx context.Context) *GetV1ScanRegistryNameImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) WithHTTPClient(client *http.Client) *GetV1ScanRegistryNameImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) WithXAuthToken(xAuthToken string) *GetV1ScanRegistryNameImagesParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) WithName(name string) *GetV1ScanRegistryNameImagesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 scan registry name images params
func (o *GetV1ScanRegistryNameImagesParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ScanRegistryNameImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
