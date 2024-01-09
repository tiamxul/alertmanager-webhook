package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/tiamxu/alertmanager-webhook/controllers"
	"github.com/tiamxu/alertmanager-webhook/log"
)

var cfg *Config

func init() {
	loadConfig()
	if err := cfg.Initial(); err != nil {
		log.Fatalf("配置文件错误,%v", err)
	}
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// gin.DisableConsoleColor()
	// f, _ := os.OpenFile("gin.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	// gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/dingding", controller.HandlerWebhook)
	r.POST("/feishu", controller.PrometheusAlert)

	srv := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Infof("address== %s", cfg.ListenAddress)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	log.Infoln("stop http server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed	: %s", err)
	}
	log.Infoln("successfully stopped")
}
