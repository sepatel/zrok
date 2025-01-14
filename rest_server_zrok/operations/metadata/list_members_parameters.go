// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListMembersParams creates a new ListMembersParams object
//
// There are no default values defined in the spec.
func NewListMembersParams() ListMembersParams {

	return ListMembersParams{}
}

// ListMembersParams contains all the bound params for the list members operation
// typically these are obtained from a http.Request
//
// swagger:parameters listMembers
type ListMembersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	OrganizationToken string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListMembersParams() beforehand.
func (o *ListMembersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rOrganizationToken, rhkOrganizationToken, _ := route.Params.GetOK("organizationToken")
	if err := o.bindOrganizationToken(rOrganizationToken, rhkOrganizationToken, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOrganizationToken binds and validates parameter OrganizationToken from path.
func (o *ListMembersParams) bindOrganizationToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.OrganizationToken = raw

	return nil
}