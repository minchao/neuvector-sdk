// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Service(params *GetV1ServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServiceOK, error)

	GetV1ServiceName(params *GetV1ServiceNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServiceNameOK, error)

	PatchV1ServiceConfig(params *PatchV1ServiceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceConfigOK, error)

	PatchV1ServiceConfigNetwork(params *PatchV1ServiceConfigNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceConfigNetworkOK, error)

	PatchV1ServiceConfigProfile(params *PatchV1ServiceConfigProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceConfigProfileOK, error)

	PostV1Service(params *PostV1ServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServiceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1Service gets a list of services
*/
func (a *Client) GetV1Service(params *GetV1ServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Service",
		Method:             "GET",
		PathPattern:        "/v1/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetV1ServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Service: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ServiceName shows service
*/
func (a *Client) GetV1ServiceName(params *GetV1ServiceNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServiceNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServiceNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1ServiceName",
		Method:             "GET",
		PathPattern:        "/v1/service/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServiceNameReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetV1ServiceNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1ServiceName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1ServiceConfig configures service
*/
func (a *Client) PatchV1ServiceConfig(params *PatchV1ServiceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ServiceConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1ServiceConfig",
		Method:             "PATCH",
		PathPattern:        "/v1/service/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ServiceConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PatchV1ServiceConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1ServiceConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1ServiceConfigNetwork configures services in batch
*/
func (a *Client) PatchV1ServiceConfigNetwork(params *PatchV1ServiceConfigNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceConfigNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ServiceConfigNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1ServiceConfigNetwork",
		Method:             "PATCH",
		PathPattern:        "/v1/service/config/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ServiceConfigNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PatchV1ServiceConfigNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1ServiceConfigNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1ServiceConfigProfile configures services in batch
*/
func (a *Client) PatchV1ServiceConfigProfile(params *PatchV1ServiceConfigProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceConfigProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ServiceConfigProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1ServiceConfigProfile",
		Method:             "PATCH",
		PathPattern:        "/v1/service/config/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ServiceConfigProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PatchV1ServiceConfigProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1ServiceConfigProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1Service creates service
*/
func (a *Client) PostV1Service(params *PostV1ServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1Service",
		Method:             "POST",
		PathPattern:        "/v1/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostV1ServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1Service: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
