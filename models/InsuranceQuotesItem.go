package models

import "database/sql"

type InsuranceQuotesItem struct {
	AddTime             int             `gorm:"column:add_time" json:"add_time"`
	Aid                 int             `gorm:"column:aid" json:"aid"`
	Auid                int             `gorm:"column:auid" json:"auid"`
	CarOwnerAmount      sql.NullFloat64 `gorm:"column:car_owner_amount" json:"car_owner_amount"`
	CarOwnerFixedAmount sql.NullFloat64 `gorm:"column:car_owner_fixed_amount" json:"car_owner_fixed_amount"`
	CarOwnerPercent     sql.NullFloat64 `gorm:"column:car_owner_percent" json:"car_owner_percent"`
	Cid                 int             `gorm:"column:cid" json:"cid"`
	City                sql.NullInt64   `gorm:"column:city" json:"city"`
	CityName            string          `gorm:"column:cityName" json:"cityName"`
	ClertAmount         sql.NullFloat64 `gorm:"column:clert_amount" json:"clert_amount"`
	ClertFixedAmount    sql.NullFloat64 `gorm:"column:clert_fixed_amount" json:"clert_fixed_amount"`
	ClertPercent        sql.NullFloat64 `gorm:"column:clert_percent" json:"clert_percent"`
	Coid                int             `gorm:"column:coid" json:"coid"`
	CompanyName         sql.NullString  `gorm:"column:companyName" json:"companyName"`
	Cost                sql.NullFloat64 `gorm:"column:cost" json:"cost"`
	Demandid            sql.NullInt64   `gorm:"column:demandid" json:"demandid"`
	EndTime             int64           `gorm:"column:end_time" json:"end_time"`
	Flagid              sql.NullInt64   `gorm:"column:flagid" json:"flagid"`
	ID                  int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isaccounting        int             `gorm:"column:isaccounting" json:"isaccounting"`
	Isdel               int             `gorm:"column:isdel" json:"isdel"`
	ItemChannel         int             `gorm:"column:itemChannel" json:"itemChannel"`
	ItemName            sql.NullString  `gorm:"column:itemName" json:"itemName"`
	ItemTypeName        sql.NullString  `gorm:"column:itemTypeName" json:"itemTypeName"`
	Itemid              string          `gorm:"column:itemid" json:"itemid"`
	Itemtype            int             `gorm:"column:itemtype" json:"itemtype"`
	Itid                int             `gorm:"column:itid" json:"itid"`
	Name                sql.NullString  `gorm:"column:name" json:"name"`
	Orderid             sql.NullString  `gorm:"column:orderid" json:"orderid"`
	PaidAmount          sql.NullFloat64 `gorm:"column:paid_amount" json:"paid_amount"`
	Pid                 int             `gorm:"column:pid" json:"pid"`
	Premium             sql.NullFloat64 `gorm:"column:premium" json:"premium"`
	Province            sql.NullInt64   `gorm:"column:province" json:"province"`
	ProvinceName        sql.NullString  `gorm:"column:provinceName" json:"provinceName"`
	Quotesid            sql.NullInt64   `gorm:"column:quotesid" json:"quotesid"`
	Rebate              sql.NullFloat64 `gorm:"column:rebate" json:"rebate"`
	RebateAmount        sql.NullFloat64 `gorm:"column:rebate_amount" json:"rebate_amount"`
	RebateFixedAmount   sql.NullFloat64 `gorm:"column:rebate_fixed_amount" json:"rebate_fixed_amount"`
	Remark              sql.NullString  `gorm:"column:remark" json:"remark"`
	StartTime           int             `gorm:"column:start_time" json:"start_time"`
	Status              int             `gorm:"column:status" json:"status"`
	StoreAmount         sql.NullFloat64 `gorm:"column:store_amount" json:"store_amount"`
	StoreFixedAmount    sql.NullFloat64 `gorm:"column:store_fixed_amount" json:"store_fixed_amount"`
	StorePercent        sql.NullFloat64 `gorm:"column:store_percent" json:"store_percent"`
	UpdateTime          int             `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceQuotesItem) TableName() string {
	return "insurance_quotes_item"
}
