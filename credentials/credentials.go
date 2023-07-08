package credentials

import (
	"context"
	"encoding/base64"
	"fmt"
)

type (
	Credentials struct {
		APIKey    string
		SecretKey string
	}

	CredentialsProvider interface {
		Retrieve(ctx context.Context) (*Credentials, error)
	}
)

func (c *Credentials) HasKeys() bool {
	return len(c.APIKey) > 0 && len(c.SecretKey) > 0
}

func (c *Credentials) Header() (string, error) {
	if !c.HasKeys() {
		return "", ErrCredentialsNotProvided
	}

	return fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", c.APIKey, c.SecretKey)))), nil
}
