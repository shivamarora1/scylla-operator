// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// FailureDetectorCountEndpointDownGetReader is a Reader for the FailureDetectorCountEndpointDownGet structure.
type FailureDetectorCountEndpointDownGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FailureDetectorCountEndpointDownGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFailureDetectorCountEndpointDownGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFailureDetectorCountEndpointDownGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFailureDetectorCountEndpointDownGetOK creates a FailureDetectorCountEndpointDownGetOK with default headers values
func NewFailureDetectorCountEndpointDownGetOK() *FailureDetectorCountEndpointDownGetOK {
	return &FailureDetectorCountEndpointDownGetOK{}
}

/*
FailureDetectorCountEndpointDownGetOK handles this case with default header values.

FailureDetectorCountEndpointDownGetOK failure detector count endpoint down get o k
*/
type FailureDetectorCountEndpointDownGetOK struct {
	Payload int32
}

func (o *FailureDetectorCountEndpointDownGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *FailureDetectorCountEndpointDownGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFailureDetectorCountEndpointDownGetDefault creates a FailureDetectorCountEndpointDownGetDefault with default headers values
func NewFailureDetectorCountEndpointDownGetDefault(code int) *FailureDetectorCountEndpointDownGetDefault {
	return &FailureDetectorCountEndpointDownGetDefault{
		_statusCode: code,
	}
}

/*
FailureDetectorCountEndpointDownGetDefault handles this case with default header values.

internal server error
*/
type FailureDetectorCountEndpointDownGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the failure detector count endpoint down get default response
func (o *FailureDetectorCountEndpointDownGetDefault) Code() int {
	return o._statusCode
}

func (o *FailureDetectorCountEndpointDownGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FailureDetectorCountEndpointDownGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FailureDetectorCountEndpointDownGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
