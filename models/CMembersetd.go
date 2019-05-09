package models

import "database/sql"

type CMembersetd struct {
	Bonusamount   sql.NullFloat64 `gorm:"column:bonusamount" json:"bonusamount"`
	Chargeamount  sql.NullFloat64 `gorm:"column:chargeamount" json:"chargeamount"`
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Class         sql.NullString  `gorm:"column:class" json:"class"`
	ConsumePrice  sql.NullFloat64 `gorm:"column:consume_price" json:"consume_price"`
	Costprice     sql.NullFloat64 `gorm:"column:costprice" json:"costprice"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image         sql.NullString  `gorm:"column:image" json:"image"`
	Indexid       sql.NullInt64   `gorm:"column:indexid" json:"indexid"`
	Intervalday   sql.NullInt64   `gorm:"column:intervalday" json:"intervalday"`
	Isgive        sql.NullInt64   `gorm:"column:isgive" json:"isgive"`
	Isintegral    sql.NullInt64   `gorm:"column:isintegral" json:"isintegral"`
	Isreserve     sql.NullInt64   `gorm:"column:isreserve" json:"isreserve"`
	Itemsprice    sql.NullFloat64 `gorm:"column:itemsprice" json:"itemsprice"`
	Jobcode       sql.NullString  `gorm:"column:jobcode" json:"jobcode"`
	Jobname       sql.NullString  `gorm:"column:jobname" json:"jobname"`
	Jobtype       sql.NullString  `gorm:"column:jobtype" json:"jobtype"`
	Mainworker    sql.NullString  `gorm:"column:mainworker" json:"mainworker"`
	Memberprice   sql.NullFloat64 `gorm:"column:memberprice" json:"memberprice"`
	Membersetcode sql.NullString  `gorm:"column:membersetcode" json:"membersetcode"`
	Orderid       sql.NullInt64   `gorm:"column:orderid" json:"orderid"`
	Pid           sql.NullInt64   `gorm:"column:pid" json:"pid"`
	Remarks       sql.NullString  `gorm:"column:remarks" json:"remarks"`
	Royaltymethod sql.NullString  `gorm:"column:royaltymethod" json:"royaltymethod"`
	Servicecount  sql.NullFloat64 `gorm:"column:servicecount" json:"servicecount"`
	Servicecycle  sql.NullInt64   `gorm:"column:servicecycle" json:"servicecycle"`
	Setlimit      sql.NullInt64   `gorm:"column:setlimit" json:"setlimit"`
	Setlimitunit  sql.NullString  `gorm:"column:setlimitunit" json:"setlimitunit"`
	StandardPrice sql.NullFloat64 `gorm:"column:standard_price" json:"standard_price"`
	Subworker     sql.NullString  `gorm:"column:subworker" json:"subworker"`
	SubworkerID   sql.NullString  `gorm:"column:subworker_id" json:"subworker_id"`
	Thumb         sql.NullString  `gorm:"column:thumb" json:"thumb"`
	Workhour      sql.NullFloat64 `gorm:"column:workhour" json:"workhour"`
	Workplace     sql.NullString  `gorm:"column:workplace" json:"workplace"`
	Worktype      sql.NullString  `gorm:"column:worktype" json:"worktype"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetd) TableName() string {
	return "c_membersetd"
}
