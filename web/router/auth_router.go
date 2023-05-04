package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (r *RouterContainer) CreateAuthRoutes() {
	group := r.engine.Group("/auth")

	group.POST("/login", func(c echo.Context) error {

		body := new(LoginRequestBody)
		if err := c.Bind(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := r.validator.Struct(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestFromError(err))
		}

		user, err := r.authService.Login(c.Request().Context(), body.Email, body.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		token, err := r.authService.GenerateToken(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		cookie := createCookieFromToken(token)
		c.SetCookie(cookie)

		return c.JSON(http.StatusCreated, echo.Map{
			"success": true,
		})
	})

	group.POST("/register", func(c echo.Context) error {
		body := new(RegisterRequestBody)

		if err := c.Bind(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := r.validator.Struct(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestFromError(err))
		}

		user, err := r.authService.Register(c.Request().Context(), body.Name, body.Email, body.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		token, err := r.authService.GenerateToken(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		cookie := createCookieFromToken(token)
		c.SetCookie(cookie)

		return c.JSON(http.StatusCreated, echo.Map{
			"success": true,
		})
	})
}

func createCookieFromToken(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "analytics.session",
		Domain:   "localhost",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}
}

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
