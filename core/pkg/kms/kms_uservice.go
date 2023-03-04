package kms

import "github.com/obarbier/custom-app/core/pkg/models"

type Storage interface {
	// Create will create a key and user wrapping algorithm for safe manipulation
	Store([]byte) error
	Get(id int64) ([]byte, error)
}

type Service struct {
}

func (s *Service) Create(key *models.KeyType) (*models.Key, error) {

}

func (s *Service) wrap(key *models.Key) ([]byte, error) {

}

func (s *Service) unwrap([]byte) (*models.Key, error) {

}
