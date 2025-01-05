package service

import (
	"context"

	"github.com/GreysonCarlos/gomall/app/product/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/product/biz/model"
	product "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CatagoryName)
	resp = &product.ListProductsResp{}
	// 遍历分类列表,再遍历分类包含的商品
	for _, v1 := range c{
		for _, v := range v1.Products{
			resp.Products = append(resp.Products, &product.Product{
				Id: uint32(v.ID),
				Price: v.Price,
				Name: v.Name,
				Description: v.Description,
				Picture: v.Picture,
			})
		}
	}
	return resp, nil
}
