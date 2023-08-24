// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1UserFullname(params *DeleteV1UserFullnameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1UserFullnameOK, error)

	DeleteV1UserRoleName(params *DeleteV1UserRoleNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1UserRoleNameOK, error)

	GetV1PasswordProfile(params *GetV1PasswordProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1PasswordProfileOK, error)

	GetV1PasswordProfileName(params *GetV1PasswordProfileNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1PasswordProfileNameOK, error)

	GetV1User(params *GetV1UserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserOK, error)

	GetV1UserFullname(params *GetV1UserFullnameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserFullnameOK, error)

	GetV1UserRole(params *GetV1UserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserRoleOK, error)

	GetV1UserRoleName(params *GetV1UserRoleNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserRoleNameOK, error)

	PatchV1PasswordProfileName(params *PatchV1PasswordProfileNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1PasswordProfileNameOK, error)

	PatchV1UserFullname(params *PatchV1UserFullnameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1UserFullnameOK, error)

	PatchV1UserFullnameRoleRole(params *PatchV1UserFullnameRoleRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1UserFullnameRoleRoleOK, error)

	PatchV1UserRoleName(params *PatchV1UserRoleNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1UserRoleNameOK, error)

	PostV1User(params *PostV1UserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1UserOK, error)

	PostV1UserFullnamePassword(params *PostV1UserFullnamePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1UserFullnamePasswordOK, error)

	PostV1UserRole(params *PostV1UserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1UserRoleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1UserFullname deletes user
*/
func (a *Client) DeleteV1UserFullname(params *DeleteV1UserFullnameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1UserFullnameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1UserFullnameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteV1UserFullname",
		Method:             "DELETE",
		PathPattern:        "/v1/user/{fullname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1UserFullnameReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1UserFullnameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteV1UserFullname: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteV1UserRoleName deletes a user role
*/
func (a *Client) DeleteV1UserRoleName(params *DeleteV1UserRoleNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1UserRoleNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1UserRoleNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteV1UserRoleName",
		Method:             "DELETE",
		PathPattern:        "/v1/user_role/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1UserRoleNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1UserRoleNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteV1UserRoleName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1PasswordProfile gets password profile list
*/
func (a *Client) GetV1PasswordProfile(params *GetV1PasswordProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1PasswordProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1PasswordProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1PasswordProfile",
		Method:             "GET",
		PathPattern:        "/v1/password_profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1PasswordProfileReader{formats: a.formats},
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
	success, ok := result.(*GetV1PasswordProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1PasswordProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1PasswordProfileName gets password profile
*/
func (a *Client) GetV1PasswordProfileName(params *GetV1PasswordProfileNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1PasswordProfileNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1PasswordProfileNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1PasswordProfileName",
		Method:             "GET",
		PathPattern:        "/v1/password_profile/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1PasswordProfileNameReader{formats: a.formats},
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
	success, ok := result.(*GetV1PasswordProfileNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1PasswordProfileName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1User gets a list of users
*/
func (a *Client) GetV1User(params *GetV1UserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1UserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1User",
		Method:             "GET",
		PathPattern:        "/v1/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1UserReader{formats: a.formats},
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
	success, ok := result.(*GetV1UserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1User: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1UserFullname gets a user
*/
func (a *Client) GetV1UserFullname(params *GetV1UserFullnameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserFullnameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1UserFullnameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1UserFullname",
		Method:             "GET",
		PathPattern:        "/v1/user/{fullname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1UserFullnameReader{formats: a.formats},
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
	success, ok := result.(*GetV1UserFullnameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1UserFullname: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1UserRole gets role list
*/
func (a *Client) GetV1UserRole(params *GetV1UserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1UserRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1UserRole",
		Method:             "GET",
		PathPattern:        "/v1/user_role",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1UserRoleReader{formats: a.formats},
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
	success, ok := result.(*GetV1UserRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1UserRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1UserRoleName gets role details
*/
func (a *Client) GetV1UserRoleName(params *GetV1UserRoleNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UserRoleNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1UserRoleNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1UserRoleName",
		Method:             "GET",
		PathPattern:        "/v1/user_role/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1UserRoleNameReader{formats: a.formats},
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
	success, ok := result.(*GetV1UserRoleNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1UserRoleName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1PasswordProfileName configures password profile
*/
func (a *Client) PatchV1PasswordProfileName(params *PatchV1PasswordProfileNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1PasswordProfileNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1PasswordProfileNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1PasswordProfileName",
		Method:             "PATCH",
		PathPattern:        "/v1/password_profile/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1PasswordProfileNameReader{formats: a.formats},
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
	success, ok := result.(*PatchV1PasswordProfileNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1PasswordProfileName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1UserFullname updates user
*/
func (a *Client) PatchV1UserFullname(params *PatchV1UserFullnameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1UserFullnameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1UserFullnameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1UserFullname",
		Method:             "PATCH",
		PathPattern:        "/v1/user/{fullname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1UserFullnameReader{formats: a.formats},
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
	success, ok := result.(*PatchV1UserFullnameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1UserFullname: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1UserFullnameRoleRole fors c l i to modify one role
*/
func (a *Client) PatchV1UserFullnameRoleRole(params *PatchV1UserFullnameRoleRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1UserFullnameRoleRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1UserFullnameRoleRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1UserFullnameRoleRole",
		Method:             "PATCH",
		PathPattern:        "/v1/user/{fullname}/role/{role}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1UserFullnameRoleRoleReader{formats: a.formats},
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
	success, ok := result.(*PatchV1UserFullnameRoleRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1UserFullnameRoleRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1UserRoleName configs a user role
*/
func (a *Client) PatchV1UserRoleName(params *PatchV1UserRoleNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1UserRoleNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1UserRoleNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1UserRoleName",
		Method:             "PATCH",
		PathPattern:        "/v1/user_role/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1UserRoleNameReader{formats: a.formats},
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
	success, ok := result.(*PatchV1UserRoleNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1UserRoleName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1User creates a user
*/
func (a *Client) PostV1User(params *PostV1UserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1UserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1UserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1User",
		Method:             "POST",
		PathPattern:        "/v1/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1UserReader{formats: a.formats},
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
	success, ok := result.(*PostV1UserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1User: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1UserFullnamePassword configures user login
*/
func (a *Client) PostV1UserFullnamePassword(params *PostV1UserFullnamePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1UserFullnamePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1UserFullnamePasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1UserFullnamePassword",
		Method:             "POST",
		PathPattern:        "/v1/user/{fullname}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1UserFullnamePasswordReader{formats: a.formats},
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
	success, ok := result.(*PostV1UserFullnamePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1UserFullnamePassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1UserRole creates a role
*/
func (a *Client) PostV1UserRole(params *PostV1UserRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1UserRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1UserRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1UserRole",
		Method:             "POST",
		PathPattern:        "/v1/user_role",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1UserRoleReader{formats: a.formats},
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
	success, ok := result.(*PostV1UserRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1UserRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
