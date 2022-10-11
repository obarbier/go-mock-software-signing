package mysql

import (
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/obarbier/custom-app/core/pkg/storage"
)

type UserStorage struct {
}

func (u UserStorage) CreateUser(user *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserStorage) ReadUser(i int64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserStorage) FindAllUser() ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserStorage) UpdateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserStorage) DeleteUser(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (u UserStorage) FindByUserName(username string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

var _ storage.UserStorage = &UserStorage{}
