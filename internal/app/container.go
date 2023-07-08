package app

import (
	"github.com/widcha/go-project-x/internal/app/usecase/healthcheck"
	"github.com/widcha/go-project-x/internal/pkg"
)

type Container struct {
	HealthCheckInport healthcheck.Inport
}

func NewContainer(datasource *pkg.DataSource) *Container {
	return &Container{
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),
	}
}
