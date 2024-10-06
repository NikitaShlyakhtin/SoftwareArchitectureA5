package app

import (
	"errors"
	"github.com/labstack/echo/v4"
)

func (app *Application) LikeMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return errors.New("not implemented")
	}
}
