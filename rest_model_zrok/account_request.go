// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountRequest account request
//
// swagger:model accountRequest
type AccountRequest struct {

	// email
	Email string `json:"email,omitempty"`
}

// Validate validates this account request
func (m *AccountRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account request based on context it is used
func (m *AccountRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequest) UnmarshalBinary(b []byte) error {
	var res AccountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
