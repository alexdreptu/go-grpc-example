package cmd

import (
	"github.com/alexdreptu/go-grpc-example/services/myservice/config"
	"github.com/alexdreptu/go-grpc-example/services/myservice/server"
	"github.com/alexdreptu/go-grpc-example/services/myservice/storage"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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

		conn, err := storage.New(conf)
		if err != nil {
			return err
		}
		defer conn.Close()

		s, err := server.New(conf, conn)
		if err != nil {
			return err
		}

		s.Log.Info("Starting server",
			zap.String("addr", conf.Srv.Addr),
			zap.Int("port", conf.Srv.Port),
			zap.String("dbhost", conf.DB.Addr),
			zap.Int("dbport", conf.DB.Port),
			zap.String("dbname", conf.DB.Name),
			zap.String("dbuser", conf.DB.User),
			zap.String("dbpass", conf.DB.Pass))
		defer s.Log.Sync()

		return s.Start()
	},
}

func init() {
	cobra.EnableCommandSorting = false

	serverStartCmd.Flags().StringP("listen", "l", "", "address to listen on")
	serverStartCmd.Flags().IntP("port", "p", 0, "port to listen on")

	serverCmd.AddCommand(serverStartCmd)
	RootCmd.AddCommand(serverCmd)
}
