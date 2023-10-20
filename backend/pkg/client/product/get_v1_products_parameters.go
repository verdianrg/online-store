// Code generated by go-swagger; DO NOT EDIT.

package product

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

// NewGetV1ProductsParams creates a new GetV1ProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ProductsParams() *GetV1ProductsParams {
	return &GetV1ProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ProductsParamsWithTimeout creates a new GetV1ProductsParams object
// with the ability to set a timeout on a request.
func NewGetV1ProductsParamsWithTimeout(timeout time.Duration) *GetV1ProductsParams {
	return &GetV1ProductsParams{
		timeout: timeout,
	}
}

// NewGetV1ProductsParamsWithContext creates a new GetV1ProductsParams object
// with the ability to set a context for a request.
func NewGetV1ProductsParamsWithContext(ctx context.Context) *GetV1ProductsParams {
	return &GetV1ProductsParams{
		Context: ctx,
	}
}

// NewGetV1ProductsParamsWithHTTPClient creates a new GetV1ProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ProductsParamsWithHTTPClient(client *http.Client) *GetV1ProductsParams {
	return &GetV1ProductsParams{
		HTTPClient: client,
	}
}

/*
GetV1ProductsParams contains all the parameters to send to the API endpoint

	for the get v1 products operation.

	Typically these are written to a http.Request.
*/
type GetV1ProductsParams struct {

	/* Keyword.

	   search keyword
	*/
	Keyword *string

	/* Limit.

	   The numbers of items to return.

	   Format: int32
	*/
	Limit *int32

	/* Page.

	   the number of page

	   Format: int32
	*/
	Page *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ProductsParams) WithDefaults() *GetV1ProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 products params
func (o *GetV1ProductsParams) WithTimeout(timeout time.Duration) *GetV1ProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 products params
func (o *GetV1ProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 products params
func (o *GetV1ProductsParams) WithContext(ctx context.Context) *GetV1ProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 products params
func (o *GetV1ProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 products params
func (o *GetV1ProductsParams) WithHTTPClient(client *http.Client) *GetV1ProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 products params
func (o *GetV1ProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyword adds the keyword to the get v1 products params
func (o *GetV1ProductsParams) WithKeyword(keyword *string) *GetV1ProductsParams {
	o.SetKeyword(keyword)
	return o
}

// SetKeyword adds the keyword to the get v1 products params
func (o *GetV1ProductsParams) SetKeyword(keyword *string) {
	o.Keyword = keyword
}

// WithLimit adds the limit to the get v1 products params
func (o *GetV1ProductsParams) WithLimit(limit *int32) *GetV1ProductsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get v1 products params
func (o *GetV1ProductsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get v1 products params
func (o *GetV1ProductsParams) WithPage(page *int32) *GetV1ProductsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 products params
func (o *GetV1ProductsParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Keyword != nil {

		// query param keyword
		var qrKeyword string

		if o.Keyword != nil {
			qrKeyword = *o.Keyword
		}
		qKeyword := qrKeyword
		if qKeyword != "" {

			if err := r.SetQueryParam("keyword", qKeyword); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
