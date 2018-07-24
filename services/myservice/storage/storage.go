package storage

import (
	"fmt"

	"github.com/alexdreptu/go-grpc-example/services/myservice/config"
	"github.com/alexdreptu/go-grpc-example/services/myservice/models"
	"github.com/go-pg/pg"
)

type Connector interface {
	Close() error
	AddData(serverIP, clientIP string, metadata map[string]string, msg string) error
}

type Conn struct {
	DB *pg.DB
}

func (c *Conn) Close() error {
	return c.DB.Close()
}

func (c *Conn) AddData(serverIP, clientIP string, metadata map[string]string, msg string) error {
	data := &models.Data{
		ServerIP: serverIP,
		ClientIP: clientIP,
		Metadata: metadata,
		Msg:      msg,
	}

	return c.DB.Insert(data)
}

func New(conf *config.Config) (*Conn, error) {
	opt := &pg.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.DB.Addr, conf.DB.Port),
		Database: conf.DB.Name,
		User:     conf.DB.User,
		Password: conf.DB.Pass,
	}

	db := pg.Connect(opt)

	return &Conn{db}, nil
}
