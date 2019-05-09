package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := NewConfig("../config.yaml")
	if err != nil {
		fmt.Println("fail")
		return
	}
	if config.Mysql.Host != "" {
		// fmt.Println(config.Redis.Auth=="")
		// fmt.Println(config.Mysql.DataBase)
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
