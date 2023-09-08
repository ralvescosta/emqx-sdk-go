module github.com/ralvescosta/emqx-sdk-go/examples/alarms

go 1.21

require (
	github.com/ralvescosta/dotenv v1.0.4
	github.com/ralvescosta/emqx-sdk-go v0.0.0-20230717184733-e0be6f28f102
	github.com/ralvescosta/emqx-sdk-go/service/alarms v0.0.0-20230717184733-e0be6f28f102
)

require (
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)

replace github.com/ralvescosta/emqx-sdk-go => ../../

replace github.com/ralvescosta/emqx-sdk-go/service/alarms => ../../service/alarms
