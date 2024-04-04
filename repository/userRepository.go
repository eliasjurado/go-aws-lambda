package repository

import (
	"log"

	"github.com/eliasjurado/qhatu-user/db"
	"github.com/eliasjurado/qhatu-user/models"
)

func CreateUser(c models.SignUpCognito) {
	query := "insert into users (UserUUID, UserEmail) values (?,?)"

	_, err := db.Exec(query, c.UserUUID, c.UserEmail)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("New user registered: %s\n", c.UserEmail)
}
