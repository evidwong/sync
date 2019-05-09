package models

import (
	"database/sql"
	"time"
)

type TVehicle struct {
	TCIEndTime                       time.Time       `gorm:"column:TCI_end_time" json:"TCI_end_time"`
	VCIAmount                        sql.NullFloat64 `gorm:"column:VCI_amount" json:"VCI_amount"`
	VCIEndTime                       time.Time       `gorm:"column:VCI_end_time" json:"VCI_end_time"`
	VVTEndTime                       time.Time       `gorm:"column:VVT_end_time" json:"VVT_end_time"`
	Allnature                        sql.NullString  `gorm:"column:allnature" json:"allnature"`
	AnnualticketTime                 time.Time       `gorm:"column:annualticket_time" json:"annualticket_time"`
	AreaID                           sql.NullInt64   `gorm:"column:area_id" json:"area_id"`
	CarBysc                          sql.NullString  `gorm:"column:car_bysc" json:"car_bysc"`
	CarDjzs                          sql.NullString  `gorm:"column:car_djzs" json:"car_djzs"`
	CarInfo                          sql.NullString  `gorm:"column:car_info" json:"car_info"`
	CarInvoice                       sql.NullString  `gorm:"column:car_invoice" json:"car_invoice"`
	CarKey                           sql.NullInt64   `gorm:"column:car_key" json:"car_key"`
	CarWszm                          sql.NullString  `gorm:"column:car_wszm" json:"car_wszm"`
	CarXsz                           sql.NullString  `gorm:"column:car_xsz" json:"car_xsz"`
	CarYhsc                          sql.NullString  `gorm:"column:car_yhsc" json:"car_yhsc"`
	Carattr                          sql.NullString  `gorm:"column:carattr" json:"carattr"`
	Cartype                          sql.NullString  `gorm:"column:cartype" json:"cartype"`
	Cid                              sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ClueID                           sql.NullInt64   `gorm:"column:clue_id" json:"clue_id"`
	CreateBy                         sql.NullInt64   `gorm:"column:create_by" json:"create_by"`
	CreateTime                       time.Time       `gorm:"column:create_time" json:"create_time"`
	CustomerID                       sql.NullInt64   `gorm:"column:customer_id" json:"customer_id"`
	DealCreateTime                   time.Time       `gorm:"column:deal_create_time" json:"deal_create_time"`
	EmissionStandard                 sql.NullString  `gorm:"column:emission_standard" json:"emission_standard"`
	EmptyLicheng                     sql.NullFloat64 `gorm:"column:empty_licheng" json:"empty_licheng"`
	ExteriorState                    sql.NullString  `gorm:"column:exterior_state" json:"exterior_state"`
	FaultRemark                      sql.NullString  `gorm:"column:fault_remark" json:"fault_remark"`
	Guohudidian                      sql.NullString  `gorm:"column:guohudidian" json:"guohudidian"`
	HandTime                         time.Time       `gorm:"column:hand_time" json:"hand_time"`
	ID                               int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	InquiryHighestBuyer              sql.NullString  `gorm:"column:inquiry_highest_buyer" json:"inquiry_highest_buyer"`
	InquiryHighestBuyerID            sql.NullInt64   `gorm:"column:inquiry_highest_buyer_id" json:"inquiry_highest_buyer_id"`
	InquiryHighestBuyerPhonenum      sql.NullString  `gorm:"column:inquiry_highest_buyer_phonenum" json:"inquiry_highest_buyer_phonenum"`
	InquiryHighestPrice              sql.NullFloat64 `gorm:"column:inquiry_highest_price" json:"inquiry_highest_price"`
	InquiryPersonCount               sql.NullInt64   `gorm:"column:inquiry_person_count" json:"inquiry_person_count"`
	InspectionBeginTime              time.Time       `gorm:"column:inspection_begin_time" json:"inspection_begin_time"`
	InspectionEndTime                time.Time       `gorm:"column:inspection_end_time" json:"inspection_end_time"`
	InspectionOrder                  sql.NullString  `gorm:"column:inspection_order" json:"inspection_order"`
	InspectionTime                   time.Time       `gorm:"column:inspection_time" json:"inspection_time"`
	IsDisabled                       sql.NullInt64   `gorm:"column:is_disabled" json:"is_disabled"`
	IsFault                          sql.NullInt64   `gorm:"column:is_fault" json:"is_fault"`
	IsFirstTrade                     sql.NullInt64   `gorm:"column:is_first_trade" json:"is_first_trade"`
	IsShowlicense                    sql.NullInt64   `gorm:"column:is_showlicense" json:"is_showlicense"`
	IsStorage                        sql.NullString  `gorm:"column:is_storage" json:"is_storage"`
	LastEditBy                       sql.NullInt64   `gorm:"column:last_edit_by" json:"last_edit_by"`
	LastEditTime                     time.Time       `gorm:"column:last_edit_time" json:"last_edit_time"`
	Minaddprice                      sql.NullFloat64 `gorm:"column:minaddprice" json:"minaddprice"`
	OrderApprovalState               sql.NullString  `gorm:"column:order_approval_state" json:"order_approval_state"`
	OrderBuyerName                   sql.NullString  `gorm:"column:order_buyer_name" json:"order_buyer_name"`
	OrderCode                        sql.NullString  `gorm:"column:order_code" json:"order_code"`
	OrderCreateTime                  time.Time       `gorm:"column:order_create_time" json:"order_create_time"`
	OrderDeposit                     sql.NullFloat64 `gorm:"column:order_deposit" json:"order_deposit"`
	OrderLinkman                     sql.NullString  `gorm:"column:order_linkman" json:"order_linkman"`
	OrderLinkmanTelephone            sql.NullString  `gorm:"column:order_linkman_telephone" json:"order_linkman_telephone"`
	OrderMoney                       sql.NullFloat64 `gorm:"column:order_money" json:"order_money"`
	OrderPaymentMoney                sql.NullFloat64 `gorm:"column:order_payment_money" json:"order_payment_money"`
	OrderState                       sql.NullString  `gorm:"column:order_state" json:"order_state"`
	OrderType                        sql.NullString  `gorm:"column:order_type" json:"order_type"`
	OrderWechatInformation           sql.NullString  `gorm:"column:order_wechat_information" json:"order_wechat_information"`
	OrgID                            sql.NullInt64   `gorm:"column:org_id" json:"org_id"`
	ParentID                         sql.NullInt64   `gorm:"column:parent_id" json:"parent_id"`
	PictureLargeURL                  sql.NullString  `gorm:"column:picture_large_url" json:"picture_large_url"`
	PictureMidURL                    sql.NullString  `gorm:"column:picture_mid_url" json:"picture_mid_url"`
	PictureSmallURL                  sql.NullString  `gorm:"column:picture_small_url" json:"picture_small_url"`
	PictureURL                       sql.NullString  `gorm:"column:picture_url" json:"picture_url"`
	ProductionDate                   time.Time       `gorm:"column:production_date" json:"production_date"`
	Profit                           sql.NullFloat64 `gorm:"column:profit" json:"profit"`
	QuickInquiryHighestBuyer         sql.NullString  `gorm:"column:quick_inquiry_highest_buyer" json:"quick_inquiry_highest_buyer"`
	QuickInquiryHighestBuyerPhonenum sql.NullString  `gorm:"column:quick_inquiry_highest_buyer_phonenum" json:"quick_inquiry_highest_buyer_phonenum"`
	QuickInquiryHighestPrice         sql.NullFloat64 `gorm:"column:quick_inquiry_highest_price" json:"quick_inquiry_highest_price"`
	RealLicheng                      sql.NullFloat64 `gorm:"column:real_licheng" json:"real_licheng"`
	Remark                           sql.NullString  `gorm:"column:remark" json:"remark"`
	Resource1                        sql.NullString  `gorm:"column:resource1" json:"resource1"`
	Resource2                        sql.NullString  `gorm:"column:resource2" json:"resource2"`
	Resource3                        sql.NullString  `gorm:"column:resource3" json:"resource3"`
	ShopID                           sql.NullInt64   `gorm:"column:shop_id" json:"shop_id"`
	SourceOrgName                    sql.NullString  `gorm:"column:source_org_name" json:"source_org_name"`
	SourceShopName                   sql.NullString  `gorm:"column:source_shop_name" json:"source_shop_name"`
	TradeTimes                       sql.NullInt64   `gorm:"column:trade_times" json:"trade_times"`
	TrimColor                        sql.NullString  `gorm:"column:trim_color" json:"trim_color"`
	TrimState                        sql.NullString  `gorm:"column:trim_state" json:"trim_state"`
	Usage                            sql.NullString  `gorm:"column:usage" json:"usage"`
	VehicleColor                     sql.NullString  `gorm:"column:vehicle_color" json:"vehicle_color"`
	VehicleDealContractPath          sql.NullString  `gorm:"column:vehicle_deal_contract_path" json:"vehicle_deal_contract_path"`
	VehicleDealID                    sql.NullInt64   `gorm:"column:vehicle_deal_id" json:"vehicle_deal_id"`
	VehicleDealIsOffset              sql.NullInt64   `gorm:"column:vehicle_deal_is_offset" json:"vehicle_deal_is_offset"`
	VehicleDealPrice                 sql.NullFloat64 `gorm:"column:vehicle_deal_price" json:"vehicle_deal_price"`
	VehicleEngineNumber              sql.NullString  `gorm:"column:vehicle_engine_number" json:"vehicle_engine_number"`
	VehicleLisenceNumber             sql.NullString  `gorm:"column:vehicle_lisence_number" json:"vehicle_lisence_number"`
	VehicleLisenceVin                sql.NullString  `gorm:"column:vehicle_lisence_vin" json:"vehicle_lisence_vin"`
	VehicleLocation                  sql.NullString  `gorm:"column:vehicle_location" json:"vehicle_location"`
	VehicleModelID                   sql.NullInt64   `gorm:"column:vehicle_model_id" json:"vehicle_model_id"`
	VehicleName                      sql.NullString  `gorm:"column:vehicle_name" json:"vehicle_name"`
	VehicleRegisteDate               time.Time       `gorm:"column:vehicle_registe_date" json:"vehicle_registe_date"`
	VehicleState                     sql.NullString  `gorm:"column:vehicle_state" json:"vehicle_state"`
	Weizhangds                       sql.NullInt64   `gorm:"column:weizhangds" json:"weizhangds"`
}

// TableName sets the insert table name for this struct type
func (t *TVehicle) TableName() string {
	return "t_vehicle"
}
