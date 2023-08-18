// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1GroupName(params *DeleteV1GroupNameParams, opts ...ClientOption) (*DeleteV1GroupNameOK, error)

	GetV1Group(params *GetV1GroupParams, opts ...ClientOption) (*GetV1GroupOK, error)

	GetV1GroupName(params *GetV1GroupNameParams, opts ...ClientOption) (*GetV1GroupNameOK, error)

	PatchV1GroupName(params *PatchV1GroupNameParams, opts ...ClientOption) (*PatchV1GroupNameOK, error)

	PostV1Group(params *PostV1GroupParams, opts ...ClientOption) (*PostV1GroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1GroupName deletes group
*/
func (a *Client) DeleteV1GroupName(params *DeleteV1GroupNameParams, opts ...ClientOption) (*DeleteV1GroupNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1GroupNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteV1GroupName",
		Method:             "DELETE",
		PathPattern:        "/v1/group/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1GroupNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1GroupNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteV1GroupName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1Group gets a list of groups
*/
func (a *Client) GetV1Group(params *GetV1GroupParams, opts ...ClientOption) (*GetV1GroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1GroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Group",
		Method:             "GET",
		PathPattern:        "/v1/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1GroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1GroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Group: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1GroupName shows group
*/
func (a *Client) GetV1GroupName(params *GetV1GroupNameParams, opts ...ClientOption) (*GetV1GroupNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1GroupNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1GroupName",
		Method:             "GET",
		PathPattern:        "/v1/group/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1GroupNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1GroupNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1GroupName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1GroupName updates group
*/
func (a *Client) PatchV1GroupName(params *PatchV1GroupNameParams, opts ...ClientOption) (*PatchV1GroupNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1GroupNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1GroupName",
		Method:             "PATCH",
		PathPattern:        "/v1/group/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1GroupNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchV1GroupNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1GroupName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1Group creates group
*/
func (a *Client) PostV1Group(params *PostV1GroupParams, opts ...ClientOption) (*PostV1GroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1GroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1Group",
		Method:             "POST",
		PathPattern:        "/v1/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1GroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1GroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1Group: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}