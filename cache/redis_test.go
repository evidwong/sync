package cache

import (
	"testing"
	"time"

	"baisiyi.net/config"
)

func TestRedis(t *testing.T) {
	var err error
	conf, err := config.NewConfig("../config.yaml")
	if err != nil {
		panic(err)
	}
	r := NewRedis(conf)

	timeoutDuration := 1 * time.Second

	if err = r.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !r.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := r.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = r.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
