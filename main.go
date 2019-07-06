package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
	"strconv"
	"net/http"
	"bytes"

	log "github.com/sirupsen/logrus"

	"github.com/silenceper/wechat/cache"

	"baisiyi.net/config"
	"baisiyi.net/util"

	"github.com/gomodule/redigo/redis"
	_ "github.com/icattlecoder/godaemon"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/template"
)

var (
	mydb          *Mydb
	pool          *cache.Redis
	redisConStr   string
	mysqlConStr   string
	logger        *log.Logger
	wmu           sync.Mutex
	smu           sync.Mutex
	pushMu		  sync.Mutex
	wechatLogChan chan string
	smsLogChan    chan string
	sysLogChan    chan string
	pushLogChan	  chan string
	startTime     time.Time
	endTime       time.Time
	domain		  string
)

// 定义常量
const (
	configFile      = "./config.yml"
	timeFormat      = "2006-01-02 15:04:05"
	dateFormat      = "2006-01-02"
	wechatCacheList = "wechat:message:template"
	smsCacheList    = "sms:message"
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
	Cid     string `json:"cid"`
	Type    string `json:"type"`
	Comno   string `json:"comno"`
	Phone   string `json:"phone"`
	Smsnum  int    `json:"smsnum"`
	Job     string `json:"job"`
	Jobtype string `json:"jobtype"`
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
	Cid            string
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

func GetPushInfo(k string) (map[string]interface{},string){
	r := pool.Conn.Get()
	_, err := r.Do("Select", 10)
	if err != nil {
		sysLogChan <- "public,,选择数据库 10 失败. " + err.Error()
		r.Close()
		return nil, ""
	}
	key,err:=redis.String(r.Do("Rpop",k));
	if err != nil {
		sysLogChan <- "public,Rpop error: " + err.Error()
		r.Close()
		return nil, ""
	}
	sysLogChan <- "public,key: " + key
	content, err := redis.String(r.Do("Get", key))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			sysLogChan <- "public,获取key对应的数据,失败. " + err.Error()
		}
		r.Do("Lpush", k, key)
		r.Close()
		return nil, ""
	}
	var msg map[string]interface{}
	err = json.Unmarshal([]byte(content), &msg)
	if err != nil {
		sysLogChan <- "public,,JSON解析key对应的数据,失败. " + err.Error()
		r.Do("Lpush", k, key)
		r.Close()
		return nil, ""
	}
	if _,ok:=msg["companyId"];!ok{
		r.Close()
		return nil, ""
	}
	if _,ok:=msg["msgType"];!ok{
		r.Close()
		return nil, ""
	}
	r.Close()
	return msg, key
}
// GetInfo 获取redis
func GetInfo(t string, k string, i int) (WechatMessage, string) {
	r := pool.Conn.Get()
	_, err := r.Do("Select", 10)
	if err != nil {
		sysLogChan <- "public,,选择数据库 10 失败. " + err.Error()
		r.Close()
		return WechatMessage{}, ""
	}
	key, err := redis.String(r.Do("Rpop", k))
	if err != nil {
		r.Close()
		return WechatMessage{}, ""
	}
	content, err := redis.String(r.Do("Get", key))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			sysLogChan <- "public,,获取key对应的数据,失败. " + err.Error()
		}
		r.Do("Lpush", k, key)
		r.Close()
		return WechatMessage{}, ""
	}
	var redisData WechatMessage
	err = json.Unmarshal([]byte(content), &redisData)
	if err != nil {
		sysLogChan <- "public,,JSON解析key对应的数据,失败. " + err.Error()
		r.Do("Lpush", k, key)
		r.Close()
		return WechatMessage{}, ""
	}
	if redisData.Cid == "" || redisData.Cid == "0" {
		r.Close()
		return WechatMessage{}, ""
	}
	r.Do("Select", 1)
	cronStr, err := redis.String(r.Do("Hget", "cron:config:"+redisData.Cid, redisData.Jobtype))
	if err != nil {
		sysLogChan <- "public,[" + redisData.Cid + "],获取商户推送配置,失败. " + err.Error()
		r.Do("Select", 10)
		r.Do("Lpush", k, key)
		r.Close()
		return WechatMessage{}, ""
	}

	var cron map[string]interface{}
	err = json.Unmarshal([]byte(cronStr), &cron)
	if err != nil {
		sysLogChan <- "public,[" + redisData.Cid + "],解析商户推送配置,失败. " + err.Error()
		r.Do("Select", 10)
		r.Do("Lpush", k, key)
		r.Close()
		return WechatMessage{}, ""
	}
	if _, ok := cron["start_at"]; ok {
		start, ok := (cron["start_at"]).(string)
		if ok && start != "" && time.Now().Format("15:04:05") < start {
			sysLogChan <- "public,[" + redisData.Cid + "],推送时间不符合条件,失败. 推送开始时间: " + start
			r.Do("Select", 10)
			r.Do("Lpush", k, key)
			r.Close()
			return WechatMessage{}, ""
		}
	}
	if _, ok := cron["end_at"]; ok {
		end, ok := (cron["end_at"]).(string)
		if ok && end != "" && time.Now().Format("15:04:05") > end {
			sysLogChan <- "wechat,[" + redisData.Cid + "],推送时间不符合条件,失败. 推送结束时间: " + end
			r.Do("Select", 10)
			r.Do("Lpush", k, key)
			r.Close()
			return WechatMessage{}, ""
		}
	}
	r.Close()
	return redisData, key

}
func sendWechatTemplateMessage(mydb *Mydb, i int) {
	var cid string
	for {
		wmu.Lock()
		job, key := GetInfo("wechat", wechatCacheList, i)
		if key == "" {
			wmu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		if job.Cid == "" || job.Cid == "0" || job.Job == "" {
			wmu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		cid = job.Cid
		r := pool.Conn.Get()
		_, err := r.Do("Select", 1)
		if err != nil || job.Job == "" {
			wechatLogChan <- "wechat,[" + cid + "],切换到 1 号数据库,失败. " + err.Error()
			r.Do("Lpush", wechatCacheList, key)
			r.Close()
			wmu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}

		cfg, err := redis.StringMap(r.Do("HgetAll", "wechat_config:"+cid))
		if err != nil {
			wechatLogChan <- "wechat,[" + cid + "],获取商户配置信息,失败. " + err.Error()
			r.Do("Select", 10)
			r.Do("Lpush", wechatCacheList, key)
			r.Close()
			wmu.Unlock()
			time.Sleep(time.Second * 3)
			continue
		}
		s,_:=json.Marshal(cfg)
		wechatLogChan <- "wechat,[" + cid + "],config,. " + string(s)
		
		wechat := NewWechat(&Config{
			AppID:     cfg["appid"],
			AppSecret: cfg["appsecret"],
			Token:     cfg["token"],
			Cache:     pool,
			Cid:       cid,
		})
		wechatLogChan <- "wechat,[" + cid + "],job,. " + job.Job
		var msg template.Message
		err = json.Unmarshal([]byte(job.Job), &msg)
		if err != nil {
			wechatLogChan <- "wechat,[" + cid + "],JSON解析微信模板消息,失败. " + err.Error()
			r.Do("Select", 10)
			r.Do("Lpush", wechatCacheList, key)
			r.Close()
			wmu.Unlock()
			time.Sleep(time.Second * 2)
			continue
		}
		tplMsg := template.NewTemplate(wechat.Context)
		_, err = tplMsg.Send(&msg)
		if err != nil {
			r.Do("Select", 10)
			r.Do("Lpush", wechatCacheList, key)
			wechatLogChan <- "wechat,[" + cid + "],发送微信模板消息,失败. " + err.Error()
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"status": -1, "fail_content": err.Error()})
		} else {
			wechatLogChan <- "wechat,[" + cid + "],发送微信模板消息,成功. "
			r.Do("Select", 10)
			r.Do("del", key)
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"action_at": time.Now().Format(timeFormat), "status": 1})
		}
		r.Close()
		wmu.Unlock()
		time.Sleep(time.Second * 3)
	}
}
func sendSmsMessage(mydb *Mydb, i int) {
	var cid string
	for {
		smu.Lock()
		// err := recover()
		job, key := GetInfo("sms", smsCacheList, i)
		if key == "" {
			smu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		if job.Cid == "0" || job.Cid == "" || job.Job == "" {
			smu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}

		cid = job.Cid
		r := pool.Conn.Get()
		_, err := r.Do("Select", 1)
		if err != nil || job.Job == "" {
			smsLogChan <- "sms,[" + cid + "],切换到 1 号数据库,失败. " + err.Error()
			r.Do("Lpush", smsCacheList, key)
			r.Close()
			smu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		cfg, err := redis.StringMap(r.Do("HgetAll", "wechat_config:"+cid))
		if err != nil {
			smsLogChan <- "sms,[" + cid + "],获取配置信息,失败. " + err.Error()
			r.Do("Select", 10)
			r.Do("Lpush", smsCacheList, key)
			r.Close()
			smu.Unlock()
			time.Sleep(time.Second * 3)
			continue
		}
		_, err = r.Do("hIncrBy", "company:"+cid, "sms_lave", -1)
		if err != nil {
			smsLogChan <- "sms,[" + cid + "],获取商户信息,失败. " + err.Error()
			r.Do("Select", 10)
			r.Do("Lpush", smsCacheList, key)
			r.Close()
			smu.Unlock()
			time.Sleep(time.Second * 3)
			continue
		}
		var sms *util.Mwsms
		sms = util.NewMwsms("http://"+cfg["sms_ip"], cfg["sms_address"], cfg["sms_port"], cfg["sms_account"], cfg["sms_passcode"], "")
		err = sms.Send(job.Phone, job.Job, "*", time.Now().UnixNano(), 1)
		if err != nil {
			r.Do("hIncrBy", "company:"+cid, "sms_lave")
			r.Do("Select", 10)
			r.Do("Lpush", smsCacheList, key)
			smsLogChan <- "[" + cid + "] 发送短信 失败. " + err.Error()
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"status": -1, "fail_content": err.Error()})
		} else {
			smsLogChan <- "[" + cid + "] 发送短信 成功. "
			r.Do("Select", 10)
			r.Do("del", key)
			mydb.db.Table("remind_job").Where("redis_key_index=?", key).Update(map[string]interface{}{"action_at": time.Now().Format(timeFormat), "status": 1})
			mydb.db.Table("r_smsrecord").Where("flag_content=?", key).Update(map[string]interface{}{"send_time": time.Now().Format(timeFormat), "status": 1})
			// mydb.db.Table("company").Where("id=?", cid).Update(map[string]interface{}{"sms_slave": gorm.Expr("sms_lave-?", 1)})
		}
		r.Close()
		smu.Unlock()
		time.Sleep(time.Second * 3)
	}
}
func sendPushMessage(i int){
	for{
		pushMu.Lock()
		var k string="Push_Message_Async:"+strconv.Itoa(i)
		job, key := GetPushInfo(k)
		if key == "" || job==nil {
			pushMu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		if job["companyId"]== 0{
			pushLogChan <- "无效的cid"
			pushMu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		r := pool.Conn.Get()
		var url string
		
		if v,ok:=(job["msgType"]).(float64);ok && v>0{
			url=domain+"/Api/Sms/"
		}else if v,ok:=(job["msgType"]).(float64);ok && v==0{
			url=domain+"/Api/Api/"
		}else{
			r.Do("Lpush",k,key)
			pushLogChan <- "解析URL地址失败 url: "+url
			r.Close()
			pushMu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		
		if v,ok:=(job["push"]).(string);ok{
			url=url+v
		}
		param,err:=json.Marshal(job)
		if err!=nil{
			r.Do("Lpush",k,key)
			pushLogChan <- "json.Marshal请求参数失败"
			r.Close()
			pushMu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		response, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte(param)))
		if err != nil {
			pushLogChan <- "请求接口["+url+"] 参数："+string(param)+"失败"
			r.Do("Lpush",k,key)
			r.Close()
			pushMu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		if response.StatusCode != http.StatusOK {
			pushLogChan <- "请求接口["+url+"] 参数："+string(param)+"失败"
			r.Do("Lpush",k,key)
			r.Close()
			pushMu.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		// body, _ := ioutil.ReadAll(response.Body)
		r.Do("del",key)
		response.Body.Close()
		r.Close()
		pushMu.Unlock()
		time.Sleep(time.Second * 2)
	}
}

func init() {
	startTime = time.Now()
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		panic(err)
	}
	domain=cfg.Web.Domain
	wechatLogChan = make(chan string, 10000)
	smsLogChan = make(chan string, 100000)
	sysLogChan = make(chan string, 100000)
	pushLogChan = make(chan string, 100000)

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
func writeLog(filename string, logInfo string) {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: timeFormat,
	})
	logFileHandle, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 666)
	if err != nil {
		sysLogChan <- "启用用日志文件[" + filename + "]失败. " + err.Error()
		logFileHandle.Close()
		fmt.Println("启用用日志文件[" + filename + "]失败. " + err.Error())
		return
	}
	log.SetOutput(logFileHandle)
	log.Info(logInfo)
	logFileHandle.Close()
}
func main() {
	defer mydb.db.Close()
	sysLogChan <- "system start"
	go func() {
		for {
			var filename, logInfo string
			select {
			case logInfo = <-wechatLogChan:
				filename = "./logs/wechat-log_" + time.Now().Format(dateFormat) + ".log"
			case logInfo = <-smsLogChan:
				filename = "./logs/sms-log_" + time.Now().Format(dateFormat) + ".log"
			case logInfo = <-sysLogChan:
				filename = "./logs/sys-log_" + time.Now().Format(dateFormat) + ".log"
			case logInfo = <-pushLogChan:
				filename = "./logs/push-log_" + time.Now().Format(dateFormat) + ".log"
			default:
				time.Sleep(time.Second * 10)
			}
			if filename != "" && logInfo != "" {
				writeLog(filename, logInfo)
				time.Sleep(time.Second * 5)
			}
		}
	}()
	for i := 0; i < 5; i++ {
		go sendWechatTemplateMessage(mydb, i)
		go sendSmsMessage(mydb, i)
	}

	for i:=0;i<8;i++{
		// go sendPushMessage(i);
	}

	endTime = time.Now()
	var dur = endTime.Sub(startTime)
	fmt.Println("used time: ", dur.String())
	select {}
}
