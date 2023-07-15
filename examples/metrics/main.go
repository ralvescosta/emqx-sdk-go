package main

import (
	"context"
	"fmt"

	"github.com/ralvescosta/dotenv"
	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics"
)

func GetMetrics() {
	client := metrics.NewFromConfig(configs.DefaultConfig())

	ctx := context.Background()
	res, err := client.GetMetrics(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetMetrics: %v", res)
}

func GetMonitorCurrentNode() {
	client := metrics.NewFromConfig(configs.DefaultConfig())

	ctx := context.Background()
	res, err := client.GetMonitorCurrentNode(ctx, "emqx@node1.emqx.io")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nGetMonitorCurrentNode: %v", *res)
}

func main() {
	dotenv.Configure(".env")
	GetMetrics()
	GetMonitorCurrentNode()
}
