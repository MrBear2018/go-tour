package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog-service/global"
	"github.com/go-tour/blog-service/internal/model"
	"github.com/go-tour/blog-service/internal/routers"
	"github.com/go-tour/blog-service/utils/config"
	"github.com/go-tour/blog-service/utils/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupConfigSetting()
	if err != nil {
		log.Fatalf("init.setupConfigSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	global.Logger.Info()
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewCusRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupConfigSetting() error {
	setting, err := config.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	dbEngine, err := model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.DBEngine = dbEngine
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags)

	return nil
}
