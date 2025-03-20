package service

import (
	"context"

	"github.com/GreysonCarlos/gomall/app/order/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/order/biz/model"
	order "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(50003, "items is empty")
		return
	}
	// 事务开启
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId: orderId.String(),
			UserId: uint(req.UserId),
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}
		// 写入订单表
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		// 写入订单商品
		var items []model.OrderItem
		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId: v.Item.ProductId,
				Quantity: v.Item.Quantity,
				Cost: float64(v.Cost),
			})
		}

		// 上传订单信息
		if err := tx.Create(items).Error; err != nil {
			return err
		}

		// 处理响应数据,传递订单
		resp = &order.PlaceOrderResp{
			Result: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	return
}
