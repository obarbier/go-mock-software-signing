package software

import (
	"crypto/rsa"
	"crypto/x509"
	"github.com/obarbier/custom-app/core/pkg/kms"
	"github.com/obarbier/custom-app/core/pkg/models"
	"os"
)

type KMS struct {
}

func (K *KMS) Create(keyType *models.KeyType) (*models.Key, error) {
	key, err := rsa.GenerateKey(os.Stdout, 4096)
	if err != nil {
		return nil, err
	}

	res := &models.Key{
		Data: map[string]interface{}{
			"privateKey": x509.MarshalPKCS1PrivateKey(key),
		},
		ID:      0,
		KeyType: keyType,
	}

	return res, nil

}

var _ kms.Storage = &KMS{}
