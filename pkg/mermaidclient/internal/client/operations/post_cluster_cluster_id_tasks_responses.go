// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/mermaidclient/internal/models"
)

// PostClusterClusterIDTasksReader is a Reader for the PostClusterClusterIDTasks structure.
type PostClusterClusterIDTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClusterClusterIDTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostClusterClusterIDTasksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostClusterClusterIDTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostClusterClusterIDTasksCreated creates a PostClusterClusterIDTasksCreated with default headers values
func NewPostClusterClusterIDTasksCreated() *PostClusterClusterIDTasksCreated {
	return &PostClusterClusterIDTasksCreated{}
}

/*
PostClusterClusterIDTasksCreated handles this case with default header values.

Task added
*/
type PostClusterClusterIDTasksCreated struct {
	Location string
}

func (o *PostClusterClusterIDTasksCreated) Error() string {
	return fmt.Sprintf("[POST /cluster/{cluster_id}/tasks][%d] postClusterClusterIdTasksCreated ", 201)
}

func (o *PostClusterClusterIDTasksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewPostClusterClusterIDTasksDefault creates a PostClusterClusterIDTasksDefault with default headers values
func NewPostClusterClusterIDTasksDefault(code int) *PostClusterClusterIDTasksDefault {
	return &PostClusterClusterIDTasksDefault{
		_statusCode: code,
	}
}

/*
PostClusterClusterIDTasksDefault handles this case with default header values.

Unexpected error
*/
type PostClusterClusterIDTasksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the post cluster cluster ID tasks default response
func (o *PostClusterClusterIDTasksDefault) Code() int {
	return o._statusCode
}

func (o *PostClusterClusterIDTasksDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/{cluster_id}/tasks][%d] PostClusterClusterIDTasks default  %+v", o._statusCode, o.Payload)
}

func (o *PostClusterClusterIDTasksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostClusterClusterIDTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
