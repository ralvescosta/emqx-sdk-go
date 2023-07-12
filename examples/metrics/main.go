package examples

import (
	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics"
)

func main() {
	client := metrics.NewFromConfig(configs.DefaultConfig())
	client.GetMetrics()
}
