package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

type HttpConfig struct {
	ListenAddress string `yaml:"listen-address"`
	Mode          string `yaml:"mode"`
}
type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Address  string `yaml:"address"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type Config struct {
	HttpConfig  HttpConfig  `yaml:"http"`
	MysqlConfig MysqlConfig `yaml:"mysql"`
	LogConfig   LogConfig   `yaml:"log"`
}

func NewConfig(path string) *Config {
	conf := &Config{}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		panic(err)
	}

	return conf
}
