package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/agent/config"
	"github.com/flipped-aurora/gin-vue-admin/agent/global"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")

	// 1️⃣ 当前工作目录（IDE / go run） 也可能指定到server 下的config.yml
	//v.AddConfigPath(".")

	// 2️⃣ agent 目录（源码运行）
	v.AddConfigPath("./agent")

	// 3️⃣ 可执行文件目录（生产环境）
	if exe, err := os.Executable(); err == nil {
		v.AddConfigPath(filepath.Dir(exe))
	}

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取 config.yml 失败: %w", err))
	}
	fmt.Println("[agent] using config:", v.ConfigFileUsed())
	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	global.ET_CONFIG = cfg
}
