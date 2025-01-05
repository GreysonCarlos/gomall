package service

import (
	"context"

	common "github.com/GreysonCarlos/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/GreysonCarlos/gomall/app/frontend/infra/rpc"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type Method1Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMethod1Service(Context context.Context, RequestContext *app.RequestContext) *Method1Service {
	return &Method1Service{RequestContext: RequestContext, Context: Context}
}

func (h *Method1Service) Run(req *common.Empty) (res map[string]any, err error) {
	// todo edit your code
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
