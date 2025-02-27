// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigInitialTokenReader is a Reader for the FindConfigInitialToken structure.
type FindConfigInitialTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigInitialTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigInitialTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigInitialTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigInitialTokenOK creates a FindConfigInitialTokenOK with default headers values
func NewFindConfigInitialTokenOK() *FindConfigInitialTokenOK {
	return &FindConfigInitialTokenOK{}
}

/*
FindConfigInitialTokenOK handles this case with default header values.

Config value
*/
type FindConfigInitialTokenOK struct {
	Payload string
}

func (o *FindConfigInitialTokenOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigInitialTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigInitialTokenDefault creates a FindConfigInitialTokenDefault with default headers values
func NewFindConfigInitialTokenDefault(code int) *FindConfigInitialTokenDefault {
	return &FindConfigInitialTokenDefault{
		_statusCode: code,
	}
}

/*
FindConfigInitialTokenDefault handles this case with default header values.

unexpected error
*/
type FindConfigInitialTokenDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config initial token default response
func (o *FindConfigInitialTokenDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigInitialTokenDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigInitialTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigInitialTokenDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
