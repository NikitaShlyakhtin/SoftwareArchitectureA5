package dependencies

import "github.com/labstack/echo/v4"

// IHandler defines the methods for HTTP handlers
type IHandler interface {
	CreateMessage() echo.HandlerFunc
	LikeMessage() echo.HandlerFunc
}
