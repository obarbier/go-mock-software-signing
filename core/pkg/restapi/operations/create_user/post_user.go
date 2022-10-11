// Code generated by go-swagger; DO NOT EDIT.

package create_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUserHandlerFunc turns a function with the right signature into a post user handler
type PostUserHandlerFunc func(PostUserParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUserHandlerFunc) Handle(params PostUserParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostUserHandler interface for that can handle valid post user params
type PostUserHandler interface {
	Handle(PostUserParams, interface{}) middleware.Responder
}

// NewPostUser creates a new http.Handler for the post user operation
func NewPostUser(ctx *middleware.Context, handler PostUserHandler) *PostUser {
	return &PostUser{Context: ctx, Handler: handler}
}

/*
	PostUser swagger:route POST /user create_user postUser

create a user
*/
type PostUser struct {
	Context *middleware.Context
	Handler PostUserHandler
}

func (o *PostUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostUserParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
