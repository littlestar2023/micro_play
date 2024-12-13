package db

import (
	"github.com/littlestar2023/common_pkg/global"
	"go.uber.org/zap"
	"micro_play/pkg/model"
)

func InitDb() {

	var err error
	if err = migration(); err != nil {
		global.CMP_LOG.Info("migration database error", zap.Error(err))
	}
}

func migration() (err error) {
	return global.CMP_DB.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(
		&model.User{})
}
