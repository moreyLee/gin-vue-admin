package initialize

import "github.com/flipped-aurora/gin-vue-admin/agent/global"

func GetListenAddr() string {
	if global.ET_CONFIG.Agent.Listen == "" {
		return ":9999"
	}
	return ":" + global.ET_CONFIG.Agent.Listen
}
