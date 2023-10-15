# Vault 

The vault Go package provides an interface and an implementation for secret management. It allows you to interact with AWS Secrets Manager to securely store and retrieve secrets. This package also provides the capability to store secrets locally within the object for faster access.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Interface](#interface)
- [Examples](#examples)
- [License](#license)

### Installation 
To install the vault package, you can use the go get command:

```sh 
go get github.com/outofoffice3/vault
```

### Usage
To use the vault package, you will need to have the AWS SDK for Go installed and properly configured with your AWS credentials. You can do this by setting up the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables or by using other AWS SDK configuration methods.

Import the go package into your code : 

``` go 
import "github.com/outofoffice3/vault"
```

You can then create a new Vault instance and start using the provided methods.

```go 
// Initialize the Vault
v, err := vault.NewVault()
if err != nil {
    log.Fatalf("Error: %v", err)
}

// Get a secret from the vault
secret, err := v.GetSecret("my-secret")
if err != nil {
    log.Fatalf("Error: %v", err)
}

// Put a secret in the vault
err = v.PutSecret("my-secret", "my-secret-value")
if err != nil {
    log.Fatalf("Error: %v", err)
}

// Delete a secret from the vault
err = v.DeleteSecret("my-secret")
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

### Interface
The Vault interface defines the methods for interacting with secrets. It includes the following methods:

- `GetSecret(secretName string) (interface{}, error)`: Retrieves a secret from the vault by its name.

- `PutSecret(secretName string, secretValue interface{})` error: Stores a secret in the vault under the given name.

- `DeleteSecret(secretName string) error`: Deletes a secret from the vault by its name.

- `GetLocalSecret(secretName string) (interface{}, bool)`: Retrieves a locally stored secret by its name.

- `PutLocalSecret(secretName string, secretValue interface{}) error`: Stores a secret locally within the object.

- `DeleteLocalSecret(secretName string) error`: Deletes a locally stored secret by its name.

### License
This package is released under the MIT License. See the LICENSE file for details.

Thank you for using the vault Go package. We hope it helps you manage your secrets effectively and securely.

