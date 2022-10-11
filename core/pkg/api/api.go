package api

import (
	"github.com/go-openapi/errors"
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/create_user"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/delete_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/get_all"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/get_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/update_user"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/update_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/storage"
	"github.com/obarbier/custom-app/core/pkg/storage/memory"
	"log"
)

type Principal struct {
	User *models.User
}

func NewAPI(api *operations.CoreAPI) error {
	us = storage.NewUserService(memory.NewMemoryStorage())
	l = log.Default()
	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
		if user == "admin" && pass == "pass" {
			adminUser := "admin"
			p := Principal{
				User: &models.User{
					UserName: &adminUser,
				},
			}
			return p, nil
		}
		if AuthenticateUser(user, pass) {
			u, err := us.FindByUserName(user)
			if err != nil {
				return nil, errors.New(401, "invalid username and or password")
			}
			p := Principal{
				User: u,
			}
			return p, nil
		}
		return nil, errors.New(401, "invalid username and or password")
	}
	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer

	api.APIAuthorizer = UserAuthorizationManager()

	// USER API
	api.DeleteUserByIDDeleteUserUserIDHandler = delete_user_by_id.DeleteUserUserIDHandlerFunc(DeleteV1UserUserIDHandlerFunc)
	api.GetAllGetUserAllHandler = get_all.GetUserAllHandlerFunc(GetV1UserAllHandlerFunc)
	api.GetUserByIDGetUserUserIDHandler = get_user_by_id.GetUserUserIDHandlerFunc(GetV1UserUserIDHandlerFunc)
	api.CreateUserPostUserHandler = create_user.PostUserHandlerFunc(PostV1UserHandlerFunc)
	api.UpdateUserPutUserHandler = update_user.PutUserHandlerFunc(PutV1UserHandlerFunc)
	api.UpdateUserByIDPutUserUserIDHandler = update_user_by_id.PutUserUserIDHandlerFunc(PutUserUserIDHandlerFunc)
	return nil
}
