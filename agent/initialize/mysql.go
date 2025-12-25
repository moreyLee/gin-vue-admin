package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/agent/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMySQL() {
	cfg := global.Cfg.Mysql

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql connect error: %v", err)
	}

	global.DB = db
}
