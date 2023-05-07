package router

import (
	"_core/auth"
	"_core/website"
	"_schemas/ent"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RouterContainer struct {
	logger    *zap.Logger
	engine    *echo.Echo
	validator *validator.Validate

	authService    *auth.AuthService
	websiteService *website.WebsiteService
}

func NewRouterContainer(logger *zap.Logger, engine *echo.Echo, authService *auth.AuthService, websiteService *website.WebsiteService) *RouterContainer {
	return &RouterContainer{
		engine:         engine,
		logger:         logger,
		authService:    authService,
		websiteService: websiteService,
		validator:      validator.New(),
	}
}

func (r *RouterContainer) UserFromContext(ctx echo.Context) (*ent.User, error) {
	user, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("invalid token")
	}

	claims, ok := user.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	validUuid, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	u, err := r.authService.GetUserFromId(ctx.Request().Context(), validUuid)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return u, nil

}
