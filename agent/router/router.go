package router

import (
	"github.com/flipped-aurora/gin-vue-admin/agent/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	agentAPI := r.Group("/api")
	{
		agentAPI.POST("/esSyncMq/stop", api.StopEsSync)
	}

	return r
}
