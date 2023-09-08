module github.com/ralvescosta/emqx-sdk-go/examples/metrics

go 1.21

require (
	github.com/ralvescosta/dotenv v1.0.4
	github.com/ralvescosta/emqx-sdk-go v0.0.0-20230715012117-435122ddbd95
	github.com/ralvescosta/emqx-sdk-go/service/metrics v0.0.0-20230715012117-435122ddbd95
)

require (
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)

replace github.com/ralvescosta/emqx-sdk-go => ../../

replace github.com/ralvescosta/emqx-sdk-go/service/metrics => ../../service/metrics
