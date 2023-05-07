package main

import (
	"_core/auth"
	"_core/website"
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

	// CORS allow all
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin},
		AllowCredentials: true,
	}))

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
	websiteService := website.NewWebsiteService(conn, logger)
	routerContainer := router.NewRouterContainer(logger, e, authService, websiteService)

	routerContainer.CreateAuthRoutes()
	routerContainer.CreateWebsiteRoutes()

	e.Logger.Fatal(e.Start(":1323"))
}
