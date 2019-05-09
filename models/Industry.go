package models

import (
	"database/sql"
	"time"
)

type Industry struct {
	AddTime        time.Time       `gorm:"column:add_time" json:"add_time"`
	Address        sql.NullString  `gorm:"column:address" json:"address"`
	Area           sql.NullString  `gorm:"column:area" json:"area"`
	Auth           sql.NullInt64   `gorm:"column:auth" json:"auth"`
	BranchIds      sql.NullString  `gorm:"column:branch_ids" json:"branch_ids"`
	Brands         sql.NullString  `gorm:"column:brands" json:"brands"`
	BrandsID       sql.NullString  `gorm:"column:brands_id" json:"brands_id"`
	Business       sql.NullString  `gorm:"column:business" json:"business"`
	Cid            sql.NullInt64   `gorm:"column:cid" json:"cid"`
	City           sql.NullString  `gorm:"column:city" json:"city"`
	Code           sql.NullString  `gorm:"column:code" json:"code"`
	Comno          sql.NullString  `gorm:"column:comno" json:"comno"`
	CompanyName    sql.NullString  `gorm:"column:company_name" json:"company_name"`
	CompanyNature  sql.NullString  `gorm:"column:company_nature" json:"company_nature"`
	CompanyType    sql.NullInt64   `gorm:"column:company_type" json:"company_type"`
	Detail         sql.NullString  `gorm:"column:detail" json:"detail"`
	ID             int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isshow         sql.NullInt64   `gorm:"column:isshow" json:"isshow"`
	LImg           sql.NullString  `gorm:"column:l_img" json:"l_img"`
	Lat            sql.NullFloat64 `gorm:"column:lat" json:"lat"`
	Level          sql.NullString  `gorm:"column:level" json:"level"`
	Log            sql.NullFloat64 `gorm:"column:log" json:"log"`
	Manager        sql.NullString  `gorm:"column:manager" json:"manager"`
	ManagerTel     sql.NullString  `gorm:"column:manager_tel" json:"manager_tel"`
	MerchantWechat sql.NullString  `gorm:"column:merchant_wechat" json:"merchant_wechat"`
	PImg           sql.NullString  `gorm:"column:p_img" json:"p_img"`
	PraiseNumber   sql.NullInt64   `gorm:"column:praise_number" json:"praise_number"`
	Profile        sql.NullString  `gorm:"column:profile" json:"profile"`
	Province       sql.NullString  `gorm:"column:province" json:"province"`
	QImg           sql.NullString  `gorm:"column:q_img" json:"q_img"`
	Score          sql.NullInt64   `gorm:"column:score" json:"score"`
	StoreID        sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	UpdateTime     time.Time       `gorm:"column:update_time" json:"update_time"`
	Voter          sql.NullInt64   `gorm:"column:voter" json:"voter"`
	Vulnerable     sql.NullString  `gorm:"column:vulnerable" json:"vulnerable"`
	VulnerableID   sql.NullString  `gorm:"column:vulnerable_id" json:"vulnerable_id"`
}

// TableName sets the insert table name for this struct type
func (i *Industry) TableName() string {
	return "industry"
}
