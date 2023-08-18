// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTUnquarReq r e s t unquar req
//
// swagger:model RESTUnquarReq
type RESTUnquarReq struct {

	// group
	Group string `json:"group,omitempty"`

	// response rule
	// Example: 1007
	ResponseRule uint32 `json:"response_rule,omitempty"`
}

// Validate validates this r e s t unquar req
func (m *RESTUnquarReq) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this r e s t unquar req based on context it is used
func (m *RESTUnquarReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTUnquarReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTUnquarReq) UnmarshalBinary(b []byte) error {
	var res RESTUnquarReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
