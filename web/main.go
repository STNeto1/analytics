package main

import (
	"_core/auth"
	"_web/pkg"
	"_web/router"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conn := pkg.InitDB(logger)

	e := echo.New()

	pkg.Logger(e, logger)
	e.Use(middleware.Recover())

	authService := auth.NewAuthService(conn, os.Getenv("SECRET"), logger)
	routerContainer := router.NewRouterContainer(logger, e, authService)

	routerContainer.CreateAuthRoutes()

	e.Logger.Fatal(e.Start(":1323"))
}
