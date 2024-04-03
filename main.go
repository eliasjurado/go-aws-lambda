package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eliasjurado/qhatu-user/awsgo"
	"github.com/eliasjurado/qhatu-user/db"
	"github.com/eliasjurado/qhatu-user/models"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(ExecLambda)
}

func ExecLambda(
	ctx context.Context,
	event events.CognitoEventUserPoolsPostConfirmation) (
	events.CognitoEventUserPoolsPostConfirmation,
	error) {

	awsgo.InitAWS()
	if !ValidateParam() {
		err := errors.New("Error en parámetros.\"SecretName\" no enviado")
		return event, err
	}

	var data models.SignUp
	for row, value := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = value
			log.Printf("Email = %+v\n", data.UserEmail)
		case "sub":
			data.UserUUID = value
			log.Printf("Sub = %+v\n", data.UserUUID)
		}
	}
	err := db.ReadSecret()
	if err != nil {
		log.Printf("Error al leer secret :%+v\n", err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err
}

func ValidateParam() bool {
	var isValid bool
	_, isValid = os.LookupEnv("SecretName")
	return isValid
}
