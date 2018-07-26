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
	GetData(serverIP, clientIP string, metadata map[string]string) (*models.Data, error)
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

func (c *Conn) GetData(serverIP, clientIP string, metadata map[string]string) (*models.Data, error) {
	// mock data
	data := &models.Data{
		ServerIP: "211.31.180.207",
		ClientIP: "105.230.131.7",
		Metadata: map[string]string{
			"mollitia": "possimus",
		},
		Msg: "Possimus qui corporis numquam minus eos.",
	}

	return data, nil
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
