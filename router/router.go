package router

import (
	"fondo-go/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(addr string) {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"*", "PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"*", "x-requested-with", "Origin", "Content-Length", "Content-Type"}

	router.Use(cors.New(corsConfig))
	//router.Use(middleware.Auth(false))

	controller.UserController(router)

	router.Run(":" + addr)
}
