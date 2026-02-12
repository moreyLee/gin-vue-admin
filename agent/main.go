package main

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/agent/global"
	"github.com/flipped-aurora/gin-vue-admin/agent/initialize"
	"github.com/flipped-aurora/gin-vue-admin/agent/router"
)

func main() {
	// 1️⃣ 配置
	initialize.InitConfig()
	// 3️⃣ 路由
	r := router.SetupRouter()
	fmt.Println("Agent listen on :", initialize.GetListenAddr())
	fmt.Printf("[agent] config allow_ips=%v\n", global.ET_CONFIG.Agent.AllowIPs)
	r.Run(initialize.GetListenAddr())
}
