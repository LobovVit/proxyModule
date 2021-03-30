package main

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"proxymodule/server"
	"proxymodule/utils/config"
	"proxymodule/utils/logger"
	"proxymodule/utils/oracle"
	"syscall"
)

var Config config.Cfg
var Log *zap.Logger
var FahConnect *sql.DB

func init() {
	//Читаем переменные окружения
	err := Config.InitConfig()
	if err != nil {
		log.Fatalf("Не удалось прочитать конфигурацию: %v", err)
	}
	//запускаем логер
	Log, err = logger.LogInit(&Config)
	if err != nil {
		log.Fatalf("Не удалось запустить ZAP LOGGER: %v", err)
	}
	Log.Info("init - ОК")
}
func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	Log.Info("PLAY")
	_ /*FahConnect*/, err := oracle.InitConn(Log, &Config)
	if err != nil {
		Log.Info("Не удалось подключиться к FAH: %v", zap.Error(err))
	}
	go server.Start(Log, &Config)
	Log.Info("GRPC")

	//выключаемся os.Signal = 1
	<-done
	server.Shutdown()
	Log.Info("!!!!!!!!Server Stopped!!!!!!!!")
}
