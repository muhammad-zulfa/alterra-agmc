package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/muhammad-zulfa/alterra-agmc/controllers"
	"net/http"
)

func Init(bookController controllers.BookController) *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/v1/books", bookController.FindAll)
	e.GET("/v1/books/:id", bookController.FindById)
	e.POST("/v1/books", bookController.Create)
	e.PUT("/v1/books/:id", bookController.Update)
	e.DELETE("/v1/books/:id", bookController.Delete)

	return e
}
