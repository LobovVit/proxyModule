package main

import (
	"go.uber.org/zap"
	"log"
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
	Log, err = logger.LogInit()
	if err != nil {
		log.Fatalf("Не удалось запустить ZAP LOGGER: %v", err)
	}
}
func main() {

	Log.Info("Play")
	Log.Info("Play", zap.Bool("DEBUG", Config.DEBUG))
	Log.Info("Play", zap.Int("DB_PORT", Config.DB_PORT))
	//Подключаемся к БД
	db, err := oracle.TestCon()
	if err != nil {
		Log.Info("Не удалось запустить подключится к БД")
	}
	err = db.Ping()
	if err != nil {
		Log.Info("Не удалось проверить подключится к БД", zap.Any("err", err))
	}
	if err == nil {
		Log.Info("Подключились к БД")
	}
	//Начинаем слушать входящие

}
