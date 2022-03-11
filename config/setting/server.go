package setting

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configServer struct {
	Port string
}

var ServerConfig configServer

func initServerConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		ServerConfig = configServer{"9233"}
	case gin.DebugMode:
		ServerConfig = configServer{os.Getenv("PORT")}
	case gin.TestMode:
		ServerConfig = configServer{"9233"}
	}
}