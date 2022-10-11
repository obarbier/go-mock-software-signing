package api

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/obarbier/custom-app/core/pkg/models"
	"net/http"
	"strings"
)

func UserAuthorizationManager() runtime.Authorizer {
	return runtime.AuthorizerFunc(func(r *http.Request, p interface{}) error {
		l.Println(fmt.Printf("authorization for: %+v \n", p))
		l.Println(fmt.Printf("uri for: %+v \n", r.RequestURI))
		np, ok := p.(Principal)
		if !ok {
			l.Println(fmt.Sprintf("not a valid  principal"))
			return errors.New(401, "Unauthorized request")
		}

		// TODO(obarbier): remove this
		if *np.User.UserName == "admin" {
			return nil
		}
		path := strings.Replace(r.RequestURI, "/api/v1/", "", 1) // TODO(obarbier): this can be cleaner
		policy, err := us.GetPolicy(np.User, path)
		if err != nil {
			l.Println(fmt.Sprintf("error retrieving policy for user: %s", err))
			return errors.New(401, "Unauthorized request")

		}
		// TODO(obarbier): bitmap can be useful here
		switch r.Method {
		case "GET":
			for _, c := range policy.Capabilities {
				if c == models.CapabilityRead {
					return nil
				}
			}
		case "PUT":
			for _, c := range policy.Capabilities {
				if c == models.CapabilityUpdate {
					return nil
				}
			}
		case "DELETE":
			for _, c := range policy.Capabilities {
				if c == models.CapabilityDelete {
					return nil
				}
			}
		case "POST":
			for _, c := range policy.Capabilities {
				if c == models.CapabilityCreate {
					return nil
				}
			}
		default:
			l.Println(fmt.Sprintf("not valid http method"))
			return errors.New(401, "Unauthorized request")
		}
		return errors.New(401, "Unauthorized request")
	})
}
