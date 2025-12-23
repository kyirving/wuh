package main

import (
	"basketball/internal/component/db"
	"basketball/internal/conf"
	"basketball/internal/dao"
	"basketball/internal/model"
	"basketball/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/pflag"
)

// main 管理端入口：加载配置、初始化数据库、自动迁移模型并启动HTTP服务
func main() {
	CommandArgs := &dao.CommandArgs{}
	pflag.StringVar(&CommandArgs.HomeDir, "homeDir", "/Users/wuh/work/code/wuh/basketball", "Home directory")
	pflag.StringVar(&CommandArgs.ConfigFile, "configFile", "config.yaml", "Config file")
	pflag.Parse()

	config := conf.InitConfig(CommandArgs)

	db := db.GetDb(&config.Db)
	// 自动迁移 Admin 模型表结构

	if config.Manager.Env == "devlopment" {
		// 开发环境 打印SQL语句
		db = db.Debug()

		if err := db.AutoMigrate(&model.User{}); err != nil {
			log.Fatalf("auto migrate failed: %v", err)
		}
	}

	r := router.InitRouter(db, config)
	// r.Use(middleware.JWT()) //

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Manager.Port),
		Handler: r.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

		fmt.Println("Server is running on port 8081")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
