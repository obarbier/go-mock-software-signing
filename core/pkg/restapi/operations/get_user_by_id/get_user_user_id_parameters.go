// Code generated by go-swagger; DO NOT EDIT.

package get_user_by_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetUserUserIDParams creates a new GetUserUserIDParams object
//
// There are no default values defined in the spec.
func NewGetUserUserIDParams() GetUserUserIDParams {

	return GetUserUserIDParams{}
}

// GetUserUserIDParams contains all the bound params for the get user user ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetUserUserID
type GetUserUserIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*id of the user
	  Required: true
	  Minimum: 1
	  In: path
	*/
	UserID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUserUserIDParams() beforehand.
func (o *GetUserUserIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUserID, rhkUserID, _ := route.Params.GetOK("userId")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *GetUserUserIDParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("userId", "path", "int64", raw)
	}
	o.UserID = value

	if err := o.validateUserID(formats); err != nil {
		return err
	}

	return nil
}

// validateUserID carries on validations for parameter UserID
func (o *GetUserUserIDParams) validateUserID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("userId", "path", o.UserID, 1, false); err != nil {
		return err
	}

	return nil
}
