package app

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
	localMiddleware "github.com/lstrgiang/gitscan/internal/infra/middleware"
)

const (
	DefaultDBUrl     = "postgres://postgres:postgres@localhost:5432/gitscan?sslmode=disable"
	DefaultPort      = "8080"
	DefaultHost      = "localhost"
	DefaultLogLevel  = "error"
	DefaultApiPrefix = "/api"

	DBUrlEnv     = "DB_URL"
	PortEnv      = "SERVER_PORT"
	HostEnv      = "SERVER_HOST"
	LogLevelEnv  = "LOG_LEVEL"
	ApiPrefixEnv = "API_PREFIX"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

func GetEnv(name string, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
}

func (a *app) Parse() {
	// parse flag
	flag.StringVar(&a.DbUrl, "db", GetEnv(DBUrlEnv, DefaultDBUrl), "DB Url for DB connection")
	flag.StringVar(&a.LogLevel, "log", GetEnv(LogLevelEnv, DefaultLogLevel), "Log level")
	flag.StringVar(&a.Port, "port", GetEnv(PortEnv, DefaultPort), "Running port")
	flag.StringVar(&a.Host, "host", GetEnv(HostEnv, DefaultHost), "Running host")
	flag.StringVar(&a.ApiPrefix, "api", GetEnv(ApiPrefixEnv, DefaultApiPrefix), "Server API Prefix")
	flag.Parse()
}

func (a *app) ConfigErrHandler() {
	defaultHandler := a.e.HTTPErrorHandler
	a.e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.IsWebSocket() {
			return // connection is hijacked, can't write to response anymore
		}
		switch err.(type) {
		case errors.SystemError:
			errObj, _ := err.(errors.SystemError)
			_ = c.JSON(int(errObj), NewErrorResponse(err))
		case errors.CustomError, errors.CustomParamError:
			_ = c.JSON(http.StatusBadRequest, NewErrorResponse(err))

		default:
			defaultHandler(err, c)
		}
	}
}

//config the logging method level and format

// configLogLevel changes echo logger to the given level in server options.
// If the log level is debug, sqlboiler debug mode will be enabled.
func (a *app) ConfigLogLevel() {
	logLevel := strings.ToLower(a.LogLevel)
	switch logLevel {
	case "error":
		a.e.Logger.SetLevel(log.ERROR)
	case "info":
		a.e.Logger.SetLevel(log.INFO)
	case "warn":
		a.e.Logger.SetLevel(log.WARN)
	case "off":
		a.e.Logger.SetLevel(log.OFF)
	case "debug":
		fallthrough
	default:
		a.e.Logger.SetLevel(log.DEBUG)
	}
}

// configLogHeader change echo global log format, and adhoc log prefix
// for more readable. The default one produces log in JSON format, with is
// intended to be collected by other tools, but we're not using such tools yet.
func (a *app) ConfigLogFormat() {
	a.e.HideBanner = true
	a.e.Logger.SetOutput(os.Stderr)

	// make echo context log more readable.
	if l, ok := a.e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level} ${short_file}:${line}")
	}

	// make echo request/response log (once per request) more readable
	a.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} requestID=${id} remote_ip=${remote_ip} ` +
			`${method} ${path} status=${status} err=${error} ` +
			`latency=${latency_human} user_agent=${user_agent}`,
	}))
}

func (a *app) ConfigMiddleware() {
	// use recover middleware
	a.e.Use(middleware.Recover())

	// use dependencies injection middleware
	a.e.Use(localMiddleware.Inject(a.DbUrl))
}
