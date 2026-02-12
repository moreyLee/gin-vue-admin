package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/agent/global"
	"github.com/flipped-aurora/gin-vue-admin/agent/middleware"
	"github.com/flipped-aurora/gin-vue-admin/agent/model/command"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(c *gin.Context) {
	// ✅ IP 白名单
	clientIP := c.ClientIP()
	if !middleware.IsIPAllowed(clientIP) {
		//global.ET_LOG
		fmt.Printf("[agent] reject ws ip=%s allow=%v path=%s", clientIP,
			global.ET_CONFIG.Agent.AllowIPs, c.Request.URL.Path,
		)
		c.AbortWithStatusJSON(403,
			gin.H{"error_message": "IP not allowed"})
		return
	}
	// ✅ token 校验
	token := c.GetHeader("X-Agent-Token")
	if token != global.ET_CONFIG.Agent.Token {
		c.AbortWithStatusJSON(401,
			gin.H{"error_message": "Token not allowed"})
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var req command.ExecCommand
		if err := json.Unmarshal(msg, &req); err != nil {
			continue
		}

		if req.Type != "exec" {
			continue
		}

		go handleExec(conn, req)
	}
}

func handleExec(conn *websocket.Conn, req command.ExecCommand) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(req.Timeout)*time.Second,
	)
	defer cancel()

	cmd := exec.CommandContext(ctx, req.Command, req.Args...)
	cmd.Dir = filepath.Clean(req.WorkDir)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Start()
	if err != nil {
		send(conn, command.ExecResult{
			RequestID: req.RequestID,
			Code:      1,
			Stderr:    err.Error(),
		})
		return
	}

	err = cmd.Wait()

	code := 0
	if err != nil {
		code = 1
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		code = 2
	}

	send(conn, command.ExecResult{
		RequestID: req.RequestID,
		Code:      code,
		Stdout:    stdoutBuf.String(),
		Stderr:    stderrBuf.String(),
	})
}
func send(conn *websocket.Conn, res command.ExecResult) {
	b, _ := json.Marshal(res)
	_ = conn.WriteMessage(websocket.TextMessage, b)
}
