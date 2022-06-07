package app

import (
	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/controller"
)

func NewApp() App {
	return &app{
		e: echo.New(),
	}
}

type App interface {
	ConfigMiddleware()
	ConfigLogFormat()
	ConfigLogLevel()
	ConfigErrHandler()

	RegisterHandlers()
	Parse()
	Run() error
}

type app struct {
	e             *echo.Echo
	LogLevel      string
	DbUrl         string
	Port          string
	Host          string
	ApiPrefix     string
	RedisLocation string
}

func (a *app) Run() error {
	address := a.Host + ":" + a.Port

	return a.e.Start(address)
}

func (a *app) RegisterHandlers() {
	controller.RegisterHandlers(a.e, a.ApiPrefix)
}
