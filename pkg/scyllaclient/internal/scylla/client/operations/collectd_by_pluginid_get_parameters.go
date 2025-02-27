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

// NewCollectdByPluginidGetParams creates a new CollectdByPluginidGetParams object
// with the default values initialized.
func NewCollectdByPluginidGetParams() *CollectdByPluginidGetParams {
	var ()
	return &CollectdByPluginidGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCollectdByPluginidGetParamsWithTimeout creates a new CollectdByPluginidGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCollectdByPluginidGetParamsWithTimeout(timeout time.Duration) *CollectdByPluginidGetParams {
	var ()
	return &CollectdByPluginidGetParams{

		timeout: timeout,
	}
}

// NewCollectdByPluginidGetParamsWithContext creates a new CollectdByPluginidGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCollectdByPluginidGetParamsWithContext(ctx context.Context) *CollectdByPluginidGetParams {
	var ()
	return &CollectdByPluginidGetParams{

		Context: ctx,
	}
}

// NewCollectdByPluginidGetParamsWithHTTPClient creates a new CollectdByPluginidGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCollectdByPluginidGetParamsWithHTTPClient(client *http.Client) *CollectdByPluginidGetParams {
	var ()
	return &CollectdByPluginidGetParams{
		HTTPClient: client,
	}
}

/*
CollectdByPluginidGetParams contains all the parameters to send to the API endpoint
for the collectd by pluginid get operation typically these are written to a http.Request
*/
type CollectdByPluginidGetParams struct {

	/*Instance
	  The plugin instance

	*/
	Instance *string
	/*Pluginid
	  The plugin ID

	*/
	Pluginid string
	/*Type
	  The plugin type

	*/
	Type string
	/*TypeInstance
	  The plugin type instance

	*/
	TypeInstance *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) WithTimeout(timeout time.Duration) *CollectdByPluginidGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) WithContext(ctx context.Context) *CollectdByPluginidGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) WithHTTPClient(client *http.Client) *CollectdByPluginidGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstance adds the instance to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) WithInstance(instance *string) *CollectdByPluginidGetParams {
	o.SetInstance(instance)
	return o
}

// SetInstance adds the instance to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) SetInstance(instance *string) {
	o.Instance = instance
}

// WithPluginid adds the pluginid to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) WithPluginid(pluginid string) *CollectdByPluginidGetParams {
	o.SetPluginid(pluginid)
	return o
}

// SetPluginid adds the pluginid to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) SetPluginid(pluginid string) {
	o.Pluginid = pluginid
}

// WithType adds the typeVar to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) WithType(typeVar string) *CollectdByPluginidGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithTypeInstance adds the typeInstance to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) WithTypeInstance(typeInstance *string) *CollectdByPluginidGetParams {
	o.SetTypeInstance(typeInstance)
	return o
}

// SetTypeInstance adds the typeInstance to the collectd by pluginid get params
func (o *CollectdByPluginidGetParams) SetTypeInstance(typeInstance *string) {
	o.TypeInstance = typeInstance
}

// WriteToRequest writes these params to a swagger request
func (o *CollectdByPluginidGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Instance != nil {

		// query param instance
		var qrInstance string
		if o.Instance != nil {
			qrInstance = *o.Instance
		}
		qInstance := qrInstance
		if qInstance != "" {
			if err := r.SetQueryParam("instance", qInstance); err != nil {
				return err
			}
		}

	}

	// path param pluginid
	if err := r.SetPathParam("pluginid", o.Pluginid); err != nil {
		return err
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if o.TypeInstance != nil {

		// query param type_instance
		var qrTypeInstance string
		if o.TypeInstance != nil {
			qrTypeInstance = *o.TypeInstance
		}
		qTypeInstance := qrTypeInstance
		if qTypeInstance != "" {
			if err := r.SetQueryParam("type_instance", qTypeInstance); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
