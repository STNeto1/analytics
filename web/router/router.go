package router

import (
	"_core/auth"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RouterContainer struct {
	logger    *zap.Logger
	engine    *echo.Echo
	validator *validator.Validate

	authService *auth.AuthService
}

func NewRouterContainer(logger *zap.Logger, engine *echo.Echo, authService *auth.AuthService) *RouterContainer {
	return &RouterContainer{
		engine:      engine,
		logger:      logger,
		authService: authService,
		validator:   validator.New(),
	}
}
