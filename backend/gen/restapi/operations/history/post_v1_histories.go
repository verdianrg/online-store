// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"online-store/gen/models"
)

// PostV1HistoriesHandlerFunc turns a function with the right signature into a post v1 histories handler
type PostV1HistoriesHandlerFunc func(PostV1HistoriesParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostV1HistoriesHandlerFunc) Handle(params PostV1HistoriesParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostV1HistoriesHandler interface for that can handle valid post v1 histories params
type PostV1HistoriesHandler interface {
	Handle(PostV1HistoriesParams, *models.Principal) middleware.Responder
}

// NewPostV1Histories creates a new http.Handler for the post v1 histories operation
func NewPostV1Histories(ctx *middleware.Context, handler PostV1HistoriesHandler) *PostV1Histories {
	return &PostV1Histories{Context: ctx, Handler: handler}
}

/*
	PostV1Histories swagger:route POST /v1/histories history postV1Histories

create new history

create new history API
*/
type PostV1Histories struct {
	Context *middleware.Context
	Handler PostV1HistoriesHandler
}

func (o *PostV1Histories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostV1HistoriesParams()
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