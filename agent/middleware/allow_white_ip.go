package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/agent/global"
	"net"
	"strings"
)

func IsIPAllowed(ip string) bool {
	//if len(global.ET_CONFIG.Agent.AllowIPs) == 0 {
	//	return true // 未配置则放行（可选）
	//}
	for _, cidr := range global.ET_CONFIG.Agent.AllowIPs {
		if cidrMatch(ip, cidr) {
			return true
		}
	}
	return false
}

func cidrMatch(ipStr, cidr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	// 如果是单 IP
	if !strings.Contains(cidr, "/") {
		return ip.Equal(net.ParseIP(cidr))
	}
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	return subnet.Contains(ip)
}
