package models

import "database/sql"

type CarOrtherinfo struct {
	Buydate                       sql.NullInt64   `gorm:"column:buydate" json:"buydate"`
	Buytime                       sql.NullInt64   `gorm:"column:buytime" json:"buytime"`
	Carid                         int             `gorm:"column:carid" json:"carid"`
	Cmodel                        sql.NullString  `gorm:"column:cmodel" json:"cmodel"`
	Color                         sql.NullString  `gorm:"column:color" json:"color"`
	Createdate                    sql.NullInt64   `gorm:"column:createdate" json:"createdate"`
	Createfileperson              sql.NullString  `gorm:"column:createfileperson" json:"createfileperson"`
	Customercode                  sql.NullString  `gorm:"column:customercode" json:"customercode"`
	DriveOuttime                  sql.NullInt64   `gorm:"column:drive_outtime" json:"drive_outtime"`
	DriveRetime                   sql.NullInt64   `gorm:"column:drive_retime" json:"drive_retime"`
	Drivelicyearcheckdate         sql.NullInt64   `gorm:"column:drivelicyearcheckdate" json:"drivelicyearcheckdate"`
	DrivelicyearcheckdateInterval sql.NullInt64   `gorm:"column:drivelicyearcheckdate_interval" json:"drivelicyearcheckdate_interval"`
	DrivelicyearcheckdateNextTime sql.NullInt64   `gorm:"column:drivelicyearcheckdate_next_time" json:"drivelicyearcheckdate_next_time"`
	Driverbirthmd                 sql.NullInt64   `gorm:"column:driverbirthmd" json:"driverbirthmd"`
	Driverfancy                   sql.NullString  `gorm:"column:driverfancy" json:"driverfancy"`
	Driverhandphone               sql.NullString  `gorm:"column:driverhandphone" json:"driverhandphone"`
	DrlicOuttime                  sql.NullInt64   `gorm:"column:drlic_outtime" json:"drlic_outtime"`
	Enviromenttodate              sql.NullInt64   `gorm:"column:enviromenttodate" json:"enviromenttodate"`
	Factory                       sql.NullString  `gorm:"column:factory" json:"factory"`
	FixHistory                    sql.NullInt64   `gorm:"column:fix_history" json:"fix_history"`
	FixMoney                      sql.NullFloat64 `gorm:"column:fix_money" json:"fix_money"`
	ID                            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Insurance                     sql.NullString  `gorm:"column:insurance" json:"insurance"`
	InsuranceInterval             sql.NullInt64   `gorm:"column:insurance_interval" json:"insurance_interval"`
	InsuranceNextTime             sql.NullInt64   `gorm:"column:insurance_next_time" json:"insurance_next_time"`
	Insurancecompany              sql.NullString  `gorm:"column:insurancecompany" json:"insurancecompany"`
	Insurancedate                 sql.NullInt64   `gorm:"column:insurancedate" json:"insurancedate"`
	InsureOuttime                 sql.NullInt64   `gorm:"column:insure_outtime" json:"insure_outtime"`
	Lastinfactorydate             sql.NullInt64   `gorm:"column:lastinfactorydate" json:"lastinfactorydate"`
	Lastinmileage                 sql.NullFloat64 `gorm:"column:lastinmileage" json:"lastinmileage"`
	Province                      sql.NullString  `gorm:"column:province" json:"province"`
	Runlicdate                    sql.NullInt64   `gorm:"column:runlicdate" json:"runlicdate"`
	RunlicdateInterval            sql.NullInt64   `gorm:"column:runlicdate_interval" json:"runlicdate_interval"`
	RunlicdateNextTime            sql.NullInt64   `gorm:"column:runlicdate_next_time" json:"runlicdate_next_time"`
	UID                           int             `gorm:"column:uid" json:"uid"`
	Vehicle                       sql.NullString  `gorm:"column:vehicle" json:"vehicle"`
	Vehiclecode                   sql.NullString  `gorm:"column:vehiclecode" json:"vehiclecode"`
	Wxhistoryamount               sql.NullFloat64 `gorm:"column:wxhistoryamount" json:"wxhistoryamount"`
	Wxhistorycount                sql.NullInt64   `gorm:"column:wxhistorycount" json:"wxhistorycount"`
	Yearcheckdate                 sql.NullInt64   `gorm:"column:yearcheckdate" json:"yearcheckdate"`
	YearcheckdateInterval         sql.NullInt64   `gorm:"column:yearcheckdate_interval" json:"yearcheckdate_interval"`
	YearcheckdateNextTime         sql.NullInt64   `gorm:"column:yearcheckdate_next_time" json:"yearcheckdate_next_time"`
}

// TableName sets the insert table name for this struct type
func (c *CarOrtherinfo) TableName() string {
	return "car_ortherinfo"
}
