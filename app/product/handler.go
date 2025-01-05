package main

import (
	"context"
	product "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product"
	"github.com/GreysonCarlos/gomall/app/product/biz/service"
)

// ProductCatalogServieImpl implements the last service interface defined in the IDL.
type ProductCatalogServieImpl struct{}

// ListProducts implements the ProductCatalogServieImpl interface.
func (s *ProductCatalogServieImpl) ListProducts(ctx context.Context, req *product.ListProductReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServieImpl interface.
func (s *ProductCatalogServieImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServieImpl interface.
func (s *ProductCatalogServieImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}
