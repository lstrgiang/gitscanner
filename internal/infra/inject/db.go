package inject

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

const (
	keyDB = "inject_db"
)

func DB(e echo.Context) *sql.DB {
	val := e.Get(keyDB)
	if val == nil {
		e.Logger().Panic("nil db")
	}

	return val.(*sql.DB)
}

func SetDB(e echo.Context, db *sql.DB) {
	e.Set(keyDB, db)
}
