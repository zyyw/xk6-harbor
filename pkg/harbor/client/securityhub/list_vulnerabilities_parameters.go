// Code generated by go-swagger; DO NOT EDIT.

package securityhub

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
)

// NewListVulnerabilitiesParams creates a new ListVulnerabilitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVulnerabilitiesParams() *ListVulnerabilitiesParams {
	return &ListVulnerabilitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVulnerabilitiesParamsWithTimeout creates a new ListVulnerabilitiesParams object
// with the ability to set a timeout on a request.
func NewListVulnerabilitiesParamsWithTimeout(timeout time.Duration) *ListVulnerabilitiesParams {
	return &ListVulnerabilitiesParams{
		timeout: timeout,
	}
}

// NewListVulnerabilitiesParamsWithContext creates a new ListVulnerabilitiesParams object
// with the ability to set a context for a request.
func NewListVulnerabilitiesParamsWithContext(ctx context.Context) *ListVulnerabilitiesParams {
	return &ListVulnerabilitiesParams{
		Context: ctx,
	}
}

// NewListVulnerabilitiesParamsWithHTTPClient creates a new ListVulnerabilitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVulnerabilitiesParamsWithHTTPClient(client *http.Client) *ListVulnerabilitiesParams {
	return &ListVulnerabilitiesParams{
		HTTPClient: client,
	}
}

/*
ListVulnerabilitiesParams contains all the parameters to send to the API endpoint

	for the list vulnerabilities operation.

	Typically these are written to a http.Request.
*/
type ListVulnerabilitiesParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64 `js:"page"`

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64 `js:"pageSize"`

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string `js:"q"`

	/* TuneCount.

	   Enable to ignore X-Total-Count when the total count > 1000, if the total count is less than 1000, the real total count is returned, else -1.
	*/
	TuneCount *bool `js:"tuneCount"`

	/* WithTag.

	   Specify whether the tag information is included inside vulnerability information
	*/
	WithTag *bool `js:"withTag"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the list vulnerabilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVulnerabilitiesParams) WithDefaults() *ListVulnerabilitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list vulnerabilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVulnerabilitiesParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)

		tuneCountDefault = bool(false)

		withTagDefault = bool(false)
	)

	val := ListVulnerabilitiesParams{
		Page:      &pageDefault,
		PageSize:  &pageSizeDefault,
		TuneCount: &tuneCountDefault,
		WithTag:   &withTagDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithTimeout(timeout time.Duration) *ListVulnerabilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithContext(ctx context.Context) *ListVulnerabilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithHTTPClient(client *http.Client) *ListVulnerabilitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithXRequestID(xRequestID *string) *ListVulnerabilitiesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithPage(page *int64) *ListVulnerabilitiesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithPageSize(pageSize *int64) *ListVulnerabilitiesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithQ adds the q to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithQ(q *string) *ListVulnerabilitiesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetQ(q *string) {
	o.Q = q
}

// WithTuneCount adds the tuneCount to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithTuneCount(tuneCount *bool) *ListVulnerabilitiesParams {
	o.SetTuneCount(tuneCount)
	return o
}

// SetTuneCount adds the tuneCount to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetTuneCount(tuneCount *bool) {
	o.TuneCount = tuneCount
}

// WithWithTag adds the withTag to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) WithWithTag(withTag *bool) *ListVulnerabilitiesParams {
	o.SetWithTag(withTag)
	return o
}

// SetWithTag adds the withTag to the list vulnerabilities params
func (o *ListVulnerabilitiesParams) SetWithTag(withTag *bool) {
	o.WithTag = withTag
}

// WriteToRequest writes these params to a swagger request
func (o *ListVulnerabilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.TuneCount != nil {

		// query param tune_count
		var qrTuneCount bool

		if o.TuneCount != nil {
			qrTuneCount = *o.TuneCount
		}
		qTuneCount := swag.FormatBool(qrTuneCount)
		if qTuneCount != "" {

			if err := r.SetQueryParam("tune_count", qTuneCount); err != nil {
				return err
			}
		}
	}

	if o.WithTag != nil {

		// query param with_tag
		var qrWithTag bool

		if o.WithTag != nil {
			qrWithTag = *o.WithTag
		}
		qWithTag := swag.FormatBool(qrWithTag)
		if qWithTag != "" {

			if err := r.SetQueryParam("with_tag", qWithTag); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
