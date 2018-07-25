package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexdreptu/go-grpc-example/config"
	pb "github.com/alexdreptu/go-grpc-example/services/myservice/proto"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
)

type API struct {
	echo *echo.Echo
	Conf *config.Config
}

func (a *API) Home(ctx echo.Context) error {
	code := http.StatusNotImplemented
	return echo.NewHTTPError(code, http.StatusText(code))
}

func (a *API) Start() error {
	addr := fmt.Sprintf("%s:%d", a.Conf.Srv.Addr, a.Conf.Srv.Port)
	return a.echo.Start(addr)
}

func (a *API) AddData(ctx echo.Context) error {
	params := ctx.QueryParams()

	conn, err := newGRPCConn(a.Conf.MyService.Addr, a.Conf.MyService.Port)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	data := &pb.AddDataRequest{
		ServerIp: params.Get("server_ip"),
		ClientIp: params.Get("client_ip"),
		Metadata: map[string]string{
			params.Get("key"): params.Get("value"),
		},
		Msg: params.Get("msg"),
	}

	_, err = c.AddData(context.Background(), data)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"code":    strconv.Itoa(http.StatusOK),
		"message": "data added",
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
