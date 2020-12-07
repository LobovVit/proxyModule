package config

import "github.com/tomazk/envcfg"

// declare a type that will hold your env variables
type Cfg struct {
	LOGLEVEL        string
	DB_PORT         int
	DB_HOST         string
	FAH_CONN_STRING string
	PORT            string
}

func InitConfig(conf *Cfg) error {
	envcfg.Unmarshal(conf)
	//TODO выставить значения по умолчанию
	conf.FAH_CONN_STRING = "apps/apps@eb-arp-dev-fah.otr.ru:1529/fah"
	conf.PORT = ":80"
	conf.LOGLEVEL = "DEBUG"
	//TODO Логика проверки переменных окружения
	return nil
}
