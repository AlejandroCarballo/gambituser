package main

import (
	"context"
	"errors"
	"fmt"
	"gambituser/awsgo"
	"gambituser/bd"
	"gambituser/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parámestros. debe enviar 'SecretName'")
		err := errors.New("error en los parámentros debe enviar SecretName")
		return event, err
	}
	var datos models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email =" + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Email =" + datos.UserUUID)

		}
	}
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret" + err.Error())
		return event, err
	}
	err = bd.SignUp(datos)
	return event, err

}

func ValidoParametros() bool {
	var traeParam bool
	_, traeParam = os.LookupEnv("secretName")
	return traeParam
}
