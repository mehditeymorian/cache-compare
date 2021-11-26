package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mehditeymorian/cache-compare/internal/cache"
	"github.com/mehditeymorian/cache-compare/internal/http/response"
	"net/http"
)

type Tikv struct {
	Client cache.Cache
}

func (t Tikv) set(ctx echo.Context) error {
	key := ctx.QueryParam("key")
	value := ctx.QueryParam("val")

	err := t.Client.Set(ctx.Request().Context(), key, value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (t Tikv) get(ctx echo.Context) error {
	key := ctx.Param("key")

	get, err := t.Client.Get(ctx.Request().Context(), key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	valueResponse := response.NewValueResponse(get)

	return ctx.JSON(http.StatusOK, valueResponse)
}

func (t Tikv) del(ctx echo.Context) error {
	key := ctx.Param("key")

	err := t.Client.Del(ctx.Request().Context(), key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (t Tikv) Register(app *echo.Echo) {
	app.GET("/tikv/:key", t.get)
	app.POST("/tikv", t.set)
	app.DELETE("/tikv/:key", t.del)
}
