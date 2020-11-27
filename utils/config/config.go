package config

import "github.com/tomazk/envcfg"

// declare a type that will hold your env variables
type Cfg struct {
	DEBUG           bool
	DB_PORT         int
	DB_HOST         string
	FAH_CONN_STRING string
}

func InitConfig(conf *Cfg) error {
	envcfg.Unmarshal(conf)
	conf.FAH_CONN_STRING = "apps/apps@eb-arp-dev-fah.otr.ru:1529/fah"
	//TODO Логика проверки переменных окружения
	return nil
}
