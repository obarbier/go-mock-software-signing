package storage

import (
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserStorage interface {
	CreateUser(*models.User) (*models.User, error)
	ReadUser(int64) (*models.User, error)
	FindAllUser() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(int64) error
	FindByUserName(username string) (*models.User, error)
}

// TODO(obarbier): implement a storage layer for policy. Ideally this needs to be able to back up the tree
type PolicyStorage interface {
	CreatePolicy(models.Policy) (models.Policy, error)
	UpdatePolicy(models.Policy) (models.Policy, error)
	DeletePolicy(policy models.Policy) error
	GetPolicy(id int64) (models.Policy, error)
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
	// TODO(obarbier): consider encapsulating logging functionality
	logger *log.Logger
	impl   UserStorage
	// upm is the user policy mapping, mapping user id with policy node
	upm map[int64]*node
}

func (u *UserService) FindByUserName(username string) (*models.User, error) {
	u.logger.Println("started find by username")
	entity, err := u.impl.FindByUserName(username)
	if err != nil {
		u.logger.Println("failed to retrieve user")
		return nil, err
	}
	u.logger.Println("successfully retrieved user")
	return entity, nil
}

func (u *UserService) FindAllUser() ([]*models.User, error) {
	u.logger.Println("started getting all user")
	entity, err := u.impl.FindAllUser()
	if err != nil {
		u.logger.Println("failed to create a user")
		return nil, err
	}
	u.logger.Println("successfully created a user")
	return entity, nil
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	u.logger.Println("started creating a user")
	pwdHash, _ := HashPassword(user.Password)
	user.Password = pwdHash
	entity, err := u.impl.CreateUser(user)
	if err != nil {
		u.logger.Println("failed to create a user")
		return nil, err
	}
	u.logger.Println("successfully created a user")
	return entity, nil

}

func (u *UserService) ReadUser(i int64) (*models.User, error) {
	u.logger.Println("started reading user from storage layer")
	entity, err := u.impl.ReadUser(i)
	if err != nil {
		u.logger.Println("failed to retrieve a user")
		return nil, err
	}
	u.logger.Println("successfully retrieved user from storage layer")
	return entity, nil
}

func (u *UserService) UpdateUser(user *models.User) error {
	u.logger.Println("updating a user")
	pwdHash, _ := HashPassword(user.Password)
	user.Password = pwdHash
	err := u.impl.UpdateUser(user)
	if err != nil {
		u.logger.Println("failed to update a user")
		return err
	}
	u.logger.Println("successfully updated a user")
	return nil
}

func (u *UserService) DeleteUser(i int64) error {
	u.logger.Println("deleting a user")
	if err := u.impl.DeleteUser(i); err != nil {
		u.logger.Println("failed to delete a user")
		return err
	}
	u.logger.Println("successfully deleted a user")
	return nil
}

func (u *UserService) CreateOrUpdatePolicy(user *models.User, policy *models.Policy) error {
	// check if upm has policy set for user
	n, ok := u.upm[user.ID]
	if !ok {
		// create new policy
		newN := newNode()

		for path, acl := range *policy {
			newN.insert(path, acl)
		}
		u.upm[user.ID] = newN
		return nil
	}

	// update policy
	// TODO(obarbier): check policy for path exist and create or update
	for path, acl := range *policy {
		n.insert(path, acl)
	}

	return nil
}

func (u *UserService) GetPolicy(user *models.User, path string) (models.PolicyAnon, error) {
	n, ok := u.upm[user.ID]
	if !ok {
		return models.PolicyAnon{}, fmt.Errorf("policy not set for user" /* TODO(obarbier): better error */)
	}
	p := n.getPolicy(path)
	if p == nil {
		return models.PolicyAnon{}, fmt.Errorf("policy not defined for path" /* TODO(obarbier): better error */)
	}

	return p.acls, nil
}

// NewUserService create a new UserService
func NewUserService(impl UserStorage) *UserService {
	return &UserService{
		logger: log.Default(),
		impl:   impl,
		upm:    make(map[int64]*node),
	}
}

var _ UserStorage = &UserService{}
