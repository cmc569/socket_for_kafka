package main

import (
	"api/config/setting"
	"api/util/log"
	"api/util/websocket"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"runtime"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err == nil {
		gin.SetMode(gin.DebugMode)
		setting.InitConfig()
	}
	fmt.Println(gin.Mode())
	CreatedWebSocketPool()
	fmt.Println(setting.ServerConfig.Port)
	log.Error(Router().Run(":" + setting.ServerConfig.Port))
}

func CreatedWebSocketPool()  {
	pool := websocket.NewPool()
	go pool.Start()
}

func GoroutineStatus() {
	for {
		fmt.Println(runtime.NumGoroutine())
		time.Sleep(time.Second * 1)
	}
}
