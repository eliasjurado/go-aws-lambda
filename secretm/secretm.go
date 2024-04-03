package secretm

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/eliasjurado/qhatu-user/awsgo"
	"github.com/eliasjurado/qhatu-user/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
var secretRDSJson = models.SecretRDSJson{}
svc := secretsmanager.NewFromConfig(awsgo.Cfg)
pass, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
SecretId: aws.String(secretName),
})
if err != nil {
	log.Fatal(err.Error())
	return secretRDSJson, err
}

err = json.Unmarshal([]byte(*pass.SecretString), &secretRDSJson)
if err != nil {
	log.Fatal(err.Error())
	return secretRDSJson, err
}

return secretRDSJson,nil
}