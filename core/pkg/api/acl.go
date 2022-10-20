package api

import (
	"context"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
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

		path := strings.Replace(r.RequestURI, "/api/v1/", "", 1) // TODO(obarbier): this can be cleaner
		isAuthorize, err := Authorize(r.Context(), np.User, path, r.Method)
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

func Authorize(ctx context.Context, user *models.User, path, httpMethod string) (bool, error) {
	log_utils.Info("trying to retrieve policy from cache")
	n, ok := policyCache[user.ID] // TODO:(obarbier): cache TTL
	// TODO(obarbier): a goroutine should be implemented to update policyCache regularly based on TTL and or Schedule
	if !ok {
		log_utils.Info("cache missed. getting data from db")
		p, err := us.FindByUserName(ctx, *user.UserName)
		if err != nil {
			return false, err
		}
		if p == nil {
			return false, fmt.Errorf("policy not set for user" /* TODO(obarbier): better error */)
		}
		n = newPolicy(&p.Policy)
		log_utils.Info("updating policy cache")
		policyCache[user.ID] = n
	}
	p := n.getPolicy(path)
	if p == nil {
		return false, fmt.Errorf("policy not defined for path" /* TODO(obarbier): better error */)
	}
	return p.capability&HTTPMethodMatch[httpMethod] != 0, nil

}
