package main

import (
	"go.uber.org/zap"
	"log"
	"proxymodule/server"
	"proxymodule/utils/config"
	"proxymodule/utils/logger"
	"proxymodule/utils/oracle"
)

var Config config.Cfg
var Log *zap.Logger

func init() {
	//Читаем переменные окружения
	err := config.InitConfig(&Config)
	if err != nil {
		log.Fatalf("Не удалось прочитать переменные окружения: %v", err)
	}
	//запускаем логер
	Log, err = logger.LogInit(Config)
	if err != nil {
		log.Fatalf("Не удалось запустить ZAP LOGGER: %v", err)
	}
}
func main() {

	Log.Info("Play")
	//Подключаемся к БД
	oracle.InitConn(Log, Config.FAH_CONN_STRING)
	server.Start(Log, Config.PORT)
	server.GrpcStart(Log, Config)
}
