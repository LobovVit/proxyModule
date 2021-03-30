package config

import (
	"fmt"
	"github.com/tomazk/envcfg"
	"log"
)

// declare a type that will hold your env variables
type Cfg struct {
	LOGLEVEL        string
	TARANTOOL_PORT  string
	TARANTOOL_HOST  string
	FAH_CONN_STRING string
	PORT            string
	HOST            string
}

func (conf *Cfg) InitConfig() error {
	err := envcfg.Unmarshal(conf)
	if err != nil {
		log.Panic("не удалось прочитать переменные окружения:", err)
	}

	//TODO выставить значения по умолчанию
	//TODO Логика проверки переменных окружения
	if conf.FAH_CONN_STRING == "" {
		conf.FAH_CONN_STRING = "apps/apps@eb-arp-dev-fah.otr.ru:1529/fah"
	}
	if conf.PORT == "" {
		conf.PORT = ":8080"
	}
	if conf.LOGLEVEL == "" {
		conf.LOGLEVEL = "DEBUG"
	}
	if conf.TARANTOOL_PORT == "" {
		conf.TARANTOOL_PORT = ":3333"
	}
	if conf.TARANTOOL_HOST == "" {
		conf.TARANTOOL_HOST = "TARANTOOL_HOST"
	}
	//Вывод конфигурации
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("-----------конфигурация------------\n")
	fmt.Printf("   LOGLEVEL : %v\n", conf.LOGLEVEL)
	fmt.Printf("   HOST:PORT : %v\n", conf.HOST+conf.PORT)
	fmt.Printf("   FAH_CONN_STRING : %v\n", conf.FAH_CONN_STRING)
	fmt.Printf("   TARANTOOL : %v\n", conf.TARANTOOL_HOST+conf.TARANTOOL_PORT)
	fmt.Printf("-----------------------------------\n")

	return nil
}

func (conf *Cfg) ReInitConfig() {
	err := envcfg.Unmarshal(conf)
	if err != nil {
		log.Panic("не удалось прочитать переменные окружения:", err)
	}

	//TODO выставить значения по умолчанию
	//TODO Логика проверки переменных окружения
	conf.FAH_CONN_STRING = "apps/apps@eb-arp-dev-fah.otr.ru:1529/fah"
	if conf.PORT == "" {
		conf.PORT = ":80"
	}
	if conf.LOGLEVEL == "" {
		conf.LOGLEVEL = "DEBUG"
	}
	if conf.TARANTOOL_PORT == "" {
		conf.TARANTOOL_PORT = ":3333"
	}
	if conf.TARANTOOL_HOST == "" {
		conf.TARANTOOL_HOST = "TARANTOOL_HOST"
	}
	//Вывод конфигурации
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("-----------конфигурация------------\n")
	fmt.Printf("   LOGLEVEL : %v\n", conf.LOGLEVEL)
	fmt.Printf("   HOST:PORT : %v\n", conf.HOST+conf.PORT)
	fmt.Printf("   FAH_CONN_STRING : %v\n", conf.FAH_CONN_STRING)
	fmt.Printf("   TARANTOOL : %v\n", conf.TARANTOOL_HOST+conf.TARANTOOL_PORT)
	fmt.Printf("-----------------------------------\n")
}
