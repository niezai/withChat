package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Server struct {
	Ip   string
	Port string
}

type Log struct {
	Log_Path string
	Log_Name string
}

type config struct {
	Server Server `json:"server"`
	Log    Log    `json:"log"`
	Mysql  Mysql  `json:"mysql"`
	Redis  Redis  `json:"redis"`
}

type Mysql struct {
	Username string
	Password string
	Db       string
	Host     string
	Port     int
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Db       int
}

var Config config

func Init() {
	bytes, err := os.ReadFile("./config/config.yml")
	//bytes, err := os.ReadFile("../../config/config.yml") //测试时使用这个
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, &Config)
	if err != nil {
		panic(err)
	}

}
