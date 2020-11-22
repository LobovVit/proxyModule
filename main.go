package main

import (
	"proxymodule/utils/logger"
)

func main() {
	//Читаем переменные окружения

	//запускаем логер
	log := logger.LogInit()
	log.Info("Play")
	//Подключаемся к БД

	//Начинаем слушать входящие

}
