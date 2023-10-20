// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"online-store/gen/models"
)

// PostV1HistoriesReader is a Reader for the PostV1Histories structure.
type PostV1HistoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1HistoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1HistoriesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostV1HistoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostV1HistoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostV1HistoriesCreated creates a PostV1HistoriesCreated with default headers values
func NewPostV1HistoriesCreated() *PostV1HistoriesCreated {
	return &PostV1HistoriesCreated{}
}

/*
PostV1HistoriesCreated describes a response with status code 201, with default header values.

success create new cart
*/
type PostV1HistoriesCreated struct {
	Payload *models.BaseResponse
}

// IsSuccess returns true when this post v1 histories created response has a 2xx status code
func (o *PostV1HistoriesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 histories created response has a 3xx status code
func (o *PostV1HistoriesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 histories created response has a 4xx status code
func (o *PostV1HistoriesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 histories created response has a 5xx status code
func (o *PostV1HistoriesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 histories created response a status code equal to that given
func (o *PostV1HistoriesCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 histories created response
func (o *PostV1HistoriesCreated) Code() int {
	return 201
}

func (o *PostV1HistoriesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/histories][%d] postV1HistoriesCreated  %+v", 201, o.Payload)
}

func (o *PostV1HistoriesCreated) String() string {
	return fmt.Sprintf("[POST /v1/histories][%d] postV1HistoriesCreated  %+v", 201, o.Payload)
}

func (o *PostV1HistoriesCreated) GetPayload() *models.BaseResponse {
	return o.Payload
}

func (o *PostV1HistoriesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BaseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1HistoriesUnauthorized creates a PostV1HistoriesUnauthorized with default headers values
func NewPostV1HistoriesUnauthorized() *PostV1HistoriesUnauthorized {
	return &PostV1HistoriesUnauthorized{}
}

/*
PostV1HistoriesUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type PostV1HistoriesUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 histories unauthorized response has a 2xx status code
func (o *PostV1HistoriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 histories unauthorized response has a 3xx status code
func (o *PostV1HistoriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 histories unauthorized response has a 4xx status code
func (o *PostV1HistoriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 histories unauthorized response has a 5xx status code
func (o *PostV1HistoriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 histories unauthorized response a status code equal to that given
func (o *PostV1HistoriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post v1 histories unauthorized response
func (o *PostV1HistoriesUnauthorized) Code() int {
	return 401
}

func (o *PostV1HistoriesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/histories][%d] postV1HistoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1HistoriesUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/histories][%d] postV1HistoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1HistoriesUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1HistoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1HistoriesDefault creates a PostV1HistoriesDefault with default headers values
func NewPostV1HistoriesDefault(code int) *PostV1HistoriesDefault {
	return &PostV1HistoriesDefault{
		_statusCode: code,
	}
}

/*
PostV1HistoriesDefault describes a response with status code -1, with default header values.

error
*/
type PostV1HistoriesDefault struct {
	_statusCode int

	Payload *models.BaseResponse
}

// IsSuccess returns true when this post v1 histories default response has a 2xx status code
func (o *PostV1HistoriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post v1 histories default response has a 3xx status code
func (o *PostV1HistoriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post v1 histories default response has a 4xx status code
func (o *PostV1HistoriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post v1 histories default response has a 5xx status code
func (o *PostV1HistoriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post v1 histories default response a status code equal to that given
func (o *PostV1HistoriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post v1 histories default response
func (o *PostV1HistoriesDefault) Code() int {
	return o._statusCode
}

func (o *PostV1HistoriesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/histories][%d] PostV1Histories default  %+v", o._statusCode, o.Payload)
}

func (o *PostV1HistoriesDefault) String() string {
	return fmt.Sprintf("[POST /v1/histories][%d] PostV1Histories default  %+v", o._statusCode, o.Payload)
}

func (o *PostV1HistoriesDefault) GetPayload() *models.BaseResponse {
	return o.Payload
}

func (o *PostV1HistoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BaseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
