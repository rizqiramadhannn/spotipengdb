package util

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gitlab.com/stefarf/iferr"
	"gopkg.in/yaml.v2"

	"spotipeng/app/global"
	"spotipeng/app/util/ip"
)

func LoggerI(c echo.Context, args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"request_id": c.Get("request_id"),
		"ip":         ip.GetIPAddress(c),
	}).Info(args)
}

func MustReadYaml(filename string) {
	b, err := os.ReadFile(filename)
	iferr.Panic(err)
	iferr.Panic(yaml.Unmarshal(b, &global.Config))
}
