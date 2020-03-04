package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func (w *wrapper) canDelete(path string) bool {
	if _, err := w.client.Logical().Delete(path); err != nil {
		return false
	}

	return true
}

func (w *wrapper) canRead(path string) bool {
	if _, err := w.client.Logical().Read(path); err != nil {
		return false
	}

	return true
}

func (w *wrapper) canWrite(path string) bool {
	// The v2 of kv secret engine needs this
	data := make(map[string]interface{})
	info := map[string]string{
		"test": "test",
	}

	data["data"] = info

	if _, err := w.client.Logical().Write(path, data); err != nil {
		return false
	}

	return true
}

func (w *wrapper) loginAs(role string) error {
	if role == "" {
		return errors.New("A role is needed")
	}

	// Set the token in w, to query to Vault
	w.client.SetToken(w.authorizer)

	// Get role-id
	path := fmt.Sprintf("auth/approle/role/%s/role-id", role)
	secret, err := w.client.Logical().Read(path)
	if err != nil {
		return err
	}

	roleId := secret.Data["role_id"].(string)

	// Get secret-id
	path = fmt.Sprintf("auth/approle/role/%s/secret-id", role)
	secret, err = w.client.Logical().Write(path, nil)
	if err != nil {
		return err
	}

	secretId := secret.Data["secret_id"].(string)

	options := map[string]interface{}{
		"role_id":   roleId,
		"secret_id": secretId,
	}

	// Login with roleId and secretId
	secret, err = w.client.Logical().Write(approleLogin, options)
	if err != nil {
		return err
	}

	token, err := secret.TokenID()
	if err != nil {
		return err
	}

	w.client.SetToken(token)

	return nil
}

func newWrapper(config *api.Config) (*wrapper, error) {
	if config == nil {
		return defaultWrapper()
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &wrapper{
		client: client,
	}, nil
}

func defaultWrapper() (*wrapper, error) {
	config := &api.Config{
		Address: url,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &wrapper{
		client:   client,
		username: username,
		password: password,
	}, nil
}

func (w *wrapper) loginWithUserPassword() error {
	path := fmt.Sprintf("auth/userpass/login/%s", w.username)
	options := map[string]interface{}{
		"password": w.password,
	}

	secret, err := w.client.Logical().Write(path, options)
	if err != nil {
		return err
	}

	w.authorizer = secret.Auth.ClientToken

	return nil
}
