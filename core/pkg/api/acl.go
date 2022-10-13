package api

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"net/http"
	"strings"
)

func UserAuthorizationManager() runtime.Authorizer {
	return runtime.AuthorizerFunc(func(r *http.Request, p interface{}) error {
		l.Println(fmt.Sprintf("authorization for: %+v \n", p))
		l.Println(fmt.Sprintf("uri for: %+v \n", r.RequestURI))
		np, ok := p.(Principal)
		if !ok {
			l.Println(fmt.Sprintf("not a valid  principal"))
			return errors.New(401, "Unauthorized request")
		}

		// TODO(obarbier): remove this
		if *np.User.UserName == "admin" {
			return nil
		}

		// TODO:
		path := strings.Replace(r.RequestURI, "/api/v1/", "", 1) // TODO(obarbier): this can be cleaner
		isAuthorize, err := us.Authorize(np.User, path, r.Method)
		if err != nil {
			l.Println(fmt.Sprintf("error retrieving policy for user: %s", err))
			return errors.New(401, "Unauthorized request")

		}
		if !isAuthorize {
			return errors.New(401, "Unauthorized request")
		}
		return nil
	})
}
