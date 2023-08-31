package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TokenWriter implements runtime.ClientAuthInfoWriter interface.
type TokenWriter struct {
	Token string
}

func NewTokenWriter(token string) *TokenWriter {
	return &TokenWriter{Token: token}
}

func (a TokenWriter) AuthenticateRequest(req runtime.ClientRequest, _ strfmt.Registry) error {
	return req.SetHeaderParam("X-Auth-Token", a.Token)
}
