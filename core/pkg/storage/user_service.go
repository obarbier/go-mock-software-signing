package storage

import (
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreateUser(*models.User) (*models.User, error)
	ReadUser(int64) (*models.User, error)
	FindAllUser() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(int64) error
	FindByUserName(username string) (*models.User, error)

	// TODO(obarbier): implement a storage layer for policy. Ideally this needs to be able to back up the tree
	CreatePolicy(int64, *models.Policy) error
	GetPolicy(int64) (*models.Policy, error)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// UserService is a wrapper to storage layer it allow storage type to exist without worry of business logic
type UserService struct {
	impl UserStorage
	// policyCache is the user policy mapping, mapping user id with policy node
	policyCache map[int64]*node
}

func (u *UserService) FindByUserName(username string) (*models.User, error) {
	log_utils.Info("started find by username")
	entity, err := u.impl.FindByUserName(username)
	if err != nil {
		log_utils.Error("failed to retrieve user")
		return nil, err
	}
	log_utils.Info("successfully retrieved user")
	return entity, nil
}

func (u *UserService) FindAllUser() ([]*models.User, error) {
	log_utils.Info("started getting all user")
	entity, err := u.impl.FindAllUser()
	if err != nil {
		log_utils.Error("failed to create a user")
		return nil, err
	}
	log_utils.Info("successfully created a user")
	return entity, nil
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	log_utils.Info("started creating a user")
	pwdHash, _ := HashPassword(user.Password)
	user.Password = pwdHash
	entity, err := u.impl.CreateUser(user)
	if err != nil {
		log_utils.Error("failed to create a user")
		return nil, err
	}
	log_utils.Info("successfully created a user")
	return entity, nil

}

func (u *UserService) ReadUser(i int64) (*models.User, error) {
	log_utils.Info("started reading user from storage layer")
	entity, err := u.impl.ReadUser(i)
	if err != nil {
		log_utils.Error("failed to retrieve a user")
		return nil, err
	}
	log_utils.Info("successfully retrieved user from storage layer")
	return entity, nil
}

func (u *UserService) UpdateUser(user *models.User) error {
	log_utils.Info("updating a user")
	pwdHash, _ := HashPassword(user.Password)
	user.Password = pwdHash
	err := u.impl.UpdateUser(user)
	if err != nil {
		log_utils.Error("failed to update a user")
		return err
	}
	log_utils.Info("successfully updated a user")
	return nil
}

func (u *UserService) DeleteUser(i int64) error {
	log_utils.Info("deleting a user")
	if err := u.impl.DeleteUser(i); err != nil {
		log_utils.Error("failed to delete a user")
		return err
	}
	log_utils.Info("successfully deleted a user")
	return nil
}

func (u *UserService) CreateOrUpdatePolicy(user *models.User, policy *models.Policy) error {
	u.impl.CreatePolicy(user.ID, policy)
	return nil
}

func (u *UserService) Authorize(user *models.User, path, httpMethod string) (bool, error) {
	log_utils.Info("trying to retrieve policy from cache")
	n, ok := u.policyCache[user.ID] // TODO:(obarbier): cache TTL
	// TODO(obarbier): a goroutine should be implemented to update policyCache regularly based on TTL and or Schedule
	if !ok {
		log_utils.Info("cache missed. getting data from db")
		p, err := u.impl.GetPolicy(user.ID)
		if err != nil {
			return false, err
		}
		if p == nil {
			return false, fmt.Errorf("policy not set for user" /* TODO(obarbier): better error */)
		}
		n = newPolicy(p)
		log_utils.Info("updating policy cache")
		u.policyCache[user.ID] = n
	}
	p := n.getPolicy(path)
	if p == nil {
		return false, fmt.Errorf("policy not defined for path" /* TODO(obarbier): better error */)
	}
	return p.capability&HTTPMethodMatch[httpMethod] != 0, nil

}

// NewUserService create a new UserService
func NewUserService(impl UserStorage) *UserService {
	return &UserService{
		impl:        impl,
		policyCache: make(map[int64]*node),
	}
}
