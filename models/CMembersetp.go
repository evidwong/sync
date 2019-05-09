package models

import "database/sql"

type CMembersetp struct {
	Attribute1    sql.NullString  `gorm:"column:attribute1" json:"attribute1"`
	Attribute2    sql.NullString  `gorm:"column:attribute2" json:"attribute2"`
	Attribute3    sql.NullString  `gorm:"column:attribute3" json:"attribute3"`
	Attribute4    sql.NullString  `gorm:"column:attribute4" json:"attribute4"`
	Attribute5    sql.NullString  `gorm:"column:attribute5" json:"attribute5"`
	Attribute6    sql.NullString  `gorm:"column:attribute6" json:"attribute6"`
	Attribute7    sql.NullString  `gorm:"column:attribute7" json:"attribute7"`
	Attribute8    sql.NullString  `gorm:"column:attribute8" json:"attribute8"`
	Attributecode string          `gorm:"column:attributecode" json:"attributecode"`
	Bonusamount   sql.NullFloat64 `gorm:"column:bonusamount" json:"bonusamount"`
	Canoutqty     sql.NullFloat64 `gorm:"column:canoutqty" json:"canoutqty"`
	Chargeamount  sql.NullFloat64 `gorm:"column:chargeamount" json:"chargeamount"`
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno         sql.NullString  `gorm:"column:comno" json:"comno"`
	ConsumePrice  sql.NullFloat64 `gorm:"column:consume_price" json:"consume_price"`
	Costprice     sql.NullFloat64 `gorm:"column:costprice" json:"costprice"`
	Englishname   sql.NullString  `gorm:"column:englishname" json:"englishname"`
	Fcode         sql.NullString  `gorm:"column:fcode" json:"fcode"`
	Groupno       sql.NullString  `gorm:"column:groupno" json:"groupno"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid       sql.NullInt64   `gorm:"column:indexid" json:"indexid"`
	Isintegral    sql.NullInt64   `gorm:"column:isintegral" json:"isintegral"`
	Itemcode      sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
	Itemlib       sql.NullString  `gorm:"column:itemlib" json:"itemlib"`
	Itemname      sql.NullString  `gorm:"column:itemname" json:"itemname"`
	LimitNumber   sql.NullFloat64 `gorm:"column:limit_number" json:"limit_number"`
	Membersetcode sql.NullString  `gorm:"column:membersetcode" json:"membersetcode"`
	Orderid       sql.NullInt64   `gorm:"column:orderid" json:"orderid"`
	Pid           sql.NullInt64   `gorm:"column:pid" json:"pid"`
	Remarks       sql.NullString  `gorm:"column:remarks" json:"remarks"`
	Setlimit      sql.NullInt64   `gorm:"column:setlimit" json:"setlimit"`
	Setlimitunit  sql.NullString  `gorm:"column:setlimitunit" json:"setlimitunit"`
	StandardPrice sql.NullFloat64 `gorm:"column:standard_price" json:"standard_price"`
	StoreID       sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	Unit          sql.NullString  `gorm:"column:unit" json:"unit"`
	Userdefcode   sql.NullString  `gorm:"column:userdefcode" json:"userdefcode"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetp) TableName() string {
	return "c_membersetp"
}
