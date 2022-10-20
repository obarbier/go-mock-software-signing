package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	ReadUser(context.Context, int64) (*models.User, error)
	FindAllUser(context.Context) ([]*models.User, error)
	UpdateUser(context.Context, *models.User) error
	DeleteUser(context.Context, int64) error
	FindByUserName(context.Context, string) (*models.User, error)

	// TODO(obarbier): implement a storage layer for policy. Ideally this needs to be able to back up the tree
	CreatePolicy(context.Context, int64, *models.Policy) error
	GetPolicy(context.Context, int64) (*models.Policy, error)
}

const DefaultPolicyFormat = `
{
	"user/%d":{
				"capabilities": [ "read", "update"]
			}
}
`

func SetDefaultPolicy(id int64) string {
	return fmt.Sprintf(DefaultPolicyFormat, id)
}

// TODO(obarbier): can we user golang custom decoder to do this
func UnmarshalPolicy(data []byte) *models.Policy {
	var v models.Policy
	err := json.Unmarshal(data, &v)
	if err != nil {
		log_utils.Error(err)
	}
	return &v
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	log_utils.Debug("length of bcrypt bytes", len(bytes))
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// UserService is a wrapper to storage layer it allow storage type to exist without worry of business logic
type UserService struct {
	impl UserStorage
}

func (u *UserService) FindByUserName(ctx context.Context, username string) (*models.User, error) {
	log_utils.Info("started find by username")
	entity, err := u.impl.FindByUserName(ctx, username)
	if err != nil {
		log_utils.Error("failed to retrieve user")
		return nil, err
	}
	log_utils.Info("successfully retrieved user")
	return entity, nil
}

func (u *UserService) FindAllUser(ctx context.Context) ([]*models.User, error) {
	log_utils.Info("started getting all user")
	entity, err := u.impl.FindAllUser(ctx)
	if err != nil {
		log_utils.Error("failed to create a user")
		return nil, err
	}
	log_utils.Info("successfully created a user")
	return entity, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	log_utils.Info("started creating a user")
	pwdHash, _ := HashPassword(user.Password)
	user.Password = pwdHash
	entity, err := u.impl.CreateUser(ctx, user)
	if err != nil {
		log_utils.Error("failed to create a user")
		return nil, err
	}
	// add policy
	defaultPolicy := SetDefaultPolicy(user.ID)
	err = u.impl.CreatePolicy(ctx, user.ID, UnmarshalPolicy([]byte(defaultPolicy)))
	if err != nil {
		return nil, err
	}
	log_utils.Info("successfully created a user")
	return entity, nil

}

func (u *UserService) ReadUser(ctx context.Context, i int64) (*models.User, error) {
	log_utils.Info("started reading user from storage layer")
	entity, err := u.impl.ReadUser(ctx, i)
	if err != nil {
		log_utils.Error("failed to retrieve a user")
		return nil, err
	}
	log_utils.Info("successfully retrieved user from storage layer")
	return entity, nil
}

func (u *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	log_utils.Info("updating a user")
	pwdHash, _ := HashPassword(user.Password)
	user.Password = pwdHash
	err := u.impl.UpdateUser(ctx, user)
	if err != nil {
		log_utils.Error("failed to update a user")
		return err
	}
	log_utils.Info("successfully updated a user")
	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, i int64) error {
	log_utils.Info("deleting a user")
	if err := u.impl.DeleteUser(ctx, i); err != nil {
		log_utils.Error("failed to delete a user")
		return err
	}
	log_utils.Info("successfully deleted a user")
	return nil
}

// NewUserService create a new UserService
func NewUserService(impl UserStorage) *UserService {
	return &UserService{
		impl: impl,
	}
}
