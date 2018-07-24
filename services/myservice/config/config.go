package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Server struct {
	Addr string
	Port int
}

type Database struct {
	Addr string
	Port int
	Name string
	User string
	Pass string
}

type Config struct {
	Srv Server   `mapstructure:"server"`
	DB  Database `mapstructure:"database"`
}

func Read(cmd *cobra.Command) (*Config, error) {
	viper.BindPFlag("server.addr", cmd.Flags().Lookup("listen"))
	viper.BindPFlag("server.port", cmd.Flags().Lookup("port"))

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := &Config{}

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
