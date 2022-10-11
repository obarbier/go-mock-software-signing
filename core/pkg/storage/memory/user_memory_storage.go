package memory

import (
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/obarbier/custom-app/core/pkg/storage"
	"sync"
)

var (
	maxId int64
	us    map[int64]*models.User
)
var NotFoundError = fmt.Errorf("no user was found")

type UserStore struct {
	mx sync.RWMutex
}

func (m *UserStore) FindByUserName(username string) (*models.User, error) {
	//TODO greedy implementation make easier by creating index
	m.mx.Lock()
	defer m.mx.Unlock()

	for _, value := range us {
		if *value.UserName == username {
			return value, nil
		}
	}

	return nil, NotFoundError
}

func (m *UserStore) FindAllUser() ([]*models.User, error) {
	m.mx.Lock()
	defer m.mx.Unlock()
	v := make([]*models.User, 0, len(us))

	for _, value := range us {
		v = append(v, value)
	}

	return v, nil
}

func (m *UserStore) CreateUser(user *models.User) (*models.User, error) {
	m.mx.Lock()
	defer m.mx.Unlock()
	user.ID = maxId
	us[maxId] = user
	maxId++
	// TODO(obarbier): add unique username
	return user, nil
}

func (m *UserStore) ReadUser(i int64) (*models.User, error) {
	m.mx.Lock()
	defer m.mx.Unlock()
	return us[i], nil
}

func (m *UserStore) UpdateUser(user *models.User) error {
	m.mx.Lock()
	defer m.mx.Unlock()
	// TODO(obarbier) check if i in us
	us[user.ID] = user
	return nil
}

func (m *UserStore) DeleteUser(i int64) error {
	m.mx.Lock()
	defer m.mx.Unlock()
	// TODO(obarbier) check if i in us
	delete(us, i)
	return nil
}

var _ storage.UserStorage = &UserStore{}

func NewMemoryStorage() *UserStore {
	maxId = 1
	us = make(map[int64]*models.User)
	return &UserStore{
		sync.RWMutex{},
	}
}
