package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	// 记录商品信息
	ProductId	uint32 `gorm:"type:int(11)"`
	OrderIdRefer string	`gorm:"type:varchar(100);index"`	// 依赖主表id
	Quantity uint32 `gorm:"type:int(11)"`
	Cost float64 `gorm:"type:decimal(10,2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}