package main

import (
	"context"

	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics"
)

func Do() {
	client := metrics.NewFromConfig(configs.DefaultConfig())

	ctx := context.Background()
	client.GetMetrics(ctx)
}

func main() {

}
