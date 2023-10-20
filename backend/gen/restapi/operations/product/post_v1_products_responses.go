// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"online-store/gen/models"
)

// PostV1ProductsCreatedCode is the HTTP code returned for type PostV1ProductsCreated
const PostV1ProductsCreatedCode int = 201

/*
PostV1ProductsCreated success create new product

swagger:response postV1ProductsCreated
*/
type PostV1ProductsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewPostV1ProductsCreated creates PostV1ProductsCreated with default headers values
func NewPostV1ProductsCreated() *PostV1ProductsCreated {

	return &PostV1ProductsCreated{}
}

// WithPayload adds the payload to the post v1 products created response
func (o *PostV1ProductsCreated) WithPayload(payload *models.BaseResponse) *PostV1ProductsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 products created response
func (o *PostV1ProductsCreated) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1ProductsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostV1ProductsUnauthorizedCode is the HTTP code returned for type PostV1ProductsUnauthorized
const PostV1ProductsUnauthorizedCode int = 401

/*
PostV1ProductsUnauthorized unauthorized

swagger:response postV1ProductsUnauthorized
*/
type PostV1ProductsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostV1ProductsUnauthorized creates PostV1ProductsUnauthorized with default headers values
func NewPostV1ProductsUnauthorized() *PostV1ProductsUnauthorized {

	return &PostV1ProductsUnauthorized{}
}

// WithPayload adds the payload to the post v1 products unauthorized response
func (o *PostV1ProductsUnauthorized) WithPayload(payload interface{}) *PostV1ProductsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 products unauthorized response
func (o *PostV1ProductsUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1ProductsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
PostV1ProductsDefault error

swagger:response postV1ProductsDefault
*/
type PostV1ProductsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewPostV1ProductsDefault creates PostV1ProductsDefault with default headers values
func NewPostV1ProductsDefault(code int) *PostV1ProductsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostV1ProductsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post v1 products default response
func (o *PostV1ProductsDefault) WithStatusCode(code int) *PostV1ProductsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post v1 products default response
func (o *PostV1ProductsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post v1 products default response
func (o *PostV1ProductsDefault) WithPayload(payload *models.BaseResponse) *PostV1ProductsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 products default response
func (o *PostV1ProductsDefault) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1ProductsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
