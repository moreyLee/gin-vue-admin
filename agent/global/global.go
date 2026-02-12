package global

import (
	"github.com/flipped-aurora/gin-vue-admin/agent/config"
	"gorm.io/gorm"
)

var (
	ET_CONFIG config.Config
	ET_MYSQL  *gorm.DB
)
