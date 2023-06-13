package main

import (
	"github.com/ronyldo12/secrets/cmd"
)

// This should be in an env file in production
const MySecret string = "1234567891234567"

func main() {
	cmd.Execute()
}
