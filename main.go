package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"petal-backend/bd"
	"petal-backend/models"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"petal-backend/awsgo"
)

func main() {
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros. debe enviar 'SecretName")
		err := errors.New("error en los parametros debe enviar SecretName")
		return event, err
	}

	var datos models.Signup

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)

		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leet el Secret " + err.Error())
		return event, err
	}

}

func ValidoParametros() bool {
	var traeParametro bool = true
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
