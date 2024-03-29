// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Controller(params *GetV1ControllerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerOK, error)

	GetV1ControllerID(params *GetV1ControllerIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerIDOK, error)

	GetV1ControllerIDConfig(params *GetV1ControllerIDConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerIDConfigOK, error)

	GetV1ControllerIDStats(params *GetV1ControllerIDStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerIDStatsOK, error)

	PatchV1ControllerID(params *PatchV1ControllerIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ControllerIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1Controller gets a list of controllers
*/
func (a *Client) GetV1Controller(params *GetV1ControllerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ControllerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Controller",
		Method:             "GET",
		PathPattern:        "/v1/controller",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ControllerReader{formats: a.formats},
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
	success, ok := result.(*GetV1ControllerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ControllerID shows controller
*/
func (a *Client) GetV1ControllerID(params *GetV1ControllerIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ControllerIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1ControllerID",
		Method:             "GET",
		PathPattern:        "/v1/controller/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ControllerIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1ControllerIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1ControllerID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ControllerIDConfig controllers get configure
*/
func (a *Client) GetV1ControllerIDConfig(params *GetV1ControllerIDConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerIDConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ControllerIDConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1ControllerIDConfig",
		Method:             "GET",
		PathPattern:        "/v1/controller/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ControllerIDConfigReader{formats: a.formats},
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
	success, ok := result.(*GetV1ControllerIDConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1ControllerIDConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ControllerIDStats controllers get system statistics
*/
func (a *Client) GetV1ControllerIDStats(params *GetV1ControllerIDStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ControllerIDStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ControllerIDStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1ControllerIDStats",
		Method:             "GET",
		PathPattern:        "/v1/controller/{id}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ControllerIDStatsReader{formats: a.formats},
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
	success, ok := result.(*GetV1ControllerIDStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1ControllerIDStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1ControllerID updates controller
*/
func (a *Client) PatchV1ControllerID(params *PatchV1ControllerIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ControllerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ControllerIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1ControllerID",
		Method:             "PATCH",
		PathPattern:        "/v1/controller/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ControllerIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1ControllerIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1ControllerID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
