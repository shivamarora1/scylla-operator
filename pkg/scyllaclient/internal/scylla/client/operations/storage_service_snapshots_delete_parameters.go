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

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageServiceSnapshotsDeleteParams creates a new StorageServiceSnapshotsDeleteParams object
// with the default values initialized.
func NewStorageServiceSnapshotsDeleteParams() *StorageServiceSnapshotsDeleteParams {
	var ()
	return &StorageServiceSnapshotsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceSnapshotsDeleteParamsWithTimeout creates a new StorageServiceSnapshotsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceSnapshotsDeleteParamsWithTimeout(timeout time.Duration) *StorageServiceSnapshotsDeleteParams {
	var ()
	return &StorageServiceSnapshotsDeleteParams{

		timeout: timeout,
	}
}

// NewStorageServiceSnapshotsDeleteParamsWithContext creates a new StorageServiceSnapshotsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceSnapshotsDeleteParamsWithContext(ctx context.Context) *StorageServiceSnapshotsDeleteParams {
	var ()
	return &StorageServiceSnapshotsDeleteParams{

		Context: ctx,
	}
}

// NewStorageServiceSnapshotsDeleteParamsWithHTTPClient creates a new StorageServiceSnapshotsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceSnapshotsDeleteParamsWithHTTPClient(client *http.Client) *StorageServiceSnapshotsDeleteParams {
	var ()
	return &StorageServiceSnapshotsDeleteParams{
		HTTPClient: client,
	}
}

/*
StorageServiceSnapshotsDeleteParams contains all the parameters to send to the API endpoint
for the storage service snapshots delete operation typically these are written to a http.Request
*/
type StorageServiceSnapshotsDeleteParams struct {

	/*Kn
	  Comma seperated keyspaces name to snapshot

	*/
	Kn *string
	/*Tag
	  the tag given to the snapshot

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) WithTimeout(timeout time.Duration) *StorageServiceSnapshotsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) WithContext(ctx context.Context) *StorageServiceSnapshotsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) WithHTTPClient(client *http.Client) *StorageServiceSnapshotsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKn adds the kn to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) WithKn(kn *string) *StorageServiceSnapshotsDeleteParams {
	o.SetKn(kn)
	return o
}

// SetKn adds the kn to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) SetKn(kn *string) {
	o.Kn = kn
}

// WithTag adds the tag to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) WithTag(tag *string) *StorageServiceSnapshotsDeleteParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the storage service snapshots delete params
func (o *StorageServiceSnapshotsDeleteParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceSnapshotsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Kn != nil {

		// query param kn
		var qrKn string
		if o.Kn != nil {
			qrKn = *o.Kn
		}
		qKn := qrKn
		if qKn != "" {
			if err := r.SetQueryParam("kn", qKn); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
