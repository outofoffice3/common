package tests

import (
	"testing"

	"github.com/outofoffice3/common/vault"
	"github.com/stretchr/testify/assert"
)

func TestVaultLocalSecret(t *testing.T) {
	assertion := assert.New(t)
	vault.Init()
	vault, err := vault.NewVault()
	assert.NoError(t, err)
	assert.NotNil(t, vault)

	// put test secret in vault
	secretName := "mySecret"
	secretValue := "mySecretValue"
	vault.PutLocalSecret(secretName, secretValue)

	// retrieve secret from vault
	retrievedSecret := vault.GetLocalSecret(secretName)
	assertion.Equal(secretValue, retrievedSecret)
	assertion.IsType(secretValue, retrievedSecret)

	// delete secret from vault
	vault.DeleteLocalSecret(secretName)
	retrievedSecret = vault.GetLocalSecret(secretName)
	assertion.Empty(retrievedSecret)
}
