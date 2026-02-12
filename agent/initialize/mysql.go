package initialize

//func InitMySQL() {
//	cfg := global.ET_CONFIG.Mysql
//
//	dsn := fmt.Sprintf(
//		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
//		cfg.User,
//		cfg.Password,
//		cfg.Host,
//		cfg.Port,
//		cfg.Database,
//	)
//
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("mysql connect error: %v", err)
//	}
//
//	global.ET_CONFIG = db
//}
