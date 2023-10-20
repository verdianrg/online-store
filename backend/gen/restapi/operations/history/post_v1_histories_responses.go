// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"online-store/gen/models"
)

// PostV1HistoriesCreatedCode is the HTTP code returned for type PostV1HistoriesCreated
const PostV1HistoriesCreatedCode int = 201

/*
PostV1HistoriesCreated success create new cart

swagger:response postV1HistoriesCreated
*/
type PostV1HistoriesCreated struct {

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewPostV1HistoriesCreated creates PostV1HistoriesCreated with default headers values
func NewPostV1HistoriesCreated() *PostV1HistoriesCreated {

	return &PostV1HistoriesCreated{}
}

// WithPayload adds the payload to the post v1 histories created response
func (o *PostV1HistoriesCreated) WithPayload(payload *models.BaseResponse) *PostV1HistoriesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 histories created response
func (o *PostV1HistoriesCreated) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1HistoriesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostV1HistoriesUnauthorizedCode is the HTTP code returned for type PostV1HistoriesUnauthorized
const PostV1HistoriesUnauthorizedCode int = 401

/*
PostV1HistoriesUnauthorized unauthorized

swagger:response postV1HistoriesUnauthorized
*/
type PostV1HistoriesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostV1HistoriesUnauthorized creates PostV1HistoriesUnauthorized with default headers values
func NewPostV1HistoriesUnauthorized() *PostV1HistoriesUnauthorized {

	return &PostV1HistoriesUnauthorized{}
}

// WithPayload adds the payload to the post v1 histories unauthorized response
func (o *PostV1HistoriesUnauthorized) WithPayload(payload interface{}) *PostV1HistoriesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 histories unauthorized response
func (o *PostV1HistoriesUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1HistoriesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
PostV1HistoriesDefault error

swagger:response postV1HistoriesDefault
*/
type PostV1HistoriesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewPostV1HistoriesDefault creates PostV1HistoriesDefault with default headers values
func NewPostV1HistoriesDefault(code int) *PostV1HistoriesDefault {
	if code <= 0 {
		code = 500
	}

	return &PostV1HistoriesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post v1 histories default response
func (o *PostV1HistoriesDefault) WithStatusCode(code int) *PostV1HistoriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post v1 histories default response
func (o *PostV1HistoriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post v1 histories default response
func (o *PostV1HistoriesDefault) WithPayload(payload *models.BaseResponse) *PostV1HistoriesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 histories default response
func (o *PostV1HistoriesDefault) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1HistoriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
