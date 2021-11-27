package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mehditeymorian/cache-compare/internal/cache"
	"github.com/mehditeymorian/cache-compare/internal/http/response"
	"net/http"
)

type CacheHandler struct {
	Client cache.Cache
}

func (c CacheHandler) set(ctx echo.Context) error {
	key := ctx.QueryParam("key")
	value := ctx.QueryParam("val")

	err := c.Client.Set(ctx.Request().Context(), key, value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c CacheHandler) get(ctx echo.Context) error {
	key := ctx.Param("key")

	get, err := c.Client.Get(ctx.Request().Context(), key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	valueResponse := response.NewValueResponse(get)

	return ctx.JSON(http.StatusOK, valueResponse)
}

func (c CacheHandler) del(ctx echo.Context) error {
	key := ctx.Param("key")

	err := c.Client.Del(ctx.Request().Context(), key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c CacheHandler) Register(prefix string, app *echo.Echo) {
	app.GET(prefix+"/:key", c.get)
	app.POST(prefix, c.set)
	app.DELETE(prefix+"/:key", c.del)
}
