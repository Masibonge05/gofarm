package main

import (
	"github.com/Masibonge05/gofarm/constant"
	"github.com/Masibonge05/gofarm/router"
	"github.com/Masibonge05/gofarm/server"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()

	constant.LoadStatic(app)

	app.Renderer = constant.LoadTemplate()

	router.LoadAllRouters(app)

	server.SetServer(app)

}
