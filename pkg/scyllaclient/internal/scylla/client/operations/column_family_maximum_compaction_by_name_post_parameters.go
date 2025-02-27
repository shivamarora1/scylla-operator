// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewColumnFamilyMaximumCompactionByNamePostParams creates a new ColumnFamilyMaximumCompactionByNamePostParams object
// with the default values initialized.
func NewColumnFamilyMaximumCompactionByNamePostParams() *ColumnFamilyMaximumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNamePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMaximumCompactionByNamePostParamsWithTimeout creates a new ColumnFamilyMaximumCompactionByNamePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMaximumCompactionByNamePostParamsWithTimeout(timeout time.Duration) *ColumnFamilyMaximumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNamePostParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMaximumCompactionByNamePostParamsWithContext creates a new ColumnFamilyMaximumCompactionByNamePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMaximumCompactionByNamePostParamsWithContext(ctx context.Context) *ColumnFamilyMaximumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNamePostParams{

		Context: ctx,
	}
}

// NewColumnFamilyMaximumCompactionByNamePostParamsWithHTTPClient creates a new ColumnFamilyMaximumCompactionByNamePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMaximumCompactionByNamePostParamsWithHTTPClient(client *http.Client) *ColumnFamilyMaximumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNamePostParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMaximumCompactionByNamePostParams contains all the parameters to send to the API endpoint
for the column family maximum compaction by name post operation typically these are written to a http.Request
*/
type ColumnFamilyMaximumCompactionByNamePostParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string
	/*Value
	  The maximum number of sstables in queue before compaction kicks off

	*/
	Value int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) WithTimeout(timeout time.Duration) *ColumnFamilyMaximumCompactionByNamePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) WithContext(ctx context.Context) *ColumnFamilyMaximumCompactionByNamePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) WithHTTPClient(client *http.Client) *ColumnFamilyMaximumCompactionByNamePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) WithName(name string) *ColumnFamilyMaximumCompactionByNamePostParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) SetName(name string) {
	o.Name = name
}

// WithValue adds the value to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) WithValue(value int32) *ColumnFamilyMaximumCompactionByNamePostParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the column family maximum compaction by name post params
func (o *ColumnFamilyMaximumCompactionByNamePostParams) SetValue(value int32) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMaximumCompactionByNamePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param value
	qrValue := o.Value
	qValue := swag.FormatInt32(qrValue)
	if qValue != "" {
		if err := r.SetQueryParam("value", qValue); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
