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

// FindConfigClientEncryptionOptionsReader is a Reader for the FindConfigClientEncryptionOptions structure.
type FindConfigClientEncryptionOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigClientEncryptionOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigClientEncryptionOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigClientEncryptionOptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigClientEncryptionOptionsOK creates a FindConfigClientEncryptionOptionsOK with default headers values
func NewFindConfigClientEncryptionOptionsOK() *FindConfigClientEncryptionOptionsOK {
	return &FindConfigClientEncryptionOptionsOK{}
}

/*
FindConfigClientEncryptionOptionsOK handles this case with default header values.

Config value
*/
type FindConfigClientEncryptionOptionsOK struct {
	Payload []string
}

func (o *FindConfigClientEncryptionOptionsOK) GetPayload() []string {
	return o.Payload
}

func (o *FindConfigClientEncryptionOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigClientEncryptionOptionsDefault creates a FindConfigClientEncryptionOptionsDefault with default headers values
func NewFindConfigClientEncryptionOptionsDefault(code int) *FindConfigClientEncryptionOptionsDefault {
	return &FindConfigClientEncryptionOptionsDefault{
		_statusCode: code,
	}
}

/*
FindConfigClientEncryptionOptionsDefault handles this case with default header values.

unexpected error
*/
type FindConfigClientEncryptionOptionsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config client encryption options default response
func (o *FindConfigClientEncryptionOptionsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigClientEncryptionOptionsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigClientEncryptionOptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigClientEncryptionOptionsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
