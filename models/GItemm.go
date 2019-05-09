package models

import (
	"database/sql"
	"time"
)

type GItemm struct {
	Arrivedays        sql.NullInt64   `gorm:"column:arrivedays" json:"arrivedays"`
	Attribute1        sql.NullString  `gorm:"column:attribute1" json:"attribute1"`
	Attribute2        sql.NullString  `gorm:"column:attribute2" json:"attribute2"`
	Attribute3        sql.NullString  `gorm:"column:attribute3" json:"attribute3"`
	Attribute4        sql.NullString  `gorm:"column:attribute4" json:"attribute4"`
	Attribute5        sql.NullString  `gorm:"column:attribute5" json:"attribute5"`
	Attribute6        sql.NullString  `gorm:"column:attribute6" json:"attribute6"`
	Attribute7        sql.NullString  `gorm:"column:attribute7" json:"attribute7"`
	Attribute8        sql.NullString  `gorm:"column:attribute8" json:"attribute8"`
	Attributecode     sql.NullString  `gorm:"column:attributecode" json:"attributecode"`
	Bomcode           sql.NullString  `gorm:"column:bomcode" json:"bomcode"`
	Box               sql.NullString  `gorm:"column:box" json:"box"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno             sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode         sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	Deliverperiod     sql.NullInt64   `gorm:"column:deliverperiod" json:"deliverperiod"`
	Englishname       sql.NullString  `gorm:"column:englishname" json:"englishname"`
	Fcode             sql.NullString  `gorm:"column:fcode" json:"fcode"`
	Grade             sql.NullString  `gorm:"column:grade" json:"grade"`
	Groupno           sql.NullString  `gorm:"column:groupno" json:"groupno"`
	Height            sql.NullFloat64 `gorm:"column:height" json:"height"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Inqty             sql.NullInt64   `gorm:"column:inqty" json:"inqty"`
	Iscancel          sql.NullInt64   `gorm:"column:iscancel" json:"iscancel"`
	Iscanintegral     sql.NullInt64   `gorm:"column:iscanintegral" json:"iscanintegral"`
	Iscanpresent      sql.NullInt64   `gorm:"column:iscanpresent" json:"iscanpresent"`
	Itemcode          string          `gorm:"column:itemcode" json:"itemcode"`
	Itemimages        sql.NullString  `gorm:"column:itemimages" json:"itemimages"`
	Itemlib           sql.NullString  `gorm:"column:itemlib" json:"itemlib"`
	Itemmemo          sql.NullString  `gorm:"column:itemmemo" json:"itemmemo"`
	Itemmemo2         sql.NullString  `gorm:"column:itemmemo2" json:"itemmemo2"`
	Itemmemo3         sql.NullString  `gorm:"column:itemmemo3" json:"itemmemo3"`
	Itemname          sql.NullString  `gorm:"column:itemname" json:"itemname"`
	LastInfo          sql.NullString  `gorm:"column:lastInfo" json:"lastInfo"`
	Length            sql.NullFloat64 `gorm:"column:length" json:"length"`
	Minorderqty       sql.NullFloat64 `gorm:"column:minorderqty" json:"minorderqty"`
	Modifydate        time.Time       `gorm:"column:modifydate" json:"modifydate"`
	Modifyip          sql.NullString  `gorm:"column:modifyip" json:"modifyip"`
	Modifypersoncode  sql.NullString  `gorm:"column:modifypersoncode" json:"modifypersoncode"`
	Modifypersonname  sql.NullString  `gorm:"column:modifypersonname" json:"modifypersonname"`
	Modifyworkstation sql.NullString  `gorm:"column:modifyworkstation" json:"modifyworkstation"`
	Orderperiod       sql.NullInt64   `gorm:"column:orderperiod" json:"orderperiod"`
	Outorderqty       sql.NullFloat64 `gorm:"column:outorderqty" json:"outorderqty"`
	Outqty            sql.NullInt64   `gorm:"column:outqty" json:"outqty"`
	Packqty           sql.NullFloat64 `gorm:"column:packqty" json:"packqty"`
	Packunit          sql.NullString  `gorm:"column:packunit" json:"packunit"`
	Partscomno        sql.NullString  `gorm:"column:partscomno" json:"partscomno"`
	Partsintegral     sql.NullFloat64 `gorm:"column:partsintegral" json:"partsintegral"`
	Pycode            sql.NullString  `gorm:"column:pycode" json:"pycode"`
	Relationcode      sql.NullString  `gorm:"column:relationcode" json:"relationcode"`
	Relationno        sql.NullInt64   `gorm:"column:relationno" json:"relationno"`
	Singlevol         sql.NullFloat64 `gorm:"column:singlevol" json:"singlevol"`
	StoreID           sql.NullString  `gorm:"column:store_id" json:"store_id"`
	StoreName         sql.NullString  `gorm:"column:store_name" json:"store_name"`
	Suitqty           sql.NullFloat64 `gorm:"column:suitqty" json:"suitqty"`
	Sumamount         sql.NullFloat64 `gorm:"column:sumamount" json:"sumamount"`
	Sumcost           sql.NullFloat64 `gorm:"column:sumcost" json:"sumcost"`
	Tmpfield1         sql.NullString  `gorm:"column:tmpfield1" json:"tmpfield1"`
	Tmpfield2         time.Time       `gorm:"column:tmpfield2" json:"tmpfield2"`
	Type1             sql.NullString  `gorm:"column:type1" json:"type1"`
	Type2             sql.NullString  `gorm:"column:type2" json:"type2"`
	Type3             sql.NullString  `gorm:"column:type3" json:"type3"`
	Type4             sql.NullString  `gorm:"column:type4" json:"type4"`
	Type5             sql.NullString  `gorm:"column:type5" json:"type5"`
	Unit              sql.NullString  `gorm:"column:unit" json:"unit"`
	Userdefcode       sql.NullString  `gorm:"column:userdefcode" json:"userdefcode"`
	Vol               sql.NullFloat64 `gorm:"column:vol" json:"vol"`
	Wbcode            sql.NullString  `gorm:"column:wbcode" json:"wbcode"`
	Weight            sql.NullFloat64 `gorm:"column:weight" json:"weight"`
	Width             sql.NullFloat64 `gorm:"column:width" json:"width"`
}

// TableName sets the insert table name for this struct type
func (g *GItemm) TableName() string {
	return "g_itemm"
}
