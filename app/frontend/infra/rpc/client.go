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
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient 		userservice.Client
	ProductClient	productcatalogservie.Client
	CartClient		cartservice.Client
	CheckoutClient	checkoutservice.Client
	OrderClient		orderservice.Client
	once sync.Once
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
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservie.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleError(err)
}