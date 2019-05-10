package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/beevik/etree"
)

var (
	errCode = []int{-1, -12, -13, -14, -999, -10001, -10003, -10011, -10029, -10030, -10031, -10057, -10056}
)

// Mwsms 梦网短信接口
type Mwsms struct {
	ServerIP       string
	ServerAddr     string
	ServerPort     string
	ServerUserName string
	ServerPasscode string
	ServerChannel  string
}

// Msg 参数
type Msg struct {
	UserID     string `json:"userId"`
	Password   string `json:"password"`
	PszMobis   string `json:"pszMobis"`
	PszMsg     string `json:"pszMsg"`
	IMobiCount int    `json:"iMobiCount"`
	PszSubPort string `json:"pszSubPort"`
	MsgID      int64  `json:"msgId"`
}

// NewMwsms 获取实例
func NewMwsms(ip string, addr string, port string, username string, password string, channel string) *Mwsms {
	return &Mwsms{
		ServerIP:       ip,
		ServerAddr:     addr,
		ServerPort:     port,
		ServerUserName: username,
		ServerPasscode: password,
		ServerChannel:  channel,
	}
}

// Send 发送短信
func (sms *Mwsms) Send(phone string, content string, pszSubPort string, msgID int64, iMobiCount int) bool {
	// msg := Msg{
	// 	UserID:     sms.ServerUserName,
	// 	Password:   sms.ServerPasscode,
	// 	PszMobis:   phone,
	// 	PszMsg:     content,
	// 	IMobiCount: iMobiCount,
	// 	PszSubPort: pszSubPort,
	// 	MsgID:      msgID,
	// }
	// fmt.Println(msg)
	// jsonData, err := json.Marshal(msg)
	// if err != nil {
	// 	return false
	// }
	// body := bytes.NewBuffer(jsonData)

	// response, err := http.Post(sms.ServerIP+":"+sms.ServerPort+sms.ServerAddr, "application/json", body)
	param := strings.NewReader("userId=" + sms.ServerUserName + "&password=" + sms.ServerPasscode + "&pszMobis=" + phone + "&pszMsg=" + content + "&iMobiCount=" + strconv.Itoa(iMobiCount) + "&msgId=" + strconv.FormatInt(msgID, 10))
	// fmt.Println(param)
	response, err := http.Post(sms.ServerIP+":"+sms.ServerPort+sms.ServerAddr, "application/x-www-form-urlencoded", param)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return false
	}
	returnData, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(returnData))
	if err != nil {
		return false
	}
	doc := etree.NewDocument()
	err = doc.ReadFromBytes(returnData)
	if err != nil {
		return false
	}
	element := doc.SelectElement("string")
	if element == nil {
		return false
	}
	code, err := strconv.Atoi(element.Text())
	if err != nil {
		return false
	}
	if inArray(code, errCode) {
		return false
	}
	return true
}

// in_array 判断是否在错误码数组中
func inArray(value int, array []int) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}
