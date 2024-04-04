// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/
package main

import (
	"github.com/eliasjurado/qhatu-user/db"
	"github.com/eliasjurado/qhatu-user/models"
)

func main() {
	db.CreateTable(models.UserSchema)
}
