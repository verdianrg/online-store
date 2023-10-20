// Code generated by go-swagger; DO NOT EDIT.

package history

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

	"online-store/gen/models"
)

// NewPostV1HistoriesParams creates a new PostV1HistoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1HistoriesParams() *PostV1HistoriesParams {
	return &PostV1HistoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1HistoriesParamsWithTimeout creates a new PostV1HistoriesParams object
// with the ability to set a timeout on a request.
func NewPostV1HistoriesParamsWithTimeout(timeout time.Duration) *PostV1HistoriesParams {
	return &PostV1HistoriesParams{
		timeout: timeout,
	}
}

// NewPostV1HistoriesParamsWithContext creates a new PostV1HistoriesParams object
// with the ability to set a context for a request.
func NewPostV1HistoriesParamsWithContext(ctx context.Context) *PostV1HistoriesParams {
	return &PostV1HistoriesParams{
		Context: ctx,
	}
}

// NewPostV1HistoriesParamsWithHTTPClient creates a new PostV1HistoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1HistoriesParamsWithHTTPClient(client *http.Client) *PostV1HistoriesParams {
	return &PostV1HistoriesParams{
		HTTPClient: client,
	}
}

/*
PostV1HistoriesParams contains all the parameters to send to the API endpoint

	for the post v1 histories operation.

	Typically these are written to a http.Request.
*/
type PostV1HistoriesParams struct {

	/* Data.

	   product request body
	*/
	Data *models.HistoryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 histories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1HistoriesParams) WithDefaults() *PostV1HistoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 histories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1HistoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 histories params
func (o *PostV1HistoriesParams) WithTimeout(timeout time.Duration) *PostV1HistoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 histories params
func (o *PostV1HistoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 histories params
func (o *PostV1HistoriesParams) WithContext(ctx context.Context) *PostV1HistoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 histories params
func (o *PostV1HistoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 histories params
func (o *PostV1HistoriesParams) WithHTTPClient(client *http.Client) *PostV1HistoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 histories params
func (o *PostV1HistoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the post v1 histories params
func (o *PostV1HistoriesParams) WithData(data *models.HistoryRequest) *PostV1HistoriesParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the post v1 histories params
func (o *PostV1HistoriesParams) SetData(data *models.HistoryRequest) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1HistoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}