package main

import (
	v1 "api/api/v1"
	_ "api/docs"
	"api/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "authorization", "Access-Control-Allow-Origin"},
	}))
	router.GET("/swagger-hello-word-xd/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	wsV1Group := router.Group("/ws/v1")
	wsV1Group.Use(service.WebSocketAuthMiddleware(false))
	Websocket(wsV1Group.Group("/"))
	return router
}
//======================================================================
//					not auth group
//======================================================================

func Websocket(router *gin.RouterGroup) {
	router.GET("/connect", v1.Connection)
}