package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/db"
	"github.com/lstrgiang/gitscan/internal/infra/inject"
	"github.com/lstrgiang/gitscan/internal/infra/redis"
)

func Inject(dbUrl string, redisLocation string) echo.MiddlewareFunc {
	// initialize database connection
	db, err := db.NewDB(dbUrl)
	if err != nil {
		fmt.Println(err)
		panic("Invalid database connection")
	}

	asynqClient := redis.NewAsynqClient(redisLocation)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			// make the DB available in echo Context for next handlers

			inject.SetAsynqClient(e, asynqClient)
			inject.SetDB(e, db)
			inject.SetDbUrl(e, dbUrl)

			return next(e)
		}
	}
}
