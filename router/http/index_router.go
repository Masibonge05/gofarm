package http

import (
	"github.com/labstack/echo/v4"
	"github.com/Masibonge05/gofarm/controller/context/pages"

)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext) 

}