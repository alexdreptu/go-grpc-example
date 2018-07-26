package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alexdreptu/go-grpc-example/config"
	pb "github.com/alexdreptu/go-grpc-example/services/myservice/proto"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
)

type HTTPResponse struct {
	Status  string
	Code    int
	Message string
	Error   string
}

type API struct {
	echo *echo.Echo
	Conf *config.Config
}

func (a *API) Home(ctx echo.Context) error {
	code := http.StatusNotImplemented

	return ctx.JSON(code, &HTTPResponse{
		Status:  "error",
		Code:    code,
		Message: http.StatusText(code),
	})
}

func (a *API) Start() error {
	addr := fmt.Sprintf("%s:%d", a.Conf.Srv.Addr, a.Conf.Srv.Port)
	return a.echo.Start(addr)
}

func (a *API) AddData(ctx echo.Context) error {
	params := ctx.QueryParams()
	code := http.StatusInternalServerError

	conn, err := newGRPCConn(a.Conf.MyService.Addr, a.Conf.MyService.Port)
	if err != nil {
		return ctx.JSON(code, &HTTPResponse{
			Status:  "error",
			Code:    code,
			Message: http.StatusText(code),
			Error:   err.Error(),
		})
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	data := &pb.AddDataRequest{
		Data: &pb.Data{
			ServerIp: params.Get("server_ip"),
			ClientIp: params.Get("client_ip"),
			Metadata: map[string]string{
				params.Get("key"): params.Get("value"),
			},
			Msg: params.Get("msg"),
		},
	}

	_, err = c.AddData(context.Background(), data)

	if err != nil {
		return ctx.JSON(code, &HTTPResponse{
			Status:  "error",
			Code:    code,
			Message: http.StatusText(code),
			Error:   err.Error(),
		})
	}

	code = http.StatusOK

	return ctx.JSON(code, &HTTPResponse{
		Status:  "ok",
		Code:    code,
		Message: "data added",
	})
}

func (a *API) GetData(ctx echo.Context) error {
	params := ctx.QueryParams()
	code := http.StatusInternalServerError
	spew.Dump(params) // debugging

	conn, err := newGRPCConn(a.Conf.MyService.Addr, a.Conf.MyService.Port)
	if err != nil {
		return ctx.JSON(code, &HTTPResponse{
			Status:  "error",
			Code:    code,
			Message: http.StatusText(code),
			Error:   err.Error(),
		})
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	data := &pb.GetDataRequest{
		Data: &pb.Query{
			ServerIp: params.Get("server_ip"),
			ClientIp: params.Get("client_ip"),
			Metadata: map[string]string{
				params.Get("key"): params.Get("value"),
			},
		},
	}

	resp, err := c.GetData(context.Background(), data)
	if err != nil {
		return ctx.JSON(code, &HTTPResponse{
			Status:  "error",
			Code:    code,
			Message: http.StatusText(code),
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"server_ip": resp.GetData().GetServerIp(),
		"client_ip": resp.GetData().GetClientIp(),
		"metadata":  resp.GetData().GetMetadata(),
		"msg":       resp.GetData().GetMsg(),
	})
}

func New(conf *config.Config) *API {
	a := &API{
		echo: echo.New(),
		Conf: conf,
	}

	a.echo.HideBanner = true

	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())

	a.echo.GET("/", a.Home)
	a.echo.GET("/get", a.GetData)
	a.echo.POST("/add", a.AddData)

	return a
}

func newGRPCConn(addr string, port int) (*grpc.ClientConn, error) {
	s := fmt.Sprintf("%s:%d", addr, port)
	conn, err := grpc.Dial(s, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
