package vault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

var (
	DefaultConfig *api.Config
)

func Client(token string) (*api.Client, error) {
	if DefaultConfig == nil {
		return nil, fmt.Errorf("Config was nil")
	}
	client, err := api.NewClient(DefaultConfig)
	if err != nil {
		return nil, err
	}
	client.SetToken(token)
	return client, nil
}
