//+build !test

package main

import (
	"api/config/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	setting.InitConfig()
}
