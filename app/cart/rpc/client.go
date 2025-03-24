package rpc

import (
	"sync"

	"github.com/GreysonCarlos/gomall/app/cart/conf"
	cartUtils "github.com/GreysonCarlos/gomall/app/cart/utils"
	"github.com/GreysonCarlos/gomall/common/clientsuite"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product/productcatalogservie"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient 		userservice.Client
	ProductClient	productcatalogservie.Client
	once 			sync.Once
	ServiceName 	= conf.GetConf().Kitex.Service
	RegistryAddr 	= conf.GetConf().Registry.RegistryAddress[0]
	err				error
)

func Init() {
	// 保证只初始化一次
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: 	ServiceName,
			RegistryAddr: 			RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservie.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}