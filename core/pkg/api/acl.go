package api

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"net/http"
	"strings"
)

func UserAuthorizationManager() runtime.Authorizer {
	return runtime.AuthorizerFunc(func(r *http.Request, p interface{}) error {
		log_utils.Trace("logging request", []interface{}{r.RequestURI, p})
		np, ok := p.(Principal)
		if !ok {
			log_utils.Trace("not a valid  principal. therefore unauthorized request")
			return errors.New(401, "Unauthorized request")
		}

		// TODO(obarbier): remove this
		if *np.User.UserName == "admin" {
			log_utils.Trace("admin user")
			return nil
		}

		// TODO:
		path := strings.Replace(r.RequestURI, "/api/v1/", "", 1) // TODO(obarbier): this can be cleaner
		isAuthorize, err := us.Authorize(np.User, path, r.Method)
		if err != nil {
			log_utils.Trace("authorization failure for path", path)
			return errors.New(401, "Unauthorized request")

		}
		if !isAuthorize {
			log_utils.Trace("unauthorized principal", path, np.User.UserName)
			return errors.New(401, "Unauthorized request")
		}
		return nil
	})
}
