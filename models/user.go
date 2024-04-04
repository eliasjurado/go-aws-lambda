package models

const UserSchema string = `CREATE TABLE IF NOT EXISTS users (
	UserUUID char(36) NOT NULL PRIMARY KEY,
	UserEmail varchar(100) NOT NULL DEFAULT '',
	UserFirstName varchar(20) DEFAULT NULL,
	UserLastName varchar(20) DEFAULT NULL,
	UserStatus char(1) NOT NULL DEFAULT '0',
	UserCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	UserUpdatedAt TIMESTAMP DEFAULT NULL,
	UNIQUE (User_Email)
	);`

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
