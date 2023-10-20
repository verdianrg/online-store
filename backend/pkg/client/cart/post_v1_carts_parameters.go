// Code generated by go-swagger; DO NOT EDIT.

package cart

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

// NewPostV1CartsParams creates a new PostV1CartsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1CartsParams() *PostV1CartsParams {
	return &PostV1CartsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1CartsParamsWithTimeout creates a new PostV1CartsParams object
// with the ability to set a timeout on a request.
func NewPostV1CartsParamsWithTimeout(timeout time.Duration) *PostV1CartsParams {
	return &PostV1CartsParams{
		timeout: timeout,
	}
}

// NewPostV1CartsParamsWithContext creates a new PostV1CartsParams object
// with the ability to set a context for a request.
func NewPostV1CartsParamsWithContext(ctx context.Context) *PostV1CartsParams {
	return &PostV1CartsParams{
		Context: ctx,
	}
}

// NewPostV1CartsParamsWithHTTPClient creates a new PostV1CartsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1CartsParamsWithHTTPClient(client *http.Client) *PostV1CartsParams {
	return &PostV1CartsParams{
		HTTPClient: client,
	}
}

/*
PostV1CartsParams contains all the parameters to send to the API endpoint

	for the post v1 carts operation.

	Typically these are written to a http.Request.
*/
type PostV1CartsParams struct {

	/* Data.

	   product request body
	*/
	Data []*models.CartRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 carts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1CartsParams) WithDefaults() *PostV1CartsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 carts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1CartsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 carts params
func (o *PostV1CartsParams) WithTimeout(timeout time.Duration) *PostV1CartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 carts params
func (o *PostV1CartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 carts params
func (o *PostV1CartsParams) WithContext(ctx context.Context) *PostV1CartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 carts params
func (o *PostV1CartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 carts params
func (o *PostV1CartsParams) WithHTTPClient(client *http.Client) *PostV1CartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 carts params
func (o *PostV1CartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the post v1 carts params
func (o *PostV1CartsParams) WithData(data []*models.CartRequest) *PostV1CartsParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the post v1 carts params
func (o *PostV1CartsParams) SetData(data []*models.CartRequest) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1CartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
