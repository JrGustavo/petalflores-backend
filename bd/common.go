package bd

import (
	"os"
	"petal-backend/models"
	"petal-backend/secretm"
)

var SecretModel models.SecretRDSJaon
var err error

func ReadSecret() error {
	SecretModel, err := secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
