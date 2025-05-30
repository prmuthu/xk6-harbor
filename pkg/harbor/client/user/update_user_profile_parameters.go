// Code generated by go-swagger; DO NOT EDIT.

package user

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
	"github.com/go-openapi/swag"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// NewUpdateUserProfileParams creates a new UpdateUserProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserProfileParams() *UpdateUserProfileParams {
	return &UpdateUserProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserProfileParamsWithTimeout creates a new UpdateUserProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateUserProfileParamsWithTimeout(timeout time.Duration) *UpdateUserProfileParams {
	return &UpdateUserProfileParams{
		timeout: timeout,
	}
}

// NewUpdateUserProfileParamsWithContext creates a new UpdateUserProfileParams object
// with the ability to set a context for a request.
func NewUpdateUserProfileParamsWithContext(ctx context.Context) *UpdateUserProfileParams {
	return &UpdateUserProfileParams{
		Context: ctx,
	}
}

// NewUpdateUserProfileParamsWithHTTPClient creates a new UpdateUserProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserProfileParamsWithHTTPClient(client *http.Client) *UpdateUserProfileParams {
	return &UpdateUserProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateUserProfileParams contains all the parameters to send to the API endpoint

	for the update user profile operation.

	Typically these are written to a http.Request.
*/
type UpdateUserProfileParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Profile.

	   Only email, realname and comment can be modified.
	*/
	Profile *models.UserProfile `js:"profile"`

	/* UserID.

	   Registered user ID

	   Format: int
	*/
	UserID int64 `js:"userID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the update user profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserProfileParams) WithDefaults() *UpdateUserProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user profile params
func (o *UpdateUserProfileParams) WithTimeout(timeout time.Duration) *UpdateUserProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user profile params
func (o *UpdateUserProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user profile params
func (o *UpdateUserProfileParams) WithContext(ctx context.Context) *UpdateUserProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user profile params
func (o *UpdateUserProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user profile params
func (o *UpdateUserProfileParams) WithHTTPClient(client *http.Client) *UpdateUserProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user profile params
func (o *UpdateUserProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update user profile params
func (o *UpdateUserProfileParams) WithXRequestID(xRequestID *string) *UpdateUserProfileParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update user profile params
func (o *UpdateUserProfileParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProfile adds the profile to the update user profile params
func (o *UpdateUserProfileParams) WithProfile(profile *models.UserProfile) *UpdateUserProfileParams {
	o.SetProfile(profile)
	return o
}

// SetProfile adds the profile to the update user profile params
func (o *UpdateUserProfileParams) SetProfile(profile *models.UserProfile) {
	o.Profile = profile
}

// WithUserID adds the userID to the update user profile params
func (o *UpdateUserProfileParams) WithUserID(userID int64) *UpdateUserProfileParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user profile params
func (o *UpdateUserProfileParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Profile != nil {
		if err := r.SetBodyParam(o.Profile); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
