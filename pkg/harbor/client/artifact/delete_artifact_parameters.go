// Code generated by go-swagger; DO NOT EDIT.

package artifact

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

// NewDeleteArtifactParams creates a new DeleteArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteArtifactParams() *DeleteArtifactParams {
	return &DeleteArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteArtifactParamsWithTimeout creates a new DeleteArtifactParams object
// with the ability to set a timeout on a request.
func NewDeleteArtifactParamsWithTimeout(timeout time.Duration) *DeleteArtifactParams {
	return &DeleteArtifactParams{
		timeout: timeout,
	}
}

// NewDeleteArtifactParamsWithContext creates a new DeleteArtifactParams object
// with the ability to set a context for a request.
func NewDeleteArtifactParamsWithContext(ctx context.Context) *DeleteArtifactParams {
	return &DeleteArtifactParams{
		Context: ctx,
	}
}

// NewDeleteArtifactParamsWithHTTPClient creates a new DeleteArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteArtifactParamsWithHTTPClient(client *http.Client) *DeleteArtifactParams {
	return &DeleteArtifactParams{
		HTTPClient: client,
	}
}

/*
DeleteArtifactParams contains all the parameters to send to the API endpoint

	for the delete artifact operation.

	Typically these are written to a http.Request.
*/
type DeleteArtifactParams struct {

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

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the delete artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteArtifactParams) WithDefaults() *DeleteArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete artifact params
func (o *DeleteArtifactParams) WithTimeout(timeout time.Duration) *DeleteArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete artifact params
func (o *DeleteArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete artifact params
func (o *DeleteArtifactParams) WithContext(ctx context.Context) *DeleteArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete artifact params
func (o *DeleteArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete artifact params
func (o *DeleteArtifactParams) WithHTTPClient(client *http.Client) *DeleteArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete artifact params
func (o *DeleteArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete artifact params
func (o *DeleteArtifactParams) WithXRequestID(xRequestID *string) *DeleteArtifactParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete artifact params
func (o *DeleteArtifactParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the delete artifact params
func (o *DeleteArtifactParams) WithProjectName(projectName string) *DeleteArtifactParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the delete artifact params
func (o *DeleteArtifactParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the delete artifact params
func (o *DeleteArtifactParams) WithReference(reference string) *DeleteArtifactParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the delete artifact params
func (o *DeleteArtifactParams) SetReference(reference string) {
	o.Reference = reference
}

// WithRepositoryName adds the repositoryName to the delete artifact params
func (o *DeleteArtifactParams) WithRepositoryName(repositoryName string) *DeleteArtifactParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the delete artifact params
func (o *DeleteArtifactParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
