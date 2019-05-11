package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

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
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logFileName := "./logs/system.log"
	logFileHandle, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Panicln("fail to create " + logFileName + " file!")
	}
	log.SetOutput(logFileHandle)
	logFileHandle.Close()
	for i := 0; i < 1; i++ {
		go sendWechatTemplateMessage(mydb)
		go sendSmsMessage(mydb)
	}
}

// GetInfo 获取redis
func GetInfo(t string, k string) (redis.Conn, *WechatMessage, *os.File, string) {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	filename := "./logs/" + t + "_" + time.Now().Format("2006-01-02") + ".log"
	logFileHandle, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Info("创建或使用日志文件失败. " + err.Error())
		time.Sleep(time.Second * 5)
		return nil, nil, nil, ""
	}
	log.SetOutput(logFileHandle)

	r := pool.Conn.Get()
	_, err = r.Do("Select", 10)
	if err != nil {
		r.Close()
		log.Info("选择数据库 10 失败. " + err.Error())
		time.Sleep(time.Second * 3)
		return nil, nil, nil, ""
	}
	key, err := redis.String(r.Do("Rpop", k))
	if err != nil {
		r.Close()
		log.Info("Rpop列表 失败. " + err.Error())
		time.Sleep(time.Second * 3)
		return nil, nil, nil, ""
	}
	content, err := redis.String(r.Do("Get", key))
	if err != nil {
		r.Close()
		log.Info("获取key对应的数据 失败. " + err.Error())
		return nil, nil, nil, ""
	}
	var redisData WechatMessage
	err = json.Unmarshal([]byte(content), &redisData)
	if err != nil {
		r.Close()
		log.Info("JSON解析key对应的数据 失败. " + err.Error())
		time.Sleep(time.Second * 3)
		return nil, nil, nil, ""
	}
	return r, &redisData, logFileHandle, key

}
func sendWechatTemplateMessage(mydb *Mydb) {
	var cid string
	for {
		r, job, file, key := GetInfo("wechat", "wechat:message:template")
		if r == nil {
			time.Sleep(time.Second * 5)
			continue
		}
		_, err := r.Do("Select", 1)
		if err != nil {
			r.Close()
			log.Info("切换数据库到 1 失败. " + err.Error())
			file.Close()
			time.Sleep(time.Second * 5)
			continue
		}
		cid = strconv.Itoa(job.Cid)
		cfg, err := redis.StringMap(r.Do("HgetAll", "wechat_config:"+cid))
		if err != nil {
			r.Close()
			log.Info("[" + cid + "] 获取配置信息 失败. " + err.Error())
			file.Close()
			time.Sleep(time.Second * 3)
			continue
		}
		wechat := NewWechat(&Config{
			AppID:     cfg["appid"],
			AppSecret: cfg["appsecret"],
			Token:     cfg["token"],
			Cache:     pool,
		})
		var msg template.Message
		err = json.Unmarshal([]byte(job.Job), &msg)
		if err != nil {
			r.Close()
			log.Info("[" + cid + "] JSON解析微信模板消息 失败. " + err.Error())
			file.Close()
			time.Sleep(time.Second * 2)
			continue
		}
		tplMsg := template.NewTemplate(wechat.Context)
		_, err = tplMsg.Send(&msg)
		if err != nil {
			log.Info("[" + cid + "] 发送微信模板消息 失败. " + err.Error())
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"status": -1, "fail_content": err})
		} else {
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"action_at": time.Now().Format("2006-01-02 15:04:05"), "status": 1})
		}
		r.Close()
		time.Sleep(time.Second * 3)
	}
}
func sendSmsMessage(mydb *Mydb) {
	var cid string
	for {
		r, job, file, key := GetInfo("sms", "sms:message")
		if r == nil {
			time.Sleep(time.Second * 5)
			continue
		}
		cid = strconv.Itoa(job.Cid)
		cfg, err := redis.StringMap(r.Do("HgetAll", "wechat_config:"+cid))
		if err != nil {
			log.Info("[" + cid + "] 获取配置信息 失败. " + err.Error())
			r.Close()
			file.Close()
			time.Sleep(time.Second * 3)
			continue
		}
		_, err = r.Do("hIncrBy", "company:"+cid, "sms_lave", -1)
		if err != nil {
			r.Close()
			log.Info("[" + cid + "] 获取商户信息 失败. " + err.Error())
			file.Close()
			time.Sleep(time.Second * 3)
			continue
		}
		var sms *util.Mwsms
		sms = util.NewMwsms("http://"+cfg["sms_ip"], cfg["sms_address"], cfg["sms_port"], cfg["sms_account"], cfg["sms_passcode"], "")
		sendResult := sms.Send(job.Phone, job.Job, "*", time.Now().UnixNano(), 1)
		fmt.Println(sendResult)
		if !sendResult {
			r.Do("hIncrBy", "company:"+cid, "sms_lave")
			log.Info("[" + cid + "] 发送短信 失败. " + err.Error())
			time.Sleep(time.Second * 5)
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"status": -1, "fail_content": err})
			r.Close()
			file.Close()
			continue
		} else {
			log.Info("[" + cid + "] 发送短信 成功. ")
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"action_at": time.Now().Format("2006-01-02 15:04:05"), "status": 1})
			mydb.db.Table("r_smsrecord").Where("flag_content=?", key).Update(map[string]interface{}{"send_time": time.Now().Format("2006-01-02 15:04:05"), "status": 1})
			// mydb.db.Table("company").Where("id=?", cid).Update(map[string]interface{}{"sms_slave": gorm.Expr("sms_lave-?", 1)})
		}
		r.Close()
		file.Close()
		time.Sleep(time.Second * 3)
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
		Host:        cfg.Redis.Host,
		Password:    cfg.Redis.Auth,
		Port:        cfg.Redis.Port,
		Database:    cfg.Redis.DB,
		MaxActive:   cfg.Redis.MaxActive,
		MaxIdle:     cfg.Redis.MaxIdle,
		IdleTimeout: cfg.Redis.IdleTimeout,
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
