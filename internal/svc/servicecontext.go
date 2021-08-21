package svc

import (
	"github.com/jacksonyoudi/datacenter/internal/config"
	"github.com/jacksonyoudi/datacenter/internal/middleware"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Usercheck  rest.Middleware
	Admincheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Usercheck:  middleware.NewUsercheckMiddleware().Handle,
		Admincheck: middleware.NewAdmincheckMiddleware().Handle,
	}
}
