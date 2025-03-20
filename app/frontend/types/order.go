package types

type OrderItem struct {
	ProductName	string
	Picture		string
	Qty			uint32
	Cost		float32
}

// 创建订单结构体
type Order struct {
	OrderId	string
	CreateDate	string
	Cost	float32
	Items	[]OrderItem
}