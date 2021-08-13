package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"monkey-admin/config"
	_ "monkey-admin/dao"
	"monkey-admin/pkg/middleware/logger"
	"monkey-admin/router"
)

var (
	port, mode string
)

func init() {
	flag.StringVar(&port, "port", "3000", "server listening on, default 3000")
	flag.StringVar(&mode, "mode", "debug", "server running mode, default debug mode")
}

func main() {
	port := config.GetServerCfg().Port
	flag.Parse()
	gin.SetMode(mode)
	r := router.Init()
	r.Use(logger.LoggerToFile())
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Start server: %+v", err)
	}
}
