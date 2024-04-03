package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/eliasjurado/qhatu-user/models"
	"github.com/eliasjurado/qhatu-user/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		log.Printf("Connection: %+v\n", err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		log.Printf("Ping: %+v\n", err.Error())
		return err
	}

	log.Printf("%+v\n", "Successful connection to DB")
	return nil
}

func ConnStr(data models.SecretRDSJson) string {
	user := data.Username
	pass := data.Password
	host := data.Host
	dbName := "qhatu"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", user, pass, host, dbName)
	log.Printf("%+v\n", dsn)
	return dsn
}
