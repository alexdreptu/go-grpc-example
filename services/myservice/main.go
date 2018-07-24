package main

import (
	"os"

	"github.com/alexdreptu/go-grpc-example/services/myservice/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
