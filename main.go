package main

import (
	"os"

	"github.com/alexdreptu/go-grpc-example/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
