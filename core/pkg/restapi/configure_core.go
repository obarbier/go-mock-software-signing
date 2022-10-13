// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	api2 "github.com/obarbier/custom-app/core/pkg/api"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/obarbier/custom-app/core/pkg/restapi/operations"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/create_user"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/delete_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/get_all"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/get_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/update_user"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/update_user_by_id"
)

//go:generate swagger generate server --target ../../pkg --name Core --spec ../../swagger.yml --principal interface{}

func configureFlags(api *operations.CoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api2.NewAPI(api)
	// Applies when the Authorization header is set with the Basic scheme
	if api.BasicAuthAuth == nil {
		api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
			return nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.DeleteUserByIDDeleteUserUserIDHandler == nil {
		api.DeleteUserByIDDeleteUserUserIDHandler = delete_user_by_id.DeleteUserUserIDHandlerFunc(func(params delete_user_by_id.DeleteUserUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation delete_user_by_id.DeleteUserUserID has not yet been implemented")
		})
	}
	if api.GetAllGetUserAllHandler == nil {
		api.GetAllGetUserAllHandler = get_all.GetUserAllHandlerFunc(func(params get_all.GetUserAllParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation get_all.GetUserAll has not yet been implemented")
		})
	}
	if api.GetUserByIDGetUserUserIDHandler == nil {
		api.GetUserByIDGetUserUserIDHandler = get_user_by_id.GetUserUserIDHandlerFunc(func(params get_user_by_id.GetUserUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation get_user_by_id.GetUserUserID has not yet been implemented")
		})
	}
	if api.CreateUserPostUserHandler == nil {
		api.CreateUserPostUserHandler = create_user.PostUserHandlerFunc(func(params create_user.PostUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation create_user.PostUser has not yet been implemented")
		})
	}
	if api.UpdateUserPutUserHandler == nil {
		api.UpdateUserPutUserHandler = update_user.PutUserHandlerFunc(func(params update_user.PutUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation update_user.PutUser has not yet been implemented")
		})
	}
	if api.UpdateUserByIDPutUserUserIDHandler == nil {
		api.UpdateUserByIDPutUserUserIDHandler = update_user_by_id.PutUserUserIDHandlerFunc(func(params update_user_by_id.PutUserUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation update_user_by_id.PutUserUserID has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
