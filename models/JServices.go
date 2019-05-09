package models

import "database/sql"

type JServices struct {
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Class         sql.NullString  `gorm:"column:class" json:"class"`
	Comno         sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode     sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	Costprice     sql.NullFloat64 `gorm:"column:costprice" json:"costprice"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image         sql.NullString  `gorm:"column:image" json:"image"`
	Isgive        sql.NullInt64   `gorm:"column:isgive" json:"isgive"`
	Isintegral    sql.NullInt64   `gorm:"column:isintegral" json:"isintegral"`
	Isreserve     sql.NullInt64   `gorm:"column:isreserve" json:"isreserve"`
	Itemsname     sql.NullString  `gorm:"column:itemsname" json:"itemsname"`
	Itemsnumber   sql.NullString  `gorm:"column:itemsnumber" json:"itemsnumber"`
	Itemsprice    sql.NullFloat64 `gorm:"column:itemsprice" json:"itemsprice"`
	Itemstype     sql.NullString  `gorm:"column:itemstype" json:"itemstype"`
	Mainworker    sql.NullString  `gorm:"column:mainworker" json:"mainworker"`
	Memberprice   sql.NullFloat64 `gorm:"column:memberprice" json:"memberprice"`
	Remark        sql.NullString  `gorm:"column:remark" json:"remark"`
	Royalty       sql.NullFloat64 `gorm:"column:royalty" json:"royalty"`
	Royaltymethod sql.NullString  `gorm:"column:royaltymethod" json:"royaltymethod"`
	Servicecycle  sql.NullInt64   `gorm:"column:servicecycle" json:"servicecycle"`
	StoreName     sql.NullString  `gorm:"column:store_name" json:"store_name"`
	Storeid       sql.NullString  `gorm:"column:storeid" json:"storeid"`
	Subworker     sql.NullString  `gorm:"column:subworker" json:"subworker"`
	SubworkerID   sql.NullString  `gorm:"column:subworker_id" json:"subworker_id"`
	Thumb         sql.NullString  `gorm:"column:thumb" json:"thumb"`
	Workhour      sql.NullFloat64 `gorm:"column:workhour" json:"workhour"`
	Workplace     sql.NullString  `gorm:"column:workplace" json:"workplace"`
	Worktype      sql.NullString  `gorm:"column:worktype" json:"worktype"`
}

// TableName sets the insert table name for this struct type
func (j *JServices) TableName() string {
	return "j_services"
}
