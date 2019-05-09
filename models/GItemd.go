package models

import (
	"database/sql"
	"time"
)

type GItemd struct {
	Bin                    sql.NullString  `gorm:"column:bin" json:"bin"`
	Bonusamount            sql.NullFloat64 `gorm:"column:bonusamount" json:"bonusamount"`
	Cid                    sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno                  string          `gorm:"column:comno" json:"comno"`
	Costprice              sql.NullFloat64 `gorm:"column:costprice" json:"costprice"`
	Firstindate            time.Time       `gorm:"column:firstindate" json:"firstindate"`
	Fobcurrency            sql.NullString  `gorm:"column:fobcurrency" json:"fobcurrency"`
	Fobprice               sql.NullFloat64 `gorm:"column:fobprice" json:"fobprice"`
	Fobrate                sql.NullFloat64 `gorm:"column:fobrate" json:"fobrate"`
	ID                     int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Inprice                sql.NullFloat64 `gorm:"column:inprice" json:"inprice"`
	Itemcode               string          `gorm:"column:itemcode" json:"itemcode"`
	Itemname               sql.NullString  `gorm:"column:itemname" json:"itemname"`
	Itemremarks            sql.NullString  `gorm:"column:itemremarks" json:"itemremarks"`
	LastInfo               sql.NullString  `gorm:"column:lastInfo" json:"lastInfo"`
	Lastcostprice          sql.NullFloat64 `gorm:"column:lastcostprice" json:"lastcostprice"`
	Lastindate             time.Time       `gorm:"column:lastindate" json:"lastindate"`
	Lastinqty              sql.NullFloat64 `gorm:"column:lastinqty" json:"lastinqty"`
	Lastoutdate            time.Time       `gorm:"column:lastoutdate" json:"lastoutdate"`
	Lastpurchasedatetime   time.Time       `gorm:"column:lastpurchasedatetime" json:"lastpurchasedatetime"`
	Lastsalesdatetime      time.Time       `gorm:"column:lastsalesdatetime" json:"lastsalesdatetime"`
	Location               sql.NullString  `gorm:"column:location" json:"location"`
	Mainworkeramount       sql.NullFloat64 `gorm:"column:mainworkeramount" json:"mainworkeramount"`
	Mainworkerbonus        sql.NullInt64   `gorm:"column:mainworkerbonus" json:"mainworkerbonus"`
	Mainworkerkpiamount    sql.NullFloat64 `gorm:"column:mainworkerkpiamount" json:"mainworkerkpiamount"`
	Mainworkerkpiprofit    sql.NullFloat64 `gorm:"column:mainworkerkpiprofit" json:"mainworkerkpiprofit"`
	Mainworkerkpiqty       sql.NullFloat64 `gorm:"column:mainworkerkpiqty" json:"mainworkerkpiqty"`
	Mainworkerpercent      sql.NullFloat64 `gorm:"column:mainworkerpercent" json:"mainworkerpercent"`
	Mainworkerprofit       sql.NullFloat64 `gorm:"column:mainworkerprofit" json:"mainworkerprofit"`
	Maxqty                 sql.NullFloat64 `gorm:"column:maxqty" json:"maxqty"`
	Minqty                 sql.NullFloat64 `gorm:"column:minqty" json:"minqty"`
	Pid                    sql.NullInt64   `gorm:"column:pid" json:"pid"`
	Prefprice1             sql.NullFloat64 `gorm:"column:prefprice1" json:"prefprice1"`
	Prefprice10            sql.NullFloat64 `gorm:"column:prefprice10" json:"prefprice10"`
	Prefprice2             sql.NullFloat64 `gorm:"column:prefprice2" json:"prefprice2"`
	Prefprice3             sql.NullFloat64 `gorm:"column:prefprice3" json:"prefprice3"`
	Prefprice4             sql.NullFloat64 `gorm:"column:prefprice4" json:"prefprice4"`
	Prefprice5             sql.NullFloat64 `gorm:"column:prefprice5" json:"prefprice5"`
	Prefprice6             sql.NullFloat64 `gorm:"column:prefprice6" json:"prefprice6"`
	Prefprice7             sql.NullFloat64 `gorm:"column:prefprice7" json:"prefprice7"`
	Prefprice8             sql.NullFloat64 `gorm:"column:prefprice8" json:"prefprice8"`
	Prefprice9             sql.NullFloat64 `gorm:"column:prefprice9" json:"prefprice9"`
	Price                  sql.NullFloat64 `gorm:"column:price" json:"price"`
	Recpersonnameamount    sql.NullFloat64 `gorm:"column:recpersonnameamount" json:"recpersonnameamount"`
	Recpersonnamebonus     sql.NullInt64   `gorm:"column:recpersonnamebonus" json:"recpersonnamebonus"`
	Recpersonnamekpiamount sql.NullFloat64 `gorm:"column:recpersonnamekpiamount" json:"recpersonnamekpiamount"`
	Recpersonnamekpiprofit sql.NullFloat64 `gorm:"column:recpersonnamekpiprofit" json:"recpersonnamekpiprofit"`
	Recpersonnamekpiqty    sql.NullFloat64 `gorm:"column:recpersonnamekpiqty" json:"recpersonnamekpiqty"`
	Recpersonnamepercent   sql.NullFloat64 `gorm:"column:recpersonnamepercent" json:"recpersonnamepercent"`
	Recpersonnameprofit    sql.NullFloat64 `gorm:"column:recpersonnameprofit" json:"recpersonnameprofit"`
	Refprice1              sql.NullFloat64 `gorm:"column:refprice1" json:"refprice1"`
	Refprice2              sql.NullFloat64 `gorm:"column:refprice2" json:"refprice2"`
	Refprice3              sql.NullFloat64 `gorm:"column:refprice3" json:"refprice3"`
	Refprice4              sql.NullFloat64 `gorm:"column:refprice4" json:"refprice4"`
	Safeqty                sql.NullFloat64 `gorm:"column:safeqty" json:"safeqty"`
	Salesleaderamount      sql.NullFloat64 `gorm:"column:salesleaderamount" json:"salesleaderamount"`
	Salesleaderbonus       sql.NullInt64   `gorm:"column:salesleaderbonus" json:"salesleaderbonus"`
	Salesleaderkpiamount   sql.NullFloat64 `gorm:"column:salesleaderkpiamount" json:"salesleaderkpiamount"`
	Salesleaderkpiprofit   sql.NullFloat64 `gorm:"column:salesleaderkpiprofit" json:"salesleaderkpiprofit"`
	Salesleaderkpiqty      sql.NullFloat64 `gorm:"column:salesleaderkpiqty" json:"salesleaderkpiqty"`
	Salesleaderpercent     sql.NullFloat64 `gorm:"column:salesleaderpercent" json:"salesleaderpercent"`
	Salesleaderprofit      sql.NullFloat64 `gorm:"column:salesleaderprofit" json:"salesleaderprofit"`
	Stdcostprice           sql.NullFloat64 `gorm:"column:stdcostprice" json:"stdcostprice"`
	StoreID                sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	Suppliercode           sql.NullString  `gorm:"column:suppliercode" json:"suppliercode"`
	Suppliername           sql.NullString  `gorm:"column:suppliername" json:"suppliername"`
	Tradeprice             sql.NullFloat64 `gorm:"column:tradeprice" json:"tradeprice"`
	Transferprice          sql.NullFloat64 `gorm:"column:transferprice" json:"transferprice"`
}

// TableName sets the insert table name for this struct type
func (g *GItemd) TableName() string {
	return "g_itemd"
}
