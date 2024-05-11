// Code generated by go-swagger; DO NOT EDIT.

package schedule

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
)

// NewGetSchedulePausedParams creates a new GetSchedulePausedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSchedulePausedParams() *GetSchedulePausedParams {
	return &GetSchedulePausedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchedulePausedParamsWithTimeout creates a new GetSchedulePausedParams object
// with the ability to set a timeout on a request.
func NewGetSchedulePausedParamsWithTimeout(timeout time.Duration) *GetSchedulePausedParams {
	return &GetSchedulePausedParams{
		timeout: timeout,
	}
}

// NewGetSchedulePausedParamsWithContext creates a new GetSchedulePausedParams object
// with the ability to set a context for a request.
func NewGetSchedulePausedParamsWithContext(ctx context.Context) *GetSchedulePausedParams {
	return &GetSchedulePausedParams{
		Context: ctx,
	}
}

// NewGetSchedulePausedParamsWithHTTPClient creates a new GetSchedulePausedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSchedulePausedParamsWithHTTPClient(client *http.Client) *GetSchedulePausedParams {
	return &GetSchedulePausedParams{
		HTTPClient: client,
	}
}

/*
GetSchedulePausedParams contains all the parameters to send to the API endpoint

	for the get schedule paused operation.

	Typically these are written to a http.Request.
*/
type GetSchedulePausedParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* JobType.

	   The type of the job. 'all' stands for all job types, current only support query with all
	*/
	JobType string `js:"jobType"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get schedule paused params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchedulePausedParams) WithDefaults() *GetSchedulePausedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule paused params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchedulePausedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule paused params
func (o *GetSchedulePausedParams) WithTimeout(timeout time.Duration) *GetSchedulePausedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule paused params
func (o *GetSchedulePausedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule paused params
func (o *GetSchedulePausedParams) WithContext(ctx context.Context) *GetSchedulePausedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule paused params
func (o *GetSchedulePausedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule paused params
func (o *GetSchedulePausedParams) WithHTTPClient(client *http.Client) *GetSchedulePausedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule paused params
func (o *GetSchedulePausedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get schedule paused params
func (o *GetSchedulePausedParams) WithXRequestID(xRequestID *string) *GetSchedulePausedParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get schedule paused params
func (o *GetSchedulePausedParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithJobType adds the jobType to the get schedule paused params
func (o *GetSchedulePausedParams) WithJobType(jobType string) *GetSchedulePausedParams {
	o.SetJobType(jobType)
	return o
}

// SetJobType adds the jobType to the get schedule paused params
func (o *GetSchedulePausedParams) SetJobType(jobType string) {
	o.JobType = jobType
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchedulePausedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param job_type
	if err := r.SetPathParam("job_type", o.JobType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
