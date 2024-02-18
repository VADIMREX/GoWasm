package main

import (
	"fmt"
	"go-wasm-srv/config"
	"go-wasm-srv/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.ReadConfig("config.json")
	h := handler.NewHandler(cfg)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(h.ServeFile)

	if err := r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
