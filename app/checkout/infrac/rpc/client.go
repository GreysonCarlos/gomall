package rpc

import (
	"sync"

	"github.com/GreysonCarlos/gomall/app/checkout/conf"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product/productcatalogservie"
	"github.com/cloudwego/kitex/client"
	"github.com/GreysonCarlos/gomall/common/clientsuite"
)

var (
	CartClient	cartservice.Client
	ProductClient	productcatalogservie.Client
	PaymentClient	paymentservice.Client
	OrderClient		orderservice.Client
	once	sync.Once
	ServiceName 	= conf.GetConf().Kitex.Service
	RegistryAddr 	= conf.GetConf().Registry.RegistryAddress[0]
	err		error

)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initCartClient() {
	// 服务发现
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:		RegistryAddr,
		}),
	}
	
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	// 服务发现
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:		RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservie.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	// 服务发现
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:		RegistryAddr,
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}

func initOrderClient() {
	// 服务发现
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:		RegistryAddr,
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}