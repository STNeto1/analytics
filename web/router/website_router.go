package router

import (
	"_core/website"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (r *RouterContainer) CreateWebsiteRoutes() {
	group := r.engine.Group("/website")

	group.GET("/list", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		websites, err := r.websiteService.GetUserWebsites(c.Request().Context(), u)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, websites)
	})

	group.GET("/show/:id", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		id := c.Param("id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: "invalid id",
			})
		}

		website, err := r.websiteService.GetWebsiteByID(c.Request().Context(), u, parsedId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, NotFoundResponse{
				Message: "website not found",
			})
		}

		return c.JSON(http.StatusOK, website)
	})

	group.POST("/create", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		body := new(CreateWebsiteBody)
		if err := c.Bind(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}
		if err := r.validator.Struct(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestFromError(err))
		}

		website, err := r.websiteService.CreateWebsite(c.Request().Context(), u, body.Name, body.Domain)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, website)
	})

	group.PUT("/update/:id", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		id := c.Param("id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: "invalid id",
			})
		}

		body := new(UpdateWebsiteBody)
		if err := c.Bind(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}
		if err := r.validator.Struct(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestFromError(err))
		}

		_, err = r.websiteService.UpdateWebsite(c.Request().Context(), u, parsedId, body.Name, body.Domain)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusNoContent, nil)
	})

	group.DELETE("/delete/:id", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		id := c.Param("id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: "invalid id",
			})
		}

		err = r.websiteService.DeleteWebsite(c.Request().Context(), u, parsedId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusNoContent, nil)
	})

	group.POST("/create-event/:id", func(c echo.Context) error {
		id := c.Param("id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: "invalid id",
			})
		}

		body := new(CreateWebsiteEventBody)
		if err := c.Bind(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}
		if err := r.validator.Struct(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestFromError(err))
		}

		err = r.websiteService.CreateWebsiteEvent(c.Request().Context(), parsedId, &website.CreateWebsiteEventPayload{
			EventName:      body.EventName,
			UrlPath:        body.UrlPath,
			UrlQuery:       body.UrlQuery,
			ReferrerPath:   body.ReferrerPath,
			ReferrerQuery:  body.ReferrerQuery,
			ReferrerDomain: body.ReferrerDomain,
			PageTitle:      body.PageTitle,
			PageData:       body.PageData,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusNoContent, nil)
	})

	group.GET("/events/search/:id/:name", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		id := c.Param("id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: "invalid id",
			})
		}

		name := c.Param("name")

		events, err := r.websiteService.GetWebsiteEventsByName(c.Request().Context(), u, parsedId, name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, events)
	})

	group.GET("/events/histogram/:id", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		id := c.Param("id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: "invalid id",
			})
		}

		level := c.QueryParam("level")

		events, err := r.websiteService.GetWebsiteEventHistogram(c.Request().Context(), u, parsedId, level)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, events)
	})

	group.GET("/events/:id", func(c echo.Context) error {
		u, err := r.UserFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: err.Error(),
			})
		}

		id := c.Param("id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, BadRequestResponse{
				Message: "invalid id",
			})
		}

		events, err := r.websiteService.GetWebsiteEvents(c.Request().Context(), u, parsedId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, InternalServerErrorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, events)
	})
}

type CreateWebsiteBody struct {
	Name   string `json:"name" validate:"required"`
	Domain string `json:"domain" validate:"required"`
}

type UpdateWebsiteBody struct {
	Name   string `json:"name" validate:"required"`
	Domain string `json:"domain" validate:"required"`
}

type CreateWebsiteEventBody struct {
	EventName      string  `json:"event_name" validate:"required"`
	UrlPath        *string `json:"url_path"`
	UrlQuery       *string `json:"url_query"`
	ReferrerPath   *string `json:"referrer_path"`
	ReferrerQuery  *string `json:"referrer_query"`
	ReferrerDomain *string `json:"referrer_domain"`
	PageTitle      *string `json:"page_title"`
	PageData       *string `json:"page_data"`
}
