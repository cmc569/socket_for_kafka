package setting

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configKafka struct {
	KafkaIp string		//kafka的IP位址
}

var ConfigKafka configKafka

func initKafkaConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		ConfigKafka = configKafka{"192.168.185.122:9092"}
	case gin.DebugMode:
		ConfigKafka = configKafka{os.Getenv("KAFKA_HOST")}
	case gin.TestMode:
		ConfigKafka = configKafka{"192.168.185.122:9092"}
	}
}