package main

import (
	"flag"
	"fondo-go/config"
	"fondo-go/router"
	"github.com/gin-gonic/gin"
)

var (
	configPath = flag.String("configuration", "application.yml", "config file path")
)

func main() {
	flag.Parse()
	//casbin.InitCasbin()
	config.InitConfig(*configPath)
	gin.SetMode(config.Conf.EnvMode)
	router.Run(config.Conf.Port)
}
