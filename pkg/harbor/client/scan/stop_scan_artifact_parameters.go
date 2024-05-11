// Code generated by go-swagger; DO NOT EDIT.

package scan

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

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// NewStopScanArtifactParams creates a new StopScanArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopScanArtifactParams() *StopScanArtifactParams {
	return &StopScanArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopScanArtifactParamsWithTimeout creates a new StopScanArtifactParams object
// with the ability to set a timeout on a request.
func NewStopScanArtifactParamsWithTimeout(timeout time.Duration) *StopScanArtifactParams {
	return &StopScanArtifactParams{
		timeout: timeout,
	}
}

// NewStopScanArtifactParamsWithContext creates a new StopScanArtifactParams object
// with the ability to set a context for a request.
func NewStopScanArtifactParamsWithContext(ctx context.Context) *StopScanArtifactParams {
	return &StopScanArtifactParams{
		Context: ctx,
	}
}

// NewStopScanArtifactParamsWithHTTPClient creates a new StopScanArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopScanArtifactParamsWithHTTPClient(client *http.Client) *StopScanArtifactParams {
	return &StopScanArtifactParams{
		HTTPClient: client,
	}
}

/*
StopScanArtifactParams contains all the parameters to send to the API endpoint

	for the stop scan artifact operation.

	Typically these are written to a http.Request.
*/
type StopScanArtifactParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string `js:"projectName"`

	/* Reference.

	   The reference of the artifact, can be digest or tag
	*/
	Reference string `js:"reference"`

	/* RepositoryName.

	   The name of the repository. If it contains slash, encode it twice over with URL encoding. e.g. a/b -> a%2Fb -> a%252Fb
	*/
	RepositoryName string `js:"repositoryName"`

	/* ScanType.

	   The scan type: Vulnerabilities, SBOM
	*/
	ScanType *models.ScanType `js:"scanType"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the stop scan artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopScanArtifactParams) WithDefaults() *StopScanArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop scan artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopScanArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop scan artifact params
func (o *StopScanArtifactParams) WithTimeout(timeout time.Duration) *StopScanArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop scan artifact params
func (o *StopScanArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop scan artifact params
func (o *StopScanArtifactParams) WithContext(ctx context.Context) *StopScanArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop scan artifact params
func (o *StopScanArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop scan artifact params
func (o *StopScanArtifactParams) WithHTTPClient(client *http.Client) *StopScanArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop scan artifact params
func (o *StopScanArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the stop scan artifact params
func (o *StopScanArtifactParams) WithXRequestID(xRequestID *string) *StopScanArtifactParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the stop scan artifact params
func (o *StopScanArtifactParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the stop scan artifact params
func (o *StopScanArtifactParams) WithProjectName(projectName string) *StopScanArtifactParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the stop scan artifact params
func (o *StopScanArtifactParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the stop scan artifact params
func (o *StopScanArtifactParams) WithReference(reference string) *StopScanArtifactParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the stop scan artifact params
func (o *StopScanArtifactParams) SetReference(reference string) {
	o.Reference = reference
}

// WithRepositoryName adds the repositoryName to the stop scan artifact params
func (o *StopScanArtifactParams) WithRepositoryName(repositoryName string) *StopScanArtifactParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the stop scan artifact params
func (o *StopScanArtifactParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WithScanType adds the scanType to the stop scan artifact params
func (o *StopScanArtifactParams) WithScanType(scanType *models.ScanType) *StopScanArtifactParams {
	o.SetScanType(scanType)
	return o
}

// SetScanType adds the scanType to the stop scan artifact params
func (o *StopScanArtifactParams) SetScanType(scanType *models.ScanType) {
	o.ScanType = scanType
}

// WriteToRequest writes these params to a swagger request
func (o *StopScanArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}
	if o.ScanType != nil {
		if err := r.SetBodyParam(o.ScanType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
