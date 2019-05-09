package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config 配置项
type Config struct {
	Redis *Redis
	Mysql *Mysql
}

// Redis 配置
type Redis struct {
	Host string `yml:"host" json:"host"`
	Port string `yml:"port" json:"port"`
	Auth string `yml:"auth" json:"auth"`
	DB   string `yml:"db" json:"db"`
}

// Mysql 配置
type Mysql struct {
	Host     string `yml:"host" json:"host"`
	Port     string `yml:"port" json:"port"`
	User string `yml:"user" json:"user"`
	PassWord string `yml:"password" json:"password"`
	DataBase string `yml:"database" json:"database"`
	Charset  string `yml:"charset" json:"charset"`
}

// NewConfig 获取配置项
func NewConfig(file string) (*Config, error) {
	config := new(Config)
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
