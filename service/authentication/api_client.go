package authentication

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/credentials"
	"github.com/ralvescosta/emqx-sdk-go/endpoint"
	"github.com/ralvescosta/emqx-sdk-go/service/authentication/types"
)

type (
	Options struct {
		Logger      *zap.SugaredLogger
		Credentials credentials.CredentialsProvider
		Endpoint    endpoint.EndpointResolver
		HTTPClient  configs.HTTPClient
	}

	AuthenticationClient interface {
		DeleteUserInAuthenticator(ctx context.Context, authenticatorID, userID string) error
		GetUserFromAuthenticator(ctx context.Context, authenticatorID, userID string) (*types.ResponseUserFromAuthenticator, error)
		PostCreateUsersForAuthenticator(ctx context.Context, authenticatorID string, req *types.RequestCreateUserForAuthenticator) (*types.ResponseCreateUserForAuthenticator, error)
		PostGlobalImportUsersIntoAuthenticator(ctx context.Context, authenticatorID string, file *os.File) error
		PutMoveAuthenticator(ctx context.Context, authenticatorID, position string) error
		PutUpdateUserInAuthenticator(ctx context.Context, authenticatorID, userID string) error
	}

	authenticationClient struct {
		logger    *zap.SugaredLogger
		apiClient *client.APIClient
	}
)

func New(options *client.APIClientOptions) AuthenticationClient {
	return &authenticationClient{options.Logger, client.New(options)}
}

func NewFromConfig(cfg *configs.Config) AuthenticationClient {
	opts := &client.APIClientOptions{
		Logger:      cfg.Logger,
		Credentials: cfg.Credentials,
		Endpoint:    cfg.Endpoint,
		HTTPClient:  cfg.HTTPClient,
	}

	return New(opts)
}
