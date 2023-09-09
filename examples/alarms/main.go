package main

import (
	"context"
	"fmt"

	"github.com/ralvescosta/dotenv"
	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/service/alarms"
)

func GetAlarms() {
	client := alarms.NewFromConfig(configs.DefaultConfig())

	ctx := context.Background()

	res, err := client.GetListAlarms(ctx, 1, 50, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetAlarms: %v", res)
}

func main() {
	dotenv.Configure(".env")
	GetAlarms()
}
