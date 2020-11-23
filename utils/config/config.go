package config

import "github.com/tomazk/envcfg"

// declare a type that will hold your env variables
type Cfg struct {
	DEBUG   bool
	DB_PORT int
	DB_HOST string
}

func InitConfig(conf *Cfg) error {
	envcfg.Unmarshal(conf)
	//TODO Логика проверки переменных окружения
	return nil
}
