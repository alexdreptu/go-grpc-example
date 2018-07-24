package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-grpc-example",
	Short: "go-grpc-example command line interface",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server functionality",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "runs the api server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	serverCmd.AddCommand(serverStartCmd)
	RootCmd.AddCommand(serverCmd)
}
