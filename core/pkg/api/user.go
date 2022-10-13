package api

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
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
var l *log.Logger

func DeleteV1UserUserIDHandlerFunc(params delete_user_by_id.DeleteUserUserIDParams, principal interface{}) middleware.Responder {
	l.Println(fmt.Sprintf("handling authorization for delete user %s", principal))
	_ = us.DeleteUser(params.UserID)
	return delete_user_by_id.NewDeleteUserUserIDOK()
}

func GetV1UserAllHandlerFunc(_ get_all.GetUserAllParams, principal interface{}) middleware.Responder {
	l.Println(fmt.Sprintf("handling authorization for getting all user %s", principal))
	entities, _ := us.FindAllUser()

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
	l.Println(fmt.Sprintf("handling authorization for getting a user %s", principal))
	user, _ := us.ReadUser(params.UserID)
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
	l.Println(fmt.Sprintf("handling authorization for creating a user %s", principal))
	u := &models.User{
		Password: params.User.Password,
		UserName: params.User.UserName,
	}
	user, _ := us.CreateUser(u)

	// add policy
	defaultPolicy := storage.SetDefaultPolicy(user.ID)

	us.CreateOrUpdatePolicy(user, storage.UnmarshalPolicy([]byte(defaultPolicy)))

	res := &models.UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}
	return create_user.NewPostUserCreated().WithPayload(res)
}

func PutUserUserIDHandlerFunc(params update_user_by_id.PutUserUserIDParams, principal interface{}) middleware.Responder {
	l.Println(fmt.Sprintf("handling authorization for getting a user %s\n", principal))
	user, _ := us.ReadUser(params.UserID)
	if user == nil {
		return update_user_by_id.NewPutUserUserIDNotFound()
	}

	//TODO(obarbier): handle partial update
	update := &models.User{
		ID:       user.ID,
		Password: params.User.Password,
		UserName: params.User.UserName,
	}
	us.UpdateUser(update)
	return update_user_by_id.NewPutUserUserIDOK()
}

func PutV1UserHandlerFunc(params update_user.PutUserParams, principal interface{}) middleware.Responder {
	l.Println(fmt.Sprintf("handling authorization for updating a user %s\n", principal))
	_ = us.UpdateUser(params.User)
	return update_user.NewPutUserOK()
}

func AuthenticateUser(username string, pass string) bool {
	user, err := us.FindByUserName(username)
	if err != nil {
		return false
	}
	return storage.CheckPasswordHash(pass, user.Password)
}
