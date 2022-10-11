// Code generated by go-swagger; DO NOT EDIT.

package delete_user_by_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/obarbier/custom-app/core/pkg/models"
)

// DeleteUserUserIDOKCode is the HTTP code returned for type DeleteUserUserIDOK
const DeleteUserUserIDOKCode int = 200

/*
DeleteUserUserIDOK OK

swagger:response deleteUserUserIdOK
*/
type DeleteUserUserIDOK struct {
}

// NewDeleteUserUserIDOK creates DeleteUserUserIDOK with default headers values
func NewDeleteUserUserIDOK() *DeleteUserUserIDOK {

	return &DeleteUserUserIDOK{}
}

// WriteResponse to the client
func (o *DeleteUserUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteUserUserIDBadRequestCode is the HTTP code returned for type DeleteUserUserIDBadRequest
const DeleteUserUserIDBadRequestCode int = 400

/*
DeleteUserUserIDBadRequest The specified user ID is invalid (e.g. not a number).

swagger:response deleteUserUserIdBadRequest
*/
type DeleteUserUserIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserUserIDBadRequest creates DeleteUserUserIDBadRequest with default headers values
func NewDeleteUserUserIDBadRequest() *DeleteUserUserIDBadRequest {

	return &DeleteUserUserIDBadRequest{}
}

// WithPayload adds the payload to the delete user user Id bad request response
func (o *DeleteUserUserIDBadRequest) WithPayload(payload *models.Error) *DeleteUserUserIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user user Id bad request response
func (o *DeleteUserUserIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserUserIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserUserIDUnauthorizedCode is the HTTP code returned for type DeleteUserUserIDUnauthorized
const DeleteUserUserIDUnauthorizedCode int = 401

/*
DeleteUserUserIDUnauthorized unauthorized

swagger:response deleteUserUserIdUnauthorized
*/
type DeleteUserUserIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserUserIDUnauthorized creates DeleteUserUserIDUnauthorized with default headers values
func NewDeleteUserUserIDUnauthorized() *DeleteUserUserIDUnauthorized {

	return &DeleteUserUserIDUnauthorized{}
}

// WithPayload adds the payload to the delete user user Id unauthorized response
func (o *DeleteUserUserIDUnauthorized) WithPayload(payload *models.Error) *DeleteUserUserIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user user Id unauthorized response
func (o *DeleteUserUserIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserUserIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserUserIDNotFoundCode is the HTTP code returned for type DeleteUserUserIDNotFound
const DeleteUserUserIDNotFoundCode int = 404

/*
DeleteUserUserIDNotFound A user with the specified ID was not found.

swagger:response deleteUserUserIdNotFound
*/
type DeleteUserUserIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserUserIDNotFound creates DeleteUserUserIDNotFound with default headers values
func NewDeleteUserUserIDNotFound() *DeleteUserUserIDNotFound {

	return &DeleteUserUserIDNotFound{}
}

// WithPayload adds the payload to the delete user user Id not found response
func (o *DeleteUserUserIDNotFound) WithPayload(payload *models.Error) *DeleteUserUserIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user user Id not found response
func (o *DeleteUserUserIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
DeleteUserUserIDDefault error

swagger:response deleteUserUserIdDefault
*/
type DeleteUserUserIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserUserIDDefault creates DeleteUserUserIDDefault with default headers values
func NewDeleteUserUserIDDefault(code int) *DeleteUserUserIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUserUserIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete user user ID default response
func (o *DeleteUserUserIDDefault) WithStatusCode(code int) *DeleteUserUserIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete user user ID default response
func (o *DeleteUserUserIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete user user ID default response
func (o *DeleteUserUserIDDefault) WithPayload(payload *models.Error) *DeleteUserUserIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user user ID default response
func (o *DeleteUserUserIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserUserIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
