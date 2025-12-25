package initialize

import "github.com/flipped-aurora/gin-vue-admin/agent/global"

func GetListenAddr() string {
	if global.Cfg.Agent.Listen == "" {
		return ":9999"
	}
	return ":" + global.Cfg.Agent.Listen
}
