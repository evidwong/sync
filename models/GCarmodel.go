package models

import (
	"database/sql"
	"time"
)

type GCarmodel struct {
	AcceTime          sql.NullString  `gorm:"column:AcceTime" json:"AcceTime"`
	AcqTax            sql.NullFloat64 `gorm:"column:AcqTax" json:"AcqTax"`
	AdjustingArm      sql.NullString  `gorm:"column:AdjustingArm" json:"AdjustingArm"`
	BackHang          sql.NullString  `gorm:"column:BackHang" json:"BackHang"`
	BackTread         sql.NullString  `gorm:"column:BackTread" json:"BackTread"`
	Brake             sql.NullString  `gorm:"column:Brake" json:"Brake"`
	Brand             sql.NullString  `gorm:"column:Brand" json:"Brand"`
	CarBoxLength      sql.NullString  `gorm:"column:CarBoxLength" json:"CarBoxLength"`
	CarCOMNo          sql.NullString  `gorm:"column:CarCOMNo" json:"CarCOMNo"`
	CarFld1           sql.NullString  `gorm:"column:CarFld1" json:"CarFld1"`
	CarFld10          sql.NullString  `gorm:"column:CarFld10" json:"CarFld10"`
	CarFld11          sql.NullString  `gorm:"column:CarFld11" json:"CarFld11"`
	CarFld12          sql.NullString  `gorm:"column:CarFld12" json:"CarFld12"`
	CarFld13          sql.NullString  `gorm:"column:CarFld13" json:"CarFld13"`
	CarFld14          sql.NullString  `gorm:"column:CarFld14" json:"CarFld14"`
	CarFld15          sql.NullString  `gorm:"column:CarFld15" json:"CarFld15"`
	CarFld16          sql.NullString  `gorm:"column:CarFld16" json:"CarFld16"`
	CarFld17          sql.NullString  `gorm:"column:CarFld17" json:"CarFld17"`
	CarFld18          sql.NullString  `gorm:"column:CarFld18" json:"CarFld18"`
	CarFld19          sql.NullString  `gorm:"column:CarFld19" json:"CarFld19"`
	CarFld2           sql.NullString  `gorm:"column:CarFld2" json:"CarFld2"`
	CarFld20          sql.NullString  `gorm:"column:CarFld20" json:"CarFld20"`
	CarFld3           sql.NullString  `gorm:"column:CarFld3" json:"CarFld3"`
	CarFld4           sql.NullString  `gorm:"column:CarFld4" json:"CarFld4"`
	CarFld5           sql.NullString  `gorm:"column:CarFld5" json:"CarFld5"`
	CarFld6           sql.NullString  `gorm:"column:CarFld6" json:"CarFld6"`
	CarFld7           sql.NullString  `gorm:"column:CarFld7" json:"CarFld7"`
	CarFld8           sql.NullString  `gorm:"column:CarFld8" json:"CarFld8"`
	CarFld9           sql.NullString  `gorm:"column:CarFld9" json:"CarFld9"`
	CarLength         sql.NullString  `gorm:"column:CarLength" json:"CarLength"`
	CarLevel          sql.NullString  `gorm:"column:CarLevel" json:"CarLevel"`
	CarOutPut         sql.NullString  `gorm:"column:CarOutPut" json:"CarOutPut"`
	CarRefPrice1      sql.NullFloat64 `gorm:"column:CarRefPrice1" json:"CarRefPrice1"`
	CarRefPrice2      sql.NullFloat64 `gorm:"column:CarRefPrice2" json:"CarRefPrice2"`
	CarRefPrice3      sql.NullFloat64 `gorm:"column:CarRefPrice3" json:"CarRefPrice3"`
	CarRefPrice4      sql.NullFloat64 `gorm:"column:CarRefPrice4" json:"CarRefPrice4"`
	CarRefPrice5      sql.NullFloat64 `gorm:"column:CarRefPrice5" json:"CarRefPrice5"`
	CarRefPrice6      sql.NullFloat64 `gorm:"column:CarRefPrice6" json:"CarRefPrice6"`
	CarSize           sql.NullString  `gorm:"column:CarSize" json:"CarSize"`
	CarType           sql.NullString  `gorm:"column:CarType" json:"CarType"`
	CargoCapacity     sql.NullString  `gorm:"column:CargoCapacity" json:"CargoCapacity"`
	Clutch            sql.NullString  `gorm:"column:Clutch" json:"Clutch"`
	CostPrice         sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DirectMachine     sql.NullString  `gorm:"column:DirectMachine" json:"DirectMachine"`
	Displacement      sql.NullString  `gorm:"column:Displacement" json:"Displacement"`
	DriveType         sql.NullString  `gorm:"column:DriveType" json:"DriveType"`
	DrivingCab        sql.NullString  `gorm:"column:DrivingCab" json:"DrivingCab"`
	EmissionStd       sql.NullString  `gorm:"column:EmissionStd" json:"EmissionStd"`
	Engine            sql.NullString  `gorm:"column:Engine" json:"Engine"`
	EngineType        sql.NullString  `gorm:"column:EngineType" json:"EngineType"`
	Environment       sql.NullString  `gorm:"column:Environment" json:"Environment"`
	Equip             sql.NullString  `gorm:"column:Equip" json:"Equip"`
	EquipQuality      sql.NullString  `gorm:"column:EquipQuality" json:"EquipQuality"`
	Factory           sql.NullString  `gorm:"column:Factory" json:"Factory"`
	FrontAxle         sql.NullString  `gorm:"column:FrontAxle" json:"FrontAxle"`
	FrontHang         sql.NullString  `gorm:"column:FrontHang" json:"FrontHang"`
	FrontTread        sql.NullString  `gorm:"column:FrontTread" json:"FrontTread"`
	Fuel              sql.NullString  `gorm:"column:Fuel" json:"Fuel"`
	HModel            sql.NullString  `gorm:"column:HModel" json:"HModel"`
	HorsePower        sql.NullString  `gorm:"column:HorsePower" json:"HorsePower"`
	Innern            sql.NullString  `gorm:"column:Innern" json:"Innern"`
	MaxOutput         sql.NullString  `gorm:"column:MaxOutput" json:"MaxOutput"`
	MaxQuality        sql.NullString  `gorm:"column:MaxQuality" json:"MaxQuality"`
	MaxSpeed          sql.NullString  `gorm:"column:MaxSpeed" json:"MaxSpeed"`
	MaxTorque         sql.NullString  `gorm:"column:MaxTorque" json:"MaxTorque"`
	MinPrice          sql.NullFloat64 `gorm:"column:MinPrice" json:"MinPrice"`
	ModelCode         sql.NullString  `gorm:"column:ModelCode" json:"ModelCode"`
	ModifyDate        time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	NoTaxCostPrice    sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OilCapacity       sql.NullString  `gorm:"column:OilCapacity" json:"OilCapacity"`
	OilUse            sql.NullString  `gorm:"column:OilUse" json:"OilUse"`
	OutFactoryPrice   sql.NullFloat64 `gorm:"column:OutFactoryPrice" json:"OutFactoryPrice"`
	PowerTakeOff      sql.NullString  `gorm:"column:PowerTakeOff" json:"PowerTakeOff"`
	Price             sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Rangerover        sql.NullString  `gorm:"column:Rangerover" json:"Rangerover"`
	RefCostPrice      sql.NullFloat64 `gorm:"column:RefCostPrice" json:"RefCostPrice"`
	Remarks           sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	Saddle            sql.NullString  `gorm:"column:Saddle" json:"Saddle"`
	Seat              sql.NullInt64   `gorm:"column:Seat" json:"Seat"`
	SpeedRatio        sql.NullString  `gorm:"column:SpeedRatio" json:"SpeedRatio"`
	TireCount         sql.NullInt64   `gorm:"column:TireCount" json:"TireCount"`
	TireType          sql.NullString  `gorm:"column:TireType" json:"TireType"`
	TransAxle         sql.NullString  `gorm:"column:TransAxle" json:"TransAxle"`
	TransferPrice     sql.NullFloat64 `gorm:"column:TransferPrice" json:"TransferPrice"`
	Turning           sql.NullString  `gorm:"column:Turning" json:"Turning"`
	UserDefCode       sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	VehicleName       sql.NullString  `gorm:"column:VehicleName" json:"VehicleName"`
	Wheelbase         sql.NullString  `gorm:"column:Wheelbase" json:"Wheelbase"`
	YearStyle         sql.NullString  `gorm:"column:YearStyle" json:"YearStyle"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno             sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode         sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	StoreID           sql.NullString  `gorm:"column:store_id" json:"store_id"`
	StoreName         sql.NullString  `gorm:"column:store_name" json:"store_name"`
}

// TableName sets the insert table name for this struct type
func (g *GCarmodel) TableName() string {
	return "g_carmodel"
}
