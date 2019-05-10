package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/silenceper/wechat/cache"

	"baisiyi.net/config"
	"baisiyi.net/util"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/template"
)

var (
	mydb        *Mydb
	pool        *cache.Redis
	startTime   int64
	endTime     int64
	redisConStr string
	mysqlConStr string
	logger      *log.Logger
)

const (
	configFile = "./config.yml"
)

// Mydb 数据连接
type Mydb struct {
	db *gorm.DB
}

// NewMydb 获取数据连接
func NewMydb(cfg *config.Config) *Mydb {
	mysqlConStr = cfg.Mysql.User + ":" + cfg.Mysql.PassWord + "@tcp(" + cfg.Mysql.Host + ":" + cfg.Mysql.Port + ")/" + cfg.Mysql.DataBase + "?charset=" + cfg.Mysql.Charset + "&parseTime=True&loc=Local"
	mydb, err := gorm.Open("mysql", mysqlConStr)
	if err != nil {
		fmt.Println("connect error: ", err)
		panic("数据库连接失败")
	}
	mydb.DB().SetMaxIdleConns(10)
	mydb.DB().SetMaxOpenConns(100)
	return &Mydb{mydb}
}

// WechatMessageBase 微信信息基础
type WechatMessageBase struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// WechatMessage redis存的微信推送信息
type WechatMessage struct {
	Cid   int    `json:"cid"`
	Type  string `json:"type"`
	Comno string `json:"comno"`
	Phone string `json:"phone"`
	Job   string `json:"job"`
}

//Config 微信配置
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	PayMchID       string //支付 - 商户 ID
	PayNotifyURL   string //支付 - 接受微信支付结果通知的接口地址
	PayKey         string //支付 - 商户后台设置的支付 key
	Cache          cache.Cache
}

// Wechat 微信
type Wechat struct {
	Context *context.Context
}

// NewWechat 实例化wechat
func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}
func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.PayMchID = cfg.PayMchID
	context.PayKey = cfg.PayKey
	context.PayNotifyURL = cfg.PayNotifyURL
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

/**
 * run
 * 开启时查询所有的商户，并加入到时服务中
 */
func run(mydb *Mydb) {
	logFileName := "./logs/" + time.Now().Format("2006-01-02") + ".log"
	logFileHandle, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Fatalln("fail to create " + logFileName + " file!")
	}
	defer logFileHandle.Close()
	logger = log.New(logFileHandle, "", log.LstdFlags|log.Lshortfile) // 日志文件格式:log包含时间及文件行数
	for i := 0; i < 1; i++ {
		go sendWechatTemplateMessage(mydb)
		go sendSmsMessage(mydb)
	}
}

func sendWechatTemplateMessage(mydb *Mydb) {
	var cid string
	for {
		r := pool.Conn.Get()
		key, err := redis.String(r.Do("Rpop", "wechat:message:template"))
		if err != nil {
			r.Close()
			continue
		}
		content, err := redis.String(r.Do("Get", key))
		if err != nil {
			r.Close()
			continue
		}
		var redisData map[string]string
		err = json.Unmarshal([]byte(content), &redisData)
		if err != nil {
			r.Close()
			continue
		}
		if _, ok := redisData["job"]; !ok {
			r.Close()
			continue
		}
		var msg template.Message
		err = json.Unmarshal([]byte(redisData["job"]), &msg)
		if err != nil {
			r.Close()
			continue
		}

		cid = redisData["cid"]
		_, err = r.Do("Select", 1)
		if err != nil {
			continue
		}
		wechatConfig, err := redis.StringMap(r.Do("HgetAll", "wechat_config:"+cid))
		if err != nil {
			continue
		}
		wechat := NewWechat(&Config{
			AppID:     wechatConfig["appid"],
			AppSecret: wechatConfig["appsecret"],
			Token:     wechatConfig["token"],
			Cache:     pool,
		})
		tplMsg := template.NewTemplate(wechat.Context)
		_, err = tplMsg.Send(&msg)
		if err != nil {
			fmt.Println("发送模板消息失败")
			fmt.Println(err)
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"status": -1, "fail_content": err})
		} else {
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"action_at": time.Now().Format("2006-01-02 15:04:05"), "status": 1})
		}
	}
}
func sendSmsMessage(mydb *Mydb) {
	var cid string
	for {
		r := pool.Conn.Get()

		key, err := redis.String(r.Do("Rpop", "sms:message"))
		if err != nil {
			r.Close()
			continue
		}
		content, err := redis.String(r.Do("Get", key))
		if err != nil {
			r.Close()
			continue
		}
		var redisData map[string]string
		err = json.Unmarshal([]byte(content), &redisData)
		if err != nil {
			r.Close()
			continue
		}
		if _, ok := redisData["job"]; !ok {
			r.Close()
			continue
		}
		cid = redisData["cid"]
		cfg, err := redis.StringMap(r.Do("HgetAll", "wechat_config:"+cid))
		if err != nil {
			r.Close()
			continue
		}
		var sms *util.Mwsms
		sms = util.NewMwsms("http://"+cfg["sms_ip"], cfg["sms_address"], cfg["sms_port"], cfg["sms_account"], cfg["sms_passcode"], "")
		sendResult := sms.Send(redisData["phone"], redisData["job"], "*", time.Now().UnixNano(), 1)
		fmt.Println(sendResult)
		if !sendResult {
			r.Close()
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"status": -1, "fail_content": err})
			continue
		}
		mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"action_at": time.Now().Format("2006-01-02 15:04:05"), "status": 1})
	}
}

func init() {
	startTime = time.Now().UnixNano()
	fmt.Println("init start")
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		panic(err)
	}
	pool = cache.NewRedis(&cache.RedisOpts{
		Host:     cfg.Redis.Host,
		Password: cfg.Redis.Auth,
		Port:     cfg.Redis.Port,
		Database: cfg.Redis.DB,
	})
	mydb = NewMydb(cfg)
}
func main() {
	defer mydb.db.Close()
	run(mydb)
	endTime = time.Now().UnixNano()
	fmt.Println("used time: ", (endTime-startTime)/1000/1000/1000/1000)
	select {}
}
