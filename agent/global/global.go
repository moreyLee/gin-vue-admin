package global

import (
	"github.com/flipped-aurora/gin-vue-admin/agent/config"
	"gorm.io/gorm"
)

var (
	Cfg *config.Config
	DB  *gorm.DB
)
