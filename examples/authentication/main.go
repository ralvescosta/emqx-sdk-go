package main

import (
	"context"
	"fmt"

	"github.com/ralvescosta/dotenv"
	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/service/authentication"
)

func MoveAuthenticator() {
	client := authentication.NewFromConfig(configs.DefaultConfig())

	ctx := context.Background()

	err := client.PutMoveAuthenticator(ctx, "password_based", "before:password_based:built_in_database")
	if err != nil {
		panic(err)
	}

	fmt.Print("PutAuthenticationPosition")
}

func main() {
	dotenv.Configure(".env")

	MoveAuthenticator()
}
