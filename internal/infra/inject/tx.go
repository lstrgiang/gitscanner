package inject

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func Ctx(_ echo.Context) context.Context {
	return context.Background()
}

func DBCtx(e echo.Context) (*sql.DB, context.Context) {
	return DB(e), Ctx(e)
}

func Tx(e echo.Context, ctx context.Context) *sql.Tx {
	db := DB(e)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}
	return tx
}

func CtxTx(e echo.Context) (context.Context, *sql.Tx) {
	ctx := Ctx(e)
	return ctx, Tx(e, ctx)
}
