package utils

import (
	"net"
	"net/http"
	"strings"
)

// 获取IP地址
func GetIpAddress(request *http.Request, trustedProxyIPHeader []string) string {
	address := ""
	for _, proxyHeader := range trustedProxyIPHeader {
		header := request.Header.Get(proxyHeader)
		if len(header) > 0 {
			addresses := strings.Fields(header)
			if len(addresses) > 0 {
				address = strings.TrimRight(addresses[0], ",")
			}
		}
		if len(address) > 0 {
			return address
		}
	}
	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(request.RemoteAddr)
	}
	return address
}
