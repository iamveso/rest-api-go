package router

import (
	"company-api/controllers"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	r := echo.New()

	personAPI := r.Group("/person")
	peopleAPI := r.Group("/people")

	personAPI.POST("/", controllers.CreatePersonEndpoint)
	personAPI.GET("/:id", controllers.GetPersonEndpoint)

	peopleAPI.GET("/", controllers.GetPeopleEndpoint)

	return r
}
