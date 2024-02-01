package routes

import (
	"app/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", handlers.HomeHandler)
	e.POST("/", handlers.ContactHandler)
	e.POST("/toast/remove", handlers.ToastRemover)
}
