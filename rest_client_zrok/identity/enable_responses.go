// Code generated by go-swagger; DO NOT EDIT.

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// EnableReader is a Reader for the Enable structure.
type EnableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEnableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEnableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableCreated creates a EnableCreated with default headers values
func NewEnableCreated() *EnableCreated {
	return &EnableCreated{}
}

/* EnableCreated describes a response with status code 201, with default header values.

environment enabled
*/
type EnableCreated struct {
	Payload *rest_model_zrok.EnableResponse
}

func (o *EnableCreated) Error() string {
	return fmt.Sprintf("[POST /enable][%d] enableCreated  %+v", 201, o.Payload)
}
func (o *EnableCreated) GetPayload() *rest_model_zrok.EnableResponse {
	return o.Payload
}

func (o *EnableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model_zrok.EnableResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableNotFound creates a EnableNotFound with default headers values
func NewEnableNotFound() *EnableNotFound {
	return &EnableNotFound{}
}

/* EnableNotFound describes a response with status code 404, with default header values.

account not found
*/
type EnableNotFound struct {
}

func (o *EnableNotFound) Error() string {
	return fmt.Sprintf("[POST /enable][%d] enableNotFound ", 404)
}

func (o *EnableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableInternalServerError creates a EnableInternalServerError with default headers values
func NewEnableInternalServerError() *EnableInternalServerError {
	return &EnableInternalServerError{}
}

/* EnableInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type EnableInternalServerError struct {
	Payload rest_model_zrok.ErrorMessage
}

func (o *EnableInternalServerError) Error() string {
	return fmt.Sprintf("[POST /enable][%d] enableInternalServerError  %+v", 500, o.Payload)
}
func (o *EnableInternalServerError) GetPayload() rest_model_zrok.ErrorMessage {
	return o.Payload
}

func (o *EnableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
