package rpc

import (
	"log"
	"sync"

	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	// 保证只初始化一次
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}
