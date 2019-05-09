package models

import "database/sql"

type Goods struct {
	AddTime        sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Auditname      string         `gorm:"column:auditname" json:"auditname"`
	Audittime      string         `gorm:"column:audittime" json:"audittime"`
	BuyNum         sql.NullInt64  `gorm:"column:buy_num" json:"buy_num"`
	CanGive        int            `gorm:"column:can_give" json:"can_give"`
	CatID          sql.NullInt64  `gorm:"column:cat_id" json:"cat_id"`
	CategoryID     sql.NullString `gorm:"column:category_id" json:"category_id"`
	CategoryName   sql.NullString `gorm:"column:category_name" json:"category_name"`
	Cid            int            `gorm:"column:cid" json:"cid"`
	Comno          sql.NullString `gorm:"column:comno" json:"comno"`
	Delivery       sql.NullInt64  `gorm:"column:delivery" json:"delivery"`
	Detail         sql.NullString `gorm:"column:detail" json:"detail"`
	EndTime        int            `gorm:"column:end_time" json:"end_time"`
	ExpireDate     int            `gorm:"column:expire_date" json:"expire_date"`
	FavorableRate  float32        `gorm:"column:favorable_rate" json:"favorable_rate"`
	ID             int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Identify       sql.NullString `gorm:"column:identify" json:"identify"`
	Images         sql.NullString `gorm:"column:images" json:"images"`
	Integral       float32        `gorm:"column:integral" json:"integral"`
	IsConvert      int            `gorm:"column:is_convert" json:"is_convert"`
	IsCouponCenter int            `gorm:"column:is_coupon_center" json:"is_coupon_center"`
	IsNewsdiscount int            `gorm:"column:is_newsdiscount" json:"is_newsdiscount"`
	IsRatify       int            `gorm:"column:is_ratify" json:"is_ratify"`
	IsSales        int            `gorm:"column:is_sales" json:"is_sales"`
	Isaudit        int            `gorm:"column:isaudit" json:"isaudit"`
	Isdel          sql.NullInt64  `gorm:"column:isdel" json:"isdel"`
	MemberPrice    float32        `gorm:"column:member_price" json:"member_price"`
	Norm           sql.NullString `gorm:"column:norm" json:"norm"`
	OfferDesc      sql.NullString `gorm:"column:offer_desc" json:"offer_desc"`
	Pic            sql.NullString `gorm:"column:pic" json:"pic"`
	Price          float32        `gorm:"column:price" json:"price"`
	PriceMarket    float32        `gorm:"column:price_market" json:"price_market"`
	ReadNum        int            `gorm:"column:read_num" json:"read_num"`
	Recommond      int            `gorm:"column:recommond" json:"recommond"`
	Remark         sql.NullString `gorm:"column:remark" json:"remark"`
	SalesCount     int            `gorm:"column:sales_count" json:"sales_count"`
	Score          float32        `gorm:"column:score" json:"score"`
	SendNum        int            `gorm:"column:send_num" json:"send_num"`
	SendPercent    int            `gorm:"column:send_percent" json:"send_percent"`
	SharedNum      int            `gorm:"column:shared_num" json:"shared_num"`
	StartTime      int            `gorm:"column:start_time" json:"start_time"`
	Statement      sql.NullString `gorm:"column:statement" json:"statement"`
	Status         int            `gorm:"column:status" json:"status"`
	Stock          int            `gorm:"column:stock" json:"stock"`
	StoreID        sql.NullString `gorm:"column:store_id" json:"store_id"`
	StoreName      sql.NullString `gorm:"column:store_name" json:"store_name"`
	SuitableDesc   sql.NullString `gorm:"column:suitable_desc" json:"suitable_desc"`
	Thumb          sql.NullString `gorm:"column:thumb" json:"thumb"`
	Times          int            `gorm:"column:times" json:"times"`
	Title          string         `gorm:"column:title" json:"title"`
	Type           sql.NullInt64  `gorm:"column:type" json:"type"`
	UpdateTime     sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
	UseTimes       int            `gorm:"column:use_times" json:"use_times"`
	Voter          int            `gorm:"column:voter" json:"voter"`
}

// TableName sets the insert table name for this struct type
func (g *Goods) TableName() string {
	return "goods"
}
