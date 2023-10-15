package vault

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/outofoffice3/common/logger"
)

// Vault represents an interface for secret management.
type Vault interface {
	// GetSecret retrieves a secret from the vault by its name.
	GetSecret(secretName string) (interface{}, error)

	// PutSecret stores a secret in the vault under the given name.
	PutSecret(secretName string, secretValue interface{}) error

	// DeleteSecret deletes a secret from the vault by its name.
	DeleteSecret(secretName string) error

	// GetLocalSecret retrieves a locally stored secret by its name.
	GetLocalSecret(secretName string) (interface{}, bool)

	// PutLocalSecret stores a secret locally within the object.
	PutLocalSecret(secretName string, secretValue interface{}) error

	// DeleteLocalSecret deletes a locally stored secret by its name.
	DeleteLocalSecret(secretName string) error
}

type _AwsVault struct {
	secrets map[string]interface{}
	client  *secretsmanager.Client
}

var (
	sos logger.Logger
)

func Init() {
	sos = logger.NewConsoleLogger(logger.LogLevelInfo)
	sos.Infof("init aws vault")
}

// set log level
func SetLogLevel(level logger.LogLevel) {
	sos.SetLogLevel(level)
}

// get secret from vault by its name
func (awsv *_AwsVault) GetSecret(secretName string) (interface{}, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	}
	sos.Debugf("getting secret [%s] from vault", secretName)
	result, err := awsv.client.GetSecretValue(context.Background(), input)
	if err != nil {
		return "", errors.New("failed to get secret value")
	}
	sos.Debugf("secret retrieved: %s", *result.SecretString)
	return *result.SecretString, nil
}

// put secret in vault by its name
func (awsv *_AwsVault) PutSecret(secretName string, secretValue interface{}) error {
	sv, ok := secretValue.(string)
	if !ok {
		return errors.New("secret value is not string")
	}
	input := &secretsmanager.CreateSecretInput{
		Name:         &secretName,
		SecretString: &sv,
	}
	sos.Debugf("putting secret [%s] with value [%v]", secretName, secretValue)
	result, err := awsv.client.CreateSecret(context.Background(), input)
	if err != nil {
		return errors.New("failed to create secret")
	}
	sos.Debugf("secret [%s] created w/ ARN [%v]", *result.Name, result.ARN)
	return nil
}

// delete secret from vault by its name
func (awsv *_AwsVault) DeleteSecret(secretName string) error {
	return nil
}

// get local secret by its name
func (awsv *_AwsVault) GetLocalSecret(secretName string) interface{} {
	sos.Debugf("getting secret [%s] from local cache", secretName)
	// check if secret exists in cache
	result := awsv.secrets[secretName]
	// if secret exists in cache, return it
	sos.Debugf("checking if secret [%s], exists. found [%v]", secretName, result)
	return result
}

// put local secret by its name
func (awsv *_AwsVault) PutLocalSecret(secretName string, secretValue interface{}) {
	sos.Debugf("putting secret [%s] with value [%s] in local cache", secretName, secretValue)
	awsv.secrets[secretName] = secretValue
	return
}

// delete local secret by its name
func (awsv *_AwsVault) DeleteLocalSecret(secretName string) error {
	sos.Debugf("deleting secret [%s]", secretName)
	// check if secret exists")
	delete(awsv.secrets, secretName)
	return nil
}

// creates a new Vault
func NewVault() (*_AwsVault, error) {
	// load config
	cfg, err := config.LoadDefaultConfig(context.Background())
	// return errors
	if err != nil {
		return nil, errors.New("failed to load aws config")
	}
	// initialize vault and return
	client := secretsmanager.NewFromConfig(cfg)
	return &_AwsVault{
		secrets: make(map[string]interface{}),
		client:  client,
	}, nil
}
