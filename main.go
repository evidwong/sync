package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"baisiyi.net/models"

	"baisiyi.net/cache"

	"baisiyi.net/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/silenceper/wechat/context"
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
	configFile = "./config.yaml"
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

//Config 微信配置
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	PayMchID       string //支付 - 商户 ID
	PayNotifyURL   string //支付 - 接受微信支付结果通知的接口地址
	PayKey         string //支付 - 商户后台设置的支付 key
	Cache          cache.Redis
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
		// go sendSmsMessage()
	}
}

func sendWechatTemplateMessage(mydb *Mydb) {
	defer mydb.db.Close()
	var cid = 0
	for {
		var jobs []models.RemindJob
		mydb.db.Where("type='wechat' AND action_at='0000-00-00 00:00:00' AND status=0").Find(&jobs)
		if len(jobs) <= 0 {
			time.Sleep(time.Second * 10)
		} else {
			for _, row := range jobs {
				wechat := NewWechat(&Config{
					AppID:     "wxb8b1c53731dcb67f",
					AppSecret: "461a3886c314edac6c8acd430c35378c",
					Token:     "test",
					Cache:     pool,
				})
				accessToken, err := wechat.Context.GetAccessToken()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			time.Sleep(time.Second * 5)
		}
	}
}
func sendSmsMessage() {

}

func init() {
	startTime = time.Now().UnixNano()
	fmt.Println("init start")
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		panic(err)
	}
	pool = cache.NewRedis(cfg)
	mydb = NewMydb(cfg)
	fmt.Printf("%T", mydb)

}
func main() {
	run(mydb)
	endTime = time.Now().UnixNano()
	fmt.Println("used time: ", (endTime-startTime)/1000/1000/1000/1000)
	select {}
}
