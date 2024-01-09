package secretm

import (
	"encoding/json"
	"fmt"
	"petal-backend/aws/aws-sdk-go-v2/service/secretsmanager"
	"petal-backend/awsgo"
	"petal-backend/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJaon, error) {
	var datosSecret models.SecretRDSJaon
	fmt.Println(" > Pido Secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		secretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println("Error al obtener el secreto " + err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Secreto obtenido " + datosSecret.Username)
	return datosSecret, nil
}
