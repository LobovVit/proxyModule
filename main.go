package main

import (
	"go.uber.org/zap"
	"log"
	"proxymodule/utils/config"
	"proxymodule/utils/logger"
)

var Config config.Cfg
var Log *zap.Logger

func init() {
	//Читаем переменные окружения
	err := config.InitConfig(&Config)
	if err != nil {
		log.Fatalf("Не прочитать переменные окружения %v", err)
	}
	//запускаем логер
	Log, err = logger.LogInit()
	if err != nil {
		log.Fatalf("Не удалось запустить ZAP LOGGER %v", err)
	}
}
func main() {

	Log.Info("Play")
	Log.Info("Play", zap.Bool("DEBUG", Config.DEBUG))
	Log.Info("Play", zap.Int("DB_PORT", Config.DB_PORT))
	//Подключаемся к БД

	//Начинаем слушать входящие

}
