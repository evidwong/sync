package models

import (
	"database/sql"
	"time"
)

type Company struct {
	Address             string          `gorm:"column:address" json:"address"`
	Addtime             int             `gorm:"column:addtime" json:"addtime"`
	APIKey              sql.NullString  `gorm:"column:api_key" json:"api_key"`
	ApprovedExpire      time.Time       `gorm:"column:approved_expire" json:"approved_expire"`
	Bdcode              sql.NullString  `gorm:"column:bdcode" json:"bdcode"`
	CrontabTime         int             `gorm:"column:crontab_time" json:"crontab_time"`
	Domian              string          `gorm:"column:domian" json:"domian"`
	Email               string          `gorm:"column:email" json:"email"`
	ExpireDate          time.Time       `gorm:"column:expire_date" json:"expire_date"`
	GdwzID              int             `gorm:"column:gdwz_id" json:"gdwz_id"`
	GdwzModel           int             `gorm:"column:gdwz_model" json:"gdwz_model"`
	ID                  int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Idcode              sql.NullString  `gorm:"column:idcode" json:"idcode"`
	Idhash              sql.NullString  `gorm:"column:idhash" json:"idhash"`
	Idpass              sql.NullString  `gorm:"column:idpass" json:"idpass"`
	Imgurl              string          `gorm:"column:imgurl" json:"imgurl"`
	IndustryID1         int             `gorm:"column:industry_id1" json:"industry_id1"`
	IndustryID2         int             `gorm:"column:industry_id2" json:"industry_id2"`
	IsInitSql           int             `gorm:"column:is_init_sql" json:"is_init_sql"`
	IsMerchant          int             `gorm:"column:is_merchant" json:"is_merchant"`
	Logo                string          `gorm:"column:logo" json:"logo"`
	Name                string          `gorm:"column:name" json:"name"`
	Person              string          `gorm:"column:person" json:"person"`
	Phone               string          `gorm:"column:phone" json:"phone"`
	QrcodePic           sql.NullString  `gorm:"column:qrcodePic" json:"qrcodePic"`
	Registration        sql.NullString  `gorm:"column:registration" json:"registration"`
	Remark              sql.NullString  `gorm:"column:remark" json:"remark"`
	SalesMan            sql.NullString  `gorm:"column:sales_man" json:"sales_man"`
	SalesPrice          sql.NullFloat64 `gorm:"column:sales_price" json:"sales_price"`
	ShopID              int             `gorm:"column:shop_id" json:"shop_id"`
	ShortName           string          `gorm:"column:short_name" json:"short_name"`
	SmsLave             sql.NullInt64   `gorm:"column:sms_lave" json:"sms_lave"`
	SmsMarketingChannel sql.NullString  `gorm:"column:sms_marketing_channel" json:"sms_marketing_channel"`
	SmsNum              sql.NullInt64   `gorm:"column:sms_num" json:"sms_num"`
	SmsType             sql.NullInt64   `gorm:"column:sms_type" json:"sms_type"`
	SmsVcodeChannel     sql.NullString  `gorm:"column:sms_vcode_channel" json:"sms_vcode_channel"`
	SqlData             sql.NullString  `gorm:"column:sql_data" json:"sql_data"`
	Status              int             `gorm:"column:status" json:"status"`
	SubID               sql.NullString  `gorm:"column:sub_id" json:"sub_id"`
	SubIDJSON           sql.NullString  `gorm:"column:sub_id_json" json:"sub_id_json"`
	SubStoreID          sql.NullString  `gorm:"column:sub_store_id" json:"sub_store_id"`
	SubStoreIDJSON      sql.NullString  `gorm:"column:sub_store_id_json" json:"sub_store_id_json"`
	Tel                 string          `gorm:"column:tel" json:"tel"`
	ThumbURL            string          `gorm:"column:thumb_url" json:"thumb_url"`
	UpdateTime          sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (c *Company) TableName() string {
	return "company"
}
