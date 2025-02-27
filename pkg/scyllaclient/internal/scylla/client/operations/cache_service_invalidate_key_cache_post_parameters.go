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

// NewCacheServiceInvalidateKeyCachePostParams creates a new CacheServiceInvalidateKeyCachePostParams object
// with the default values initialized.
func NewCacheServiceInvalidateKeyCachePostParams() *CacheServiceInvalidateKeyCachePostParams {

	return &CacheServiceInvalidateKeyCachePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceInvalidateKeyCachePostParamsWithTimeout creates a new CacheServiceInvalidateKeyCachePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceInvalidateKeyCachePostParamsWithTimeout(timeout time.Duration) *CacheServiceInvalidateKeyCachePostParams {

	return &CacheServiceInvalidateKeyCachePostParams{

		timeout: timeout,
	}
}

// NewCacheServiceInvalidateKeyCachePostParamsWithContext creates a new CacheServiceInvalidateKeyCachePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceInvalidateKeyCachePostParamsWithContext(ctx context.Context) *CacheServiceInvalidateKeyCachePostParams {

	return &CacheServiceInvalidateKeyCachePostParams{

		Context: ctx,
	}
}

// NewCacheServiceInvalidateKeyCachePostParamsWithHTTPClient creates a new CacheServiceInvalidateKeyCachePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceInvalidateKeyCachePostParamsWithHTTPClient(client *http.Client) *CacheServiceInvalidateKeyCachePostParams {

	return &CacheServiceInvalidateKeyCachePostParams{
		HTTPClient: client,
	}
}

/*
CacheServiceInvalidateKeyCachePostParams contains all the parameters to send to the API endpoint
for the cache service invalidate key cache post operation typically these are written to a http.Request
*/
type CacheServiceInvalidateKeyCachePostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service invalidate key cache post params
func (o *CacheServiceInvalidateKeyCachePostParams) WithTimeout(timeout time.Duration) *CacheServiceInvalidateKeyCachePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service invalidate key cache post params
func (o *CacheServiceInvalidateKeyCachePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service invalidate key cache post params
func (o *CacheServiceInvalidateKeyCachePostParams) WithContext(ctx context.Context) *CacheServiceInvalidateKeyCachePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service invalidate key cache post params
func (o *CacheServiceInvalidateKeyCachePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service invalidate key cache post params
func (o *CacheServiceInvalidateKeyCachePostParams) WithHTTPClient(client *http.Client) *CacheServiceInvalidateKeyCachePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service invalidate key cache post params
func (o *CacheServiceInvalidateKeyCachePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceInvalidateKeyCachePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
