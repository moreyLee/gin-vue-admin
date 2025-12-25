package api

import (
	"github.com/flipped-aurora/gin-vue-admin/agent/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StopReq struct {
	Name string `json:"name"` // A1 A6  T1
}

func StopEsSync(c *gin.Context) {
	var req StopReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	mq, err := service.GetSyncMQByName(req.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}

	if err := service.StopSyncMQ(mq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "stop failed",
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "stopped",
		"name": req.Name,
	})
}
