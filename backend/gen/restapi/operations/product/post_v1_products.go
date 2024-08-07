// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"online-store/gen/models"
)

// PostV1ProductsHandlerFunc turns a function with the right signature into a post v1 products handler
type PostV1ProductsHandlerFunc func(PostV1ProductsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostV1ProductsHandlerFunc) Handle(params PostV1ProductsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostV1ProductsHandler interface for that can handle valid post v1 products params
type PostV1ProductsHandler interface {
	Handle(PostV1ProductsParams, *models.Principal) middleware.Responder
}

// NewPostV1Products creates a new http.Handler for the post v1 products operation
func NewPostV1Products(ctx *middleware.Context, handler PostV1ProductsHandler) *PostV1Products {
	return &PostV1Products{Context: ctx, Handler: handler}
}

/*
	PostV1Products swagger:route POST /v1/products product postV1Products

create new product

create new product API
*/
type PostV1Products struct {
	Context *middleware.Context
	Handler PostV1ProductsHandler
}

func (o *PostV1Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostV1ProductsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
