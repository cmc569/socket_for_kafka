package setting

import (
	"github.com/gin-gonic/gin"
	"os"
)

// TODO：文章發布的圖片上傳路徑定義
// 圖片路徑組成  =  根目錄 + 文章目錄 + 類別目錄 + 文章ID目錄 + 圖片.jpg

type configPath struct {
	VerboseLogPath     string
	ErrorLogPath       string
}

var PathConfig configPath

func pathConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		PathConfig = configPath{"/opt/api/blockchain-api/log/verbose.log",
			"log/error.log"}
	case gin.DebugMode:
		PathConfig = configPath{os.Getenv("VERBOSE_LOG_PATH"),
			os.Getenv("ERROR_LOG_PATH")}
	case gin.TestMode:
		PathConfig = configPath{"/opt/api/blockchain-api/log/verbose.log",
			"log/error.log"}
	}
}
