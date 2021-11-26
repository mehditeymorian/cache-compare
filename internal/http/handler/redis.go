package handler

import "github.com/labstack/echo/v4"

type Redis struct {
}

func (r Redis) set(ctx echo.Context) error {
	panic("implement me!")
}

func (r Redis) get(ctx echo.Context) error {
	panic("implement me!")
}

func (r Redis) del(ctx echo.Context) error {
	panic("implement me!")
}

func (r Redis) Register(e *echo.Echo) {
	panic("implement me!")
}
