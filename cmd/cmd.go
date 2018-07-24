package cmd

import (
	"github.com/alexdreptu/go-grpc-example/api"
	"github.com/alexdreptu/go-grpc-example/config"
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
		conf, err := config.Read(cmd)
		if err != nil {
			return err
		}

		return api.New(conf).Start()
	},
}

func init() {
	cobra.EnableCommandSorting = false

	serverStartCmd.Flags().StringP("listen", "l", "", "address to listen on")
	serverStartCmd.Flags().IntP("port", "p", 0, "port to listen on")

	serverCmd.AddCommand(serverStartCmd)
	RootCmd.AddCommand(serverCmd)
}
