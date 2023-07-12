package main

import (
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	os.Setenv("EMQX_HOST", "http://localhost")
	os.Setenv("EMQX_API_PORT", "18083")
	os.Setenv("EMQX_API_KEY", "1cae2222c1a22db7")
	os.Setenv("EMQX_SECRET_KEY", "PmTNJTnxcqyQFEVx4Ed9B9CZP5wn8ToEO79BAKWvFw0VeM")

	Do()
}
