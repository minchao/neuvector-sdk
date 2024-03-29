// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new container API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for container API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Workload(params *GetV1WorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadOK, error)

	GetV1WorkloadID(params *GetV1WorkloadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDOK, error)

	GetV1WorkloadIDCompliance(params *GetV1WorkloadIDComplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDComplianceOK, error)

	GetV1WorkloadIDConfig(params *GetV1WorkloadIDConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDConfigOK, error)

	GetV1WorkloadIDProcess(params *GetV1WorkloadIDProcessParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDProcessOK, error)

	GetV1WorkloadIDProcessHistory(params *GetV1WorkloadIDProcessHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDProcessHistoryOK, error)

	GetV1WorkloadIDStats(params *GetV1WorkloadIDStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDStatsOK, error)

	GetV2Workload(params *GetV2WorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV2WorkloadOK, error)

	GetV2WorkloadID(params *GetV2WorkloadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV2WorkloadIDOK, error)

	PatchV1WorkloadID(params *PatchV1WorkloadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1WorkloadIDOK, error)

	PostV1WorkloadRequestID(params *PostV1WorkloadRequestIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1WorkloadRequestIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1Workload gets container list
*/
func (a *Client) GetV1Workload(params *GetV1WorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Workload",
		Method:             "GET",
		PathPattern:        "/v1/workload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1WorkloadReader{formats: a.formats},
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
	success, ok := result.(*GetV1WorkloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Workload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1WorkloadID gets container detail
*/
func (a *Client) GetV1WorkloadID(params *GetV1WorkloadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WorkloadIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1WorkloadID",
		Method:             "GET",
		PathPattern:        "/v1/workload/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1WorkloadIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1WorkloadIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1WorkloadID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1WorkloadIDCompliance gets a container compliance report
*/
func (a *Client) GetV1WorkloadIDCompliance(params *GetV1WorkloadIDComplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDComplianceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WorkloadIDComplianceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1WorkloadIDCompliance",
		Method:             "GET",
		PathPattern:        "/v1/workload/{id}/compliance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1WorkloadIDComplianceReader{formats: a.formats},
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
	success, ok := result.(*GetV1WorkloadIDComplianceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1WorkloadIDCompliance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1WorkloadIDConfig gets a container configure
*/
func (a *Client) GetV1WorkloadIDConfig(params *GetV1WorkloadIDConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WorkloadIDConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1WorkloadIDConfig",
		Method:             "GET",
		PathPattern:        "/v1/workload/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1WorkloadIDConfigReader{formats: a.formats},
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
	success, ok := result.(*GetV1WorkloadIDConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1WorkloadIDConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1WorkloadIDProcess gets a container process
*/
func (a *Client) GetV1WorkloadIDProcess(params *GetV1WorkloadIDProcessParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDProcessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WorkloadIDProcessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1WorkloadIDProcess",
		Method:             "GET",
		PathPattern:        "/v1/workload/{id}/process",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1WorkloadIDProcessReader{formats: a.formats},
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
	success, ok := result.(*GetV1WorkloadIDProcessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1WorkloadIDProcess: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1WorkloadIDProcessHistory gets a container process history
*/
func (a *Client) GetV1WorkloadIDProcessHistory(params *GetV1WorkloadIDProcessHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDProcessHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WorkloadIDProcessHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1WorkloadIDProcessHistory",
		Method:             "GET",
		PathPattern:        "/v1/workload/{id}/process_history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1WorkloadIDProcessHistoryReader{formats: a.formats},
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
	success, ok := result.(*GetV1WorkloadIDProcessHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1WorkloadIDProcessHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1WorkloadIDStats gets container stats
*/
func (a *Client) GetV1WorkloadIDStats(params *GetV1WorkloadIDStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1WorkloadIDStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WorkloadIDStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1WorkloadIDStats",
		Method:             "GET",
		PathPattern:        "/v1/workload/{id}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1WorkloadIDStatsReader{formats: a.formats},
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
	success, ok := result.(*GetV1WorkloadIDStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1WorkloadIDStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV2Workload gets container list starting from 5 0 rest client should call this api
*/
func (a *Client) GetV2Workload(params *GetV2WorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV2WorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2WorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV2Workload",
		Method:             "GET",
		PathPattern:        "/v2/workload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV2WorkloadReader{formats: a.formats},
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
	success, ok := result.(*GetV2WorkloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV2Workload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV2WorkloadID gets container detail starting from 5 0 rest client should call this api
*/
func (a *Client) GetV2WorkloadID(params *GetV2WorkloadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV2WorkloadIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2WorkloadIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV2WorkloadID",
		Method:             "GET",
		PathPattern:        "/v2/workload/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV2WorkloadIDReader{formats: a.formats},
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
	success, ok := result.(*GetV2WorkloadIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV2WorkloadID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1WorkloadID updates container
*/
func (a *Client) PatchV1WorkloadID(params *PatchV1WorkloadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1WorkloadIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1WorkloadIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1WorkloadID",
		Method:             "PATCH",
		PathPattern:        "/v1/workload/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1WorkloadIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1WorkloadIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1WorkloadID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1WorkloadRequestID containers request
*/
func (a *Client) PostV1WorkloadRequestID(params *PostV1WorkloadRequestIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1WorkloadRequestIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1WorkloadRequestIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1WorkloadRequestID",
		Method:             "POST",
		PathPattern:        "/v1/workload/request/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1WorkloadRequestIDReader{formats: a.formats},
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
	success, ok := result.(*PostV1WorkloadRequestIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1WorkloadRequestID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
