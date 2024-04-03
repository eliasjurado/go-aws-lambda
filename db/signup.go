package db

import (
	"fmt"
	"log"

	"github.com/eliasjurado/qhatu-user/models"
	"github.com/eliasjurado/qhatu-user/tools"
)

func SignUp(s models.SignUp) error {
	log.Printf("%+v\n", "Registration begins")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := fmt.Sprintf("insert into users (User_Email,User_UUID, User_DateAdd) VALUES ('%v','%v','%v')", s.UserEmail, s.UserUUID, tools.MySqlDateTimeNow())

	_,err = Db.Exec(query)
	if err != nil {
		log.Printf("%+v\n",err.Error() )
	}
	return nil
}
