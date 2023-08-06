module github.com/ralvescosta/emqx-sdk-go/examples/authentication

go 1.20

require (
	github.com/ralvescosta/dotenv v1.0.4
	github.com/ralvescosta/emqx-sdk-go v0.0.0-20230805105905-f125aba8a22c
	github.com/ralvescosta/emqx-sdk-go/service/authentication v0.0.0-20230805105905-f125aba8a22c
)

require (
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)

replace github.com/ralvescosta/emqx-sdk-go/service/authentication => ../../service/authentication
