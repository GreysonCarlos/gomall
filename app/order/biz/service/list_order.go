package service

import (
	"context"

	"github.com/GreysonCarlos/gomall/app/order/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/order/biz/model"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	// 组装数据
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity: oi.Quantity,
				},
				Cost: float32(oi.Cost),
			})
		}
		orders = append(orders, &order.Order{
			CreatedAt: int32(v.CreatedAt.Unix()),
			OrderId: v.OrderId,
			UserId: uint32(v.UserId),
			UserCurrency: v.UserCurrency,
			Email: v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City: v.Consignee.City,
				State: v.Consignee.State,
				Country: v.Consignee.Country,
				ZipCode: v.Consignee.ZipCode,
			},
			Items: items,
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
		
	}
	return
}
