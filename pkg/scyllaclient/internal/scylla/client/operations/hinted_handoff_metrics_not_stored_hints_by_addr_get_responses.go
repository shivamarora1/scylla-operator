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

// HintedHandoffMetricsNotStoredHintsByAddrGetReader is a Reader for the HintedHandoffMetricsNotStoredHintsByAddrGet structure.
type HintedHandoffMetricsNotStoredHintsByAddrGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHintedHandoffMetricsNotStoredHintsByAddrGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHintedHandoffMetricsNotStoredHintsByAddrGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHintedHandoffMetricsNotStoredHintsByAddrGetOK creates a HintedHandoffMetricsNotStoredHintsByAddrGetOK with default headers values
func NewHintedHandoffMetricsNotStoredHintsByAddrGetOK() *HintedHandoffMetricsNotStoredHintsByAddrGetOK {
	return &HintedHandoffMetricsNotStoredHintsByAddrGetOK{}
}

/*
HintedHandoffMetricsNotStoredHintsByAddrGetOK handles this case with default header values.

HintedHandoffMetricsNotStoredHintsByAddrGetOK hinted handoff metrics not stored hints by addr get o k
*/
type HintedHandoffMetricsNotStoredHintsByAddrGetOK struct {
	Payload int32
}

func (o *HintedHandoffMetricsNotStoredHintsByAddrGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *HintedHandoffMetricsNotStoredHintsByAddrGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHintedHandoffMetricsNotStoredHintsByAddrGetDefault creates a HintedHandoffMetricsNotStoredHintsByAddrGetDefault with default headers values
func NewHintedHandoffMetricsNotStoredHintsByAddrGetDefault(code int) *HintedHandoffMetricsNotStoredHintsByAddrGetDefault {
	return &HintedHandoffMetricsNotStoredHintsByAddrGetDefault{
		_statusCode: code,
	}
}

/*
HintedHandoffMetricsNotStoredHintsByAddrGetDefault handles this case with default header values.

internal server error
*/
type HintedHandoffMetricsNotStoredHintsByAddrGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the hinted handoff metrics not stored hints by addr get default response
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetDefault) Code() int {
	return o._statusCode
}

func (o *HintedHandoffMetricsNotStoredHintsByAddrGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *HintedHandoffMetricsNotStoredHintsByAddrGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *HintedHandoffMetricsNotStoredHintsByAddrGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
