package routes

import (
	"api.meet.the/components/config"
	"api.meet.the/controllers/legends"
	"api.meet.the/controllers/levels"
	"api.meet.the/controllers/question"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetupRoutes() {

	e := echo.New()

	e.Use(middleware.Logger())

	g := e.Group("v1/")

	setupGETRoutes(g)

	e.Logger.Fatal(e.Start(config.Data.Server.Port))

}

func setupGETRoutes(g *echo.Group) {

	g.GET("levels", levels.GetLevels)
	g.GET("question", question.GetQuestion)
	g.GET("legends", legends.GetLegends)

}
