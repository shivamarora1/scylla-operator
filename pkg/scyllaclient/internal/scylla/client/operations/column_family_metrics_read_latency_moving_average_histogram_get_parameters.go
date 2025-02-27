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

// NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams creates a new ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams() *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParamsWithTimeout creates a new ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParamsWithContext creates a new ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams {

	return &ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams contains all the parameters to send to the API endpoint
for the column family metrics read latency moving average histogram get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read latency moving average histogram get params
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read latency moving average histogram get params
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read latency moving average histogram get params
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read latency moving average histogram get params
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read latency moving average histogram get params
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read latency moving average histogram get params
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
