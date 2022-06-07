package inject

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

const (
	keyDB    = "inject_db"
	keyDBUrl = "inject_db_url"
)

func DB(e echo.Context) *sql.DB {
	val := e.Get(keyDB)
	if val == nil {
		e.Logger().Panic("nil db")
	}

	return val.(*sql.DB)
}

func SetDbUrl(e echo.Context, dbUrl string) {
	e.Set(keyDBUrl, dbUrl)
}

func SetDB(e echo.Context, db *sql.DB) {
	e.Set(keyDB, db)
}

func DBUrl(e echo.Context) string {
	val := e.Get(keyDBUrl)

	return val.(string)
}
