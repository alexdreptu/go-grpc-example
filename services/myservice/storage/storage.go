package storage

import (
	"fmt"

	"github.com/alexdreptu/go-grpc-example/services/myservice/config"
	"github.com/go-pg/pg"
)

type Connector interface {
	Close() error
}

type Conn struct {
	DB *pg.DB
}

func (c *Conn) Close() error {
	return c.DB.Close()
}

func New(conf *config.Config) (*Conn, error) {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.DB.Addr, conf.DB.Port),
		Database: conf.DB.Name,
		User:     conf.DB.User,
		Password: conf.DB.Pass,
	})

	return &Conn{db}, nil
}
