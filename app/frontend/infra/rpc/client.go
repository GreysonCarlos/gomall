package rpc

import (
	"sync"

	"github.com/GreysonCarlos/gomall/app/frontend/conf"
	frontendUtils "github.com/GreysonCarlos/gomall/app/frontend/utils"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product/productcatalogservie"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/GreysonCarlos/gomall/common/clientsuite"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient 		userservice.Client
	ProductClient	productcatalogservie.Client
	CartClient		cartservice.Client
	CheckoutClient	checkoutservice.Client
	OrderClient		orderservice.Client
	ServiceName 	= frontendUtils.ServiceName
	MetricsPort 	= conf.GetConf().Hertz.MetricsPort
	RegistryAddr 	= conf.GetConf().Hertz.RegistryAddr
	once 			sync.Once
	err				error
)

func Init() {
	// 保证只初始化一次
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: 	ServiceName,
			RegistryAddr: 			RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservie.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: 	ServiceName,
			RegistryAddr: 			RegistryAddr,
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: 	ServiceName,
			RegistryAddr: 			RegistryAddr,
		}),
	}

	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: 	ServiceName,
			RegistryAddr: 			RegistryAddr,
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleError(err)
}