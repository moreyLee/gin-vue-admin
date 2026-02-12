package router

import (
	"github.com/flipped-aurora/gin-vue-admin/agent/ws"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)
	//agentAPI := r.Group("/api")
	//{
	//	agentAPI.POST("/esSyncMq/stop", api.StopEsSync)
	//}
	agentWS := r.Group("/ws")
	{
		agentWS.GET("/activity-script", ws.HandleWS)
	}
	return r
}
