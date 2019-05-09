package models

import "database/sql"

type ViewPerformance struct {
	COMNo             string         `gorm:"column:COMNo" json:"COMNo"`
	Date              sql.NullString `gorm:"column:Date" json:"Date"`
	Branch            string         `gorm:"column:branch" json:"branch"`
	Cid               int            `gorm:"column:cid" json:"cid"`
	Company           string         `gorm:"column:company" json:"company"`
	MemSalesAmount    []byte         `gorm:"column:memSalesAmount" json:"memSalesAmount"`
	MemSalesRoyalty   []byte         `gorm:"column:memSalesRoyalty" json:"memSalesRoyalty"`
	MemSaleskpi       []byte         `gorm:"column:memSaleskpi" json:"memSaleskpi"`
	Name              sql.NullString `gorm:"column:name" json:"name"`
	PartSalesAmount   []byte         `gorm:"column:partSalesAmount" json:"partSalesAmount"`
	PartSalesRoyalty  []byte         `gorm:"column:partSalesRoyalty" json:"partSalesRoyalty"`
	PartSaleskpi      []byte         `gorm:"column:partSaleskpi" json:"partSaleskpi"`
	PartWorkerAmount  []byte         `gorm:"column:partWorkerAmount" json:"partWorkerAmount"`
	PartWorkerRoyalty []byte         `gorm:"column:partWorkerRoyalty" json:"partWorkerRoyalty"`
	PartWorkerkpi     []byte         `gorm:"column:partWorkerkpi" json:"partWorkerkpi"`
	SerRecAmount      []byte         `gorm:"column:serRecAmount" json:"serRecAmount"`
	SerRecRoyalty     []byte         `gorm:"column:serRecRoyalty" json:"serRecRoyalty"`
	SerReckpi         []byte         `gorm:"column:serReckpi" json:"serReckpi"`
	Team              sql.NullString `gorm:"column:team" json:"team"`
}

// TableName sets the insert table name for this struct type
func (v *ViewPerformance) TableName() string {
	return "view_performance"
}
