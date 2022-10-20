package api

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/create_user"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/delete_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/get_all"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/get_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/update_user"
	"github.com/obarbier/custom-app/core/pkg/restapi/operations/update_user_by_id"
	"github.com/obarbier/custom-app/core/pkg/storage"
	"log"
)

var us *storage.UserService

// policyCache is the user policy mapping, mapping user id with policy node
var policyCache map[int64]*node
var l *log.Logger

func DeleteV1UserUserIDHandlerFunc(params delete_user_by_id.DeleteUserUserIDParams, principal interface{}) middleware.Responder {
	log_utils.Trace("handling delete a user", principal)
	_ = us.DeleteUser(params.HTTPRequest.Context(), params.UserID)
	return delete_user_by_id.NewDeleteUserUserIDOK()
}

func GetV1UserAllHandlerFunc(params get_all.GetUserAllParams, principal interface{}) middleware.Responder {
	log_utils.Trace("handling getting all user", principal)
	entities, _ := us.FindAllUser(params.HTTPRequest.Context())

	// TODO(obarbier): greedy implementation
	var res []*models.UserResponse

	for _, user := range entities {

		p := &models.UserResponse{
			ID:       user.ID,
			UserName: user.UserName,
		}

		res = append(res, p)
	}

	return get_all.NewGetUserAllOK().WithPayload(res)
}

func GetV1UserUserIDHandlerFunc(params get_user_by_id.GetUserUserIDParams, principal interface{}) middleware.Responder {
	log_utils.Trace("handling getting a user by id", principal)
	user, _ := us.ReadUser(params.HTTPRequest.Context(), params.UserID)
	if user == nil {
		return get_user_by_id.NewGetUserUserIDNotFound()
	}
	res := &models.UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}
	return get_user_by_id.NewGetUserUserIDOK().WithPayload(res)
}
func PostV1UserHandlerFunc(params create_user.PostUserParams, principal interface{}) middleware.Responder {
	log_utils.Trace("handling creating a user", principal)
	u := &models.User{
		Password: params.User.Password,
		UserName: params.User.UserName,
	}
	user, err := us.CreateUser(params.HTTPRequest.Context(), u)
	if err != nil {
		errMsg := err.Error()
		erroRes := &models.Error{
			Code:    500,
			Fields:  "",
			Message: &errMsg,
		}
		return create_user.NewPostUserInternalServerError().WithPayload(erroRes)
	}
	res := &models.UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}
	return create_user.NewPostUserCreated().WithPayload(res)
}

func PutUserUserIDHandlerFunc(params update_user_by_id.PutUserUserIDParams, principal interface{}) middleware.Responder {
	log_utils.Trace("handling updating a user by id", principal)
	user, _ := us.ReadUser(params.HTTPRequest.Context(), params.UserID)
	if user == nil {
		return update_user_by_id.NewPutUserUserIDNotFound()
	}

	//TODO(obarbier): handle partial update
	update := &models.User{
		ID:       user.ID,
		Password: params.User.Password,
		UserName: params.User.UserName,
	}
	err := us.UpdateUser(params.HTTPRequest.Context(), update)
	if err != nil {
		errMsg := err.Error()
		errorRes := &models.Error{
			Code:    500,
			Fields:  "",
			Message: &errMsg,
		}
		return update_user_by_id.NewPutUserUserIDInternalServerError().WithPayload(errorRes)
	}
	return update_user_by_id.NewPutUserUserIDOK()
}

func PutV1UserHandlerFunc(params update_user.PutUserParams, principal interface{}) middleware.Responder {
	log_utils.Trace("handling updating a user", principal)
	_ = us.UpdateUser(params.HTTPRequest.Context(), params.User)
	return update_user.NewPutUserOK()
}

func AuthenticateUser(username string, pass string) bool {
	user, err := us.FindByUserName(context.Background(), username)
	if err != nil {
		return false
	}
	return storage.CheckPasswordHash(pass, user.Password)
}
