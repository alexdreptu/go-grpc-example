package server

import (
	"context"
	"fmt"
	"net"

	"github.com/alexdreptu/go-grpc-example/services/myservice/config"
	pb "github.com/alexdreptu/go-grpc-example/services/myservice/proto"
	"github.com/alexdreptu/go-grpc-example/services/myservice/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	Conf *config.Config
	Log  *zap.Logger
	Conn storage.Connector
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.Conf.Srv.Addr, s.Conf.Srv.Port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	pb.RegisterMyServiceServer(srv, s)

	return srv.Serve(l)
}

func (s *Server) AddData(ctx context.Context, req *pb.AddDataRequest) (*empty.Empty, error) {
	if err := s.Conn.AddData(req.GetServerIp(), req.GetClientIp(),
		req.GetMetadata(), req.GetMsg()); err != nil {
		return nil, err
	}

	s.Log.Info("added data",
		zap.String("server_ip", req.GetServerIp()),
		zap.String("client_ip", req.GetClientIp()),
		zap.Any("matadata", req.GetMetadata()),
		zap.String("message", req.GetMsg()))

	return &empty.Empty{}, nil
}

func New(conf *config.Config, conn storage.Connector) (*Server, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	s := &Server{
		Conf: conf,
		Log:  log,
		Conn: conn,
	}

	return s, nil
}
