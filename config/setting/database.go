package setting

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configDB struct {
	Type     string
	Host     string
	Port     string
	DBName   string
	UserName string
	Password string
}

type configMongoDB struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

type configRedis struct {
	Host     string
	Port     string
	DBName   string
	Password string
}

//Global Config Variable
var DataBaseConfig configDB
var MongoConfig configMongoDB
var RedisConfig configRedis

func databaseConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		DataBaseConfig = configDB{
			"mysql",
			"192.168.185.122",
			"3306",
			"otc",
			"otc",
			"742!!C0O0a!iyLdE^$5Tf^h7&"}

		MongoConfig = configMongoDB{
			"192.168.0.100",
			"27017",
			"",
			"",
			""}

		RedisConfig = configRedis{
			"192.168.0.100",
			"6379",
			"0",
			"rich-83719767"}

	case gin.DebugMode:
		DataBaseConfig = configDB{
			os.Getenv("DB_TYPE"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_POST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD")}

		MongoConfig = configMongoDB{
			os.Getenv("MongoDB_HOST"),
			os.Getenv("MongoDB_POST"),
			os.Getenv("MongoDB_NAME"),
			os.Getenv("MongoDB_USER"),
			os.Getenv("MongoDB_PASSWORD")}

		RedisConfig = configRedis{
			os.Getenv("Redis_HOST"),
			os.Getenv("Redis_POST"),
			os.Getenv("Redis_NAME"),
			os.Getenv("Redis_PASSWORD")}

	case gin.TestMode:
		DataBaseConfig = configDB{
			"mysql",
			"172.105.207.240",
			"3306",
			"otc",
			"loudada",
			"Aad#sfj8&e2X$S9DFa"}

		MongoConfig = configMongoDB{
			"192.168.0.100",
			"27017",
			"",
			"",
			""}

		RedisConfig = configRedis{
			"192.168.0.100",
			"6379",
			"0",
			"rich-83719767"}
	}
}
