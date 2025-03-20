package rpc

import (
	"sync"

	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product/productcatalogservie"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/GreysonCarlos/gomall/app/cart/conf"
	cartUtils "github.com/GreysonCarlos/gomall/app/cart/utils"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient 		userservice.Client
	ProductClient	productcatalogservie.Client
	once sync.Once
)

func Init() {
	// 保证只初始化一次
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservie.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}