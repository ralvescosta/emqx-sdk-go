package main

import (
	"testing"

	"github.com/ralvescosta/dotenv"
)

func TestX(t *testing.T) {
	dotenv.Configure(".env")
	GetAlarms()
}
