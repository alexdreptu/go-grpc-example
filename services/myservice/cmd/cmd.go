package cmd

import (
	"github.com/alexdreptu/go-grpc-example/services/myservice/config"
	"github.com/alexdreptu/go-grpc-example/services/myservice/storage"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "myservice",
	Short: "myservice command line interface",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server functionality",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "runs the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := config.Read(cmd)
		if err != nil {
			return err
		}

		spew.Dump(conf) // for debugging purposes

		conn, err := storage.New(conf)
		if err != nil {
			return err
		}
		defer conn.Close()

		return nil
	},
}

func init() {
	cobra.EnableCommandSorting = false

	serverStartCmd.Flags().StringP("listen", "l", "", "address to listen on")
	serverStartCmd.Flags().IntP("port", "p", 0, "port to listen on")

	serverCmd.AddCommand(serverStartCmd)
	RootCmd.AddCommand(serverCmd)
}
