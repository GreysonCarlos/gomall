package service

import (
	"context"
	"time"

	order "github.com/GreysonCarlos/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/GreysonCarlos/gomall/app/frontend/infra/rpc"
	"github.com/GreysonCarlos/gomall/app/frontend/types"
	frontendUtils "github.com/GreysonCarlos/gomall/app/frontend/utils"
	rpcorder "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *order.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	var list []types.Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		for _, u := range v.Items {
			total += u.Cost
			i := u.Item
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: i.ProductId})
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}
			items = append(items, types.OrderItem{
				// 通过rpc调用来获取商品名以及相关信息
				ProductName: productResp.Product.Name,
				Picture: productResp.Product.Picture,
				Qty: i.Quantity,
				Cost: total,
			})
			
		}
		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId: v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:01"),
			Cost: total,
			Items: items,
		})
	}
	return utils.H{
		"title": "Order",
		"orders": list,
	}, nil
}
