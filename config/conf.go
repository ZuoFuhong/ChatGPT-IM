package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Conf struct {
	Server Server
	Mysql  Mysql
}

type Server struct {
	Addr string
	Port int
}

type Mysql struct {
	Addr     string
	Username string
	Password string
}

func LoadConf() *Conf {
	env := getMode()
	viper.SetConfigName("application-" + env)
	viper.AddConfigPath("conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	C := &Conf{}
	err = viper.Unmarshal(C)
	if err != nil {
		panic(env)
	}
	return C
}

func getMode() string {
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}
