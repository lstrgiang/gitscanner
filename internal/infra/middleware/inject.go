package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/db"
	"github.com/lstrgiang/gitscan/internal/infra/inject"
)

func Inject(dbUrl string) echo.MiddlewareFunc {
	fmt.Println(dbUrl)
	db, err := db.NewDB(dbUrl)

	if err != nil {
		fmt.Println(err)
		panic("Invalid database connection")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			// make the DB available in echo Context for next handlers
			inject.SetDB(e, db)

			return next(e)
		}
	}
}
