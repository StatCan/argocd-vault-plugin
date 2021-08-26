package vault

import (
	"github.com/IBM/argocd-vault-plugin/pkg/utils"
	"github.com/hashicorp/vault/api"
)

// GithubAuth is a struct for working with Vault that uses the Github Auth method

type TokenAuth struct {
	AccessToken string
}

// NewGithubAuth initializes a new GithubAuth with token
func NewTokenAuth(token string) *TokenAuth {
	tokenAuth := &TokenAuth{
		AccessToken: token,
	}

	return tokenAuth
}

// Authenticate authenticates with Vault and returns a token
func (t *TokenAuth) Authenticate(vaultClient *api.Client) error {
	payload := map[string]interface{}{
		"token": t.AccessToken,
	}

	data, err := vaultClient.Logical().Write("sys/auth/vaultAuth", payload)
	if err != nil {
		return err
	}
	

	// If we cannot write the Vault token, we'll just have to login next time. Nothing showstopping.
	err = utils.SetToken(vaultClient, data.Auth.ClientToken)
	if err != nil {
		print(err)
	}

	return nil
}
