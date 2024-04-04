package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/eliasjurado/qhatu-user/models"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	secret, err := GetConnectionValues()
	if err != nil {
		log.Fatal(err.Error())
	}

	dns := fmt.Sprintf("%v:%v@(%v:%v)/%v",
		secret.Username,
		secret.Password,
		secret.Host,
		secret.Port,
		secret.DbClusterIdentifier)

	log.Printf("DNS value : %+v\n", dns)

	//Open database connection
	connection, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = connection

	//Verify db connection
	Ping()

	log.Printf("%+v\n", "Successful DB Connection")

}

func Ping() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err.Error())
	}
}

// Close db connection
func Close() {
	DB.Close()
}

func GetConnectionValues() (models.SecretRds, error) {
	secretName := "qhatuMySQL"
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString

	// Your code goes here.
	var secretRds = models.SecretRds{}
	err = json.Unmarshal([]byte(secretString), &secretRds)
	if err != nil {
		log.Fatal(err.Error())
		return secretRds, err
	}
	return secretRds, nil
}

// Crear una tabla en la base de datos
func CreateTable(schema string) {
	_, err := Exec(schema)
	if err != nil {
		fmt.Println(err)
	}
}

// Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := DB.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

// Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := DB.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
