package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/widcha/go-project-x/internal/app"
	"github.com/widcha/go-project-x/internal/app/delivery/rest/handlers/healthcheck"
	"github.com/widcha/go-project-x/internal/pkg"
)

type Router struct {
	router     gin.IRouter
	datasource *pkg.DataSource
	container  *app.Container
}

func NewRouter(router gin.IRouter, datasource *pkg.DataSource, container *app.Container) *Router {
	return &Router{
		router:     router,
		datasource: datasource,
		container:  container,
	}
}

func (h *Router) RegisterRouter() {
	h.router.GET("/health", healthcheck.HealthCheckHandler(h.container.HealthCheckInport))
}
