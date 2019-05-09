package models

import "database/sql"

type WebDisellSetting struct {
	Ishead                    sql.NullInt64  `gorm:"column:Ishead" json:"Ishead"`
	QRcodeContent             sql.NullString `gorm:"column:QRcode_content" json:"QRcode_content"`
	QRcodeImg                 sql.NullString `gorm:"column:QRcode_img" json:"QRcode_img"`
	QRcodeThumb               sql.NullString `gorm:"column:QRcode_thumb" json:"QRcode_thumb"`
	BuyDistributor            int            `gorm:"column:buy_distributor" json:"buy_distributor"`
	Cid                       sql.NullInt64  `gorm:"column:cid" json:"cid"`
	DistributorAuditing       int            `gorm:"column:distributor_auditing" json:"distributor_auditing"`
	ExchangeStandard          sql.NullInt64  `gorm:"column:exchange_standard" json:"exchange_standard"`
	ID                        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Level1Coupon              sql.NullInt64  `gorm:"column:level_1_coupon" json:"level_1_coupon"`
	Level1CouponName          sql.NullString `gorm:"column:level_1_coupon_name" json:"level_1_coupon_name"`
	Level2Coupon              sql.NullInt64  `gorm:"column:level_2_coupon" json:"level_2_coupon"`
	Level2CouponName          sql.NullString `gorm:"column:level_2_coupon_name" json:"level_2_coupon_name"`
	Level3Coupon              sql.NullInt64  `gorm:"column:level_3_coupon" json:"level_3_coupon"`
	Level3CouponName          sql.NullString `gorm:"column:level_3_coupon_name" json:"level_3_coupon_name"`
	NewreadContent            sql.NullString `gorm:"column:newread_content" json:"newread_content"`
	RebateMechanism           sql.NullString `gorm:"column:rebate_mechanism" json:"rebate_mechanism"`
	ReminderContent           sql.NullString `gorm:"column:reminder_content" json:"reminder_content"`
	SalesorderAuditing        int            `gorm:"column:salesorder_auditing" json:"salesorder_auditing"`
	IsAutoSendIntegralForBind int            `gorm:"column:is_auto_send_integral_for_bind" json:"is_auto_send_integral_for_bind"`
}

// TableName sets the insert table name for this struct type
func (w *WebDisellSetting) TableName() string {
	return "web_disell_setting"
}
