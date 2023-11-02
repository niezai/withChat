package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

type server struct {
	IP   string `yaml:"IP"`
	Port string `yaml:"Port"`
}

type config struct {
	Server server
}

var Config config

func Init() {
	bytes, err := os.ReadFile("./configs/configs.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, &Config)
	if err != nil {
		panic(err)
	}
}
