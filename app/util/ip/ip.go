package ip

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net"
)

func GetIPAddress(c echo.Context) string {
	userIP := ""
	//CF-Connecting-IP
	if len(c.Request().Header.Get("CF-Connecting-IP")) > 1 {
		userIP = c.Request().Header.Get("CF-Connecting-IP")
		fmt.Println(net.ParseIP(userIP))
	} else if len(c.Request().Header.Get("X-Forwarded-For")) > 1 {
		userIP = c.Request().Header.Get("X-Forwarded-For")
		fmt.Println(net.ParseIP(userIP))
	} else if len(c.Request().Header.Get("X-Real-IP")) > 1 {
		userIP = c.Request().Header.Get("X-Real-IP")
		fmt.Println(net.ParseIP(userIP))
	} else {
		userIP = extractIPAddress(c.Request().RemoteAddr)

	}
	return userIP
}

func GetUserAgent(c echo.Context) string {
	if len(c.Request().Header.Get("User-Agent")) > 1 {
		return c.Request().Header.Get("User-Agent")
	} else {
		return ""
	}
}
func extractIPAddress(ip string) string {
	if len(ip) > 0 {
		for i := len(ip); i >= 0; i-- {
			offset := len(ip)
			if (i + 1) <= len(ip) {
				offset = i + 1
			}
			if ip[i:offset] == ":" {
				return ip[:i]
			}
		}
	}
	return ip
}
