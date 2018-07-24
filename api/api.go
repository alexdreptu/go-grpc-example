package api

import (
	"fmt"
	"net/http"

	"github.com/alexdreptu/go-grpc-example/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

func New(conf *config.Config) *API {
	a := &API{
		echo: echo.New(),
		Conf: conf,
	}

	a.echo.HideBanner = true

	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())

	a.echo.GET("/", a.Home)

	return a
}
