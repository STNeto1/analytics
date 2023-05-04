package main

import (
	"_core/auth"
	"_web/pkg"
	"_web/router"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
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
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		TokenLookup: "cookie:analytics.session",
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/auth/login" || c.Path() == "/auth/register" {
				return true
			}
			return false
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, router.BadRequestResponse{
				Message: "Unauthorized",
			})
		},
	}))

	authService := auth.NewAuthService(conn, os.Getenv("SECRET"), logger)
	routerContainer := router.NewRouterContainer(logger, e, authService)

	routerContainer.CreateAuthRoutes()

	e.Logger.Fatal(e.Start(":1323"))
}
