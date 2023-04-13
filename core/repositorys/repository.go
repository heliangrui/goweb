package repositorys

import (
	"go-web/core/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Orm    *gorm.DB
	Logger *zap.Logger
}

func (d *Repository) MakeOrm() {
	d.Orm = config.GDb
}
