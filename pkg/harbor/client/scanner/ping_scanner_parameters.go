// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// NewPingScannerParams creates a new PingScannerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPingScannerParams() *PingScannerParams {
	return &PingScannerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPingScannerParamsWithTimeout creates a new PingScannerParams object
// with the ability to set a timeout on a request.
func NewPingScannerParamsWithTimeout(timeout time.Duration) *PingScannerParams {
	return &PingScannerParams{
		timeout: timeout,
	}
}

// NewPingScannerParamsWithContext creates a new PingScannerParams object
// with the ability to set a context for a request.
func NewPingScannerParamsWithContext(ctx context.Context) *PingScannerParams {
	return &PingScannerParams{
		Context: ctx,
	}
}

// NewPingScannerParamsWithHTTPClient creates a new PingScannerParams object
// with the ability to set a custom HTTPClient for a request.
func NewPingScannerParamsWithHTTPClient(client *http.Client) *PingScannerParams {
	return &PingScannerParams{
		HTTPClient: client,
	}
}

/*
PingScannerParams contains all the parameters to send to the API endpoint

	for the ping scanner operation.

	Typically these are written to a http.Request.
*/
type PingScannerParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Settings.

	   A scanner registration settings to be tested.
	*/
	Settings *models.ScannerRegistrationSettings `js:"settings"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the ping scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingScannerParams) WithDefaults() *PingScannerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ping scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingScannerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ping scanner params
func (o *PingScannerParams) WithTimeout(timeout time.Duration) *PingScannerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping scanner params
func (o *PingScannerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping scanner params
func (o *PingScannerParams) WithContext(ctx context.Context) *PingScannerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping scanner params
func (o *PingScannerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping scanner params
func (o *PingScannerParams) WithHTTPClient(client *http.Client) *PingScannerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping scanner params
func (o *PingScannerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the ping scanner params
func (o *PingScannerParams) WithXRequestID(xRequestID *string) *PingScannerParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the ping scanner params
func (o *PingScannerParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSettings adds the settings to the ping scanner params
func (o *PingScannerParams) WithSettings(settings *models.ScannerRegistrationSettings) *PingScannerParams {
	o.SetSettings(settings)
	return o
}

// SetSettings adds the settings to the ping scanner params
func (o *PingScannerParams) SetSettings(settings *models.ScannerRegistrationSettings) {
	o.Settings = settings
}

// WriteToRequest writes these params to a swagger request
func (o *PingScannerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Settings != nil {
		if err := r.SetBodyParam(o.Settings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
