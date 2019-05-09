package models

import "database/sql"

type PMonthbalanced struct {
	ActionID      int             `gorm:"column:ActionID;primary_key" json:"ActionID;primary_key"`
	AdjustCost    sql.NullFloat64 `gorm:"column:AdjustCost" json:"AdjustCost"`
	CGInCost      sql.NullFloat64 `gorm:"column:CGInCost" json:"CGInCost"`
	CGInQty       sql.NullFloat64 `gorm:"column:CGInQty" json:"CGInQty"`
	COMNo         string          `gorm:"column:COMNo" json:"COMNo"`
	ChangeCost    sql.NullFloat64 `gorm:"column:ChangeCost" json:"ChangeCost"`
	ChangeQty     sql.NullFloat64 `gorm:"column:ChangeQty" json:"ChangeQty"`
	DBJInCost     sql.NullFloat64 `gorm:"column:DBJInCost" json:"DBJInCost"`
	DBJInQty      sql.NullFloat64 `gorm:"column:DBJInQty" json:"DBJInQty"`
	DBJOutAmount  sql.NullFloat64 `gorm:"column:DBJOutAmount" json:"DBJOutAmount"`
	DBJOutCost    sql.NullFloat64 `gorm:"column:DBJOutCost" json:"DBJOutCost"`
	DBJOutQty     sql.NullFloat64 `gorm:"column:DBJOutQty" json:"DBJOutQty"`
	DBWInCost     sql.NullFloat64 `gorm:"column:DBWInCost" json:"DBWInCost"`
	DBWInQty      sql.NullFloat64 `gorm:"column:DBWInQty" json:"DBWInQty"`
	DBWOutAmount  sql.NullFloat64 `gorm:"column:DBWOutAmount" json:"DBWOutAmount"`
	DBWOutCost    sql.NullFloat64 `gorm:"column:DBWOutCost" json:"DBWOutCost"`
	DBWOutQty     sql.NullFloat64 `gorm:"column:DBWOutQty" json:"DBWOutQty"`
	DHInCost      sql.NullFloat64 `gorm:"column:DHInCost" json:"DHInCost"`
	DHInQty       sql.NullFloat64 `gorm:"column:DHInQty" json:"DHInQty"`
	DXInCost      sql.NullFloat64 `gorm:"column:DXInCost" json:"DXInCost"`
	DXInQty       sql.NullFloat64 `gorm:"column:DXInQty" json:"DXInQty"`
	EndCost       sql.NullFloat64 `gorm:"column:EndCost" json:"EndCost"`
	EndQty        sql.NullFloat64 `gorm:"column:EndQty" json:"EndQty"`
	InCost        sql.NullFloat64 `gorm:"column:InCost" json:"InCost"`
	InQty         sql.NullFloat64 `gorm:"column:InQty" json:"InQty"`
	JFOutAmount   sql.NullFloat64 `gorm:"column:JFOutAmount" json:"JFOutAmount"`
	JFOutCost     sql.NullFloat64 `gorm:"column:JFOutCost" json:"JFOutCost"`
	JFOutQty      sql.NullFloat64 `gorm:"column:JFOutQty" json:"JFOutQty"`
	JPOutAmount   sql.NullFloat64 `gorm:"column:JPOutAmount" json:"JPOutAmount"`
	JPOutCost     sql.NullFloat64 `gorm:"column:JPOutCost" json:"JPOutCost"`
	JPOutQty      sql.NullFloat64 `gorm:"column:JPOutQty" json:"JPOutQty"`
	JYOutAmount   sql.NullFloat64 `gorm:"column:JYOutAmount" json:"JYOutAmount"`
	JYOutCost     sql.NullFloat64 `gorm:"column:JYOutCost" json:"JYOutCost"`
	JYOutQty      sql.NullFloat64 `gorm:"column:JYOutQty" json:"JYOutQty"`
	LossCost      sql.NullFloat64 `gorm:"column:LossCost" json:"LossCost"`
	LossQty       sql.NullFloat64 `gorm:"column:LossQty" json:"LossQty"`
	OutAmount     sql.NullFloat64 `gorm:"column:OutAmount" json:"OutAmount"`
	OutCost       sql.NullFloat64 `gorm:"column:OutCost" json:"OutCost"`
	OutQty        sql.NullFloat64 `gorm:"column:OutQty" json:"OutQty"`
	PrevEndCost   sql.NullFloat64 `gorm:"column:PrevEndCost" json:"PrevEndCost"`
	PrevEndQty    sql.NullFloat64 `gorm:"column:PrevEndQty" json:"PrevEndQty"`
	ProfitCost    sql.NullFloat64 `gorm:"column:ProfitCost" json:"ProfitCost"`
	ProfitQty     sql.NullFloat64 `gorm:"column:ProfitQty" json:"ProfitQty"`
	QCOutAmount   sql.NullFloat64 `gorm:"column:QCOutAmount" json:"QCOutAmount"`
	QCOutCost     sql.NullFloat64 `gorm:"column:QCOutCost" json:"QCOutCost"`
	QCOutQty      sql.NullFloat64 `gorm:"column:QCOutQty" json:"QCOutQty"`
	QGOutAmount   sql.NullFloat64 `gorm:"column:QGOutAmount" json:"QGOutAmount"`
	QGOutCost     sql.NullFloat64 `gorm:"column:QGOutCost" json:"QGOutCost"`
	QGOutQty      sql.NullFloat64 `gorm:"column:QGOutQty" json:"QGOutQty"`
	QJInCost      sql.NullFloat64 `gorm:"column:QJInCost" json:"QJInCost"`
	QJInQty       sql.NullFloat64 `gorm:"column:QJInQty" json:"QJInQty"`
	QSInCost      sql.NullFloat64 `gorm:"column:QSInCost" json:"QSInCost"`
	QSInQty       sql.NullFloat64 `gorm:"column:QSInQty" json:"QSInQty"`
	QTInCost      sql.NullFloat64 `gorm:"column:QTInCost" json:"QTInCost"`
	QTInQty       sql.NullFloat64 `gorm:"column:QTInQty" json:"QTInQty"`
	QTOutAmount   sql.NullFloat64 `gorm:"column:QTOutAmount" json:"QTOutAmount"`
	QTOutCost     sql.NullFloat64 `gorm:"column:QTOutCost" json:"QTOutCost"`
	QTOutQty      sql.NullFloat64 `gorm:"column:QTOutQty" json:"QTOutQty"`
	ReviseCost    sql.NullFloat64 `gorm:"column:ReviseCost" json:"ReviseCost"`
	StockACost    sql.NullFloat64 `gorm:"column:StockACost" json:"StockACost"`
	StockAInCost  sql.NullFloat64 `gorm:"column:StockAInCost" json:"StockAInCost"`
	StockAInQty   sql.NullFloat64 `gorm:"column:StockAInQty" json:"StockAInQty"`
	StockAOutCost sql.NullFloat64 `gorm:"column:StockAOutCost" json:"StockAOutCost"`
	StockAOutQty  sql.NullFloat64 `gorm:"column:StockAOutQty" json:"StockAOutQty"`
	StockAQty     sql.NullFloat64 `gorm:"column:StockAQty" json:"StockAQty"`
	StockSCost    sql.NullFloat64 `gorm:"column:StockSCost" json:"StockSCost"`
	StockSInCost  sql.NullFloat64 `gorm:"column:StockSInCost" json:"StockSInCost"`
	StockSInQty   sql.NullFloat64 `gorm:"column:StockSInQty" json:"StockSInQty"`
	StockSOutCost sql.NullFloat64 `gorm:"column:StockSOutCost" json:"StockSOutCost"`
	StockSOutQty  sql.NullFloat64 `gorm:"column:StockSOutQty" json:"StockSOutQty"`
	StockSQty     sql.NullFloat64 `gorm:"column:StockSQty" json:"StockSQty"`
	XSOutAmount   sql.NullFloat64 `gorm:"column:XSOutAmount" json:"XSOutAmount"`
	XSOutCost     sql.NullFloat64 `gorm:"column:XSOutCost" json:"XSOutCost"`
	XSOutQty      sql.NullFloat64 `gorm:"column:XSOutQty" json:"XSOutQty"`
	YLOutAmount   sql.NullFloat64 `gorm:"column:YLOutAmount" json:"YLOutAmount"`
	YLOutCost     sql.NullFloat64 `gorm:"column:YLOutCost" json:"YLOutCost"`
	YLOutQty      sql.NullFloat64 `gorm:"column:YLOutQty" json:"YLOutQty"`
	YearMonth     string          `gorm:"column:YearMonth" json:"YearMonth"`
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode      sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PMonthbalanced) TableName() string {
	return "p_monthbalanced"
}
