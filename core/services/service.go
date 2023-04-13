package services

import (
	"go-web/core/repositorys"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
}

func (e *Service) MakeRepository(c *repositorys.Repository) *Service {
	c.Logger = e.Logger
	c.MakeOrm()
	return e
}
func (e *Service) MakeService(c *Service) *Service {
	c.Logger = e.Logger
	return e
}
