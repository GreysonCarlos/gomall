package serversuite

import (
	"log"

	"github.com/GreysonCarlos/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// 添加Prometheus中间件,使用自定义mtl，将地址、端口直接置空
		server.WithTracer(prometheus.NewServerTracer("",
			"", 
			prometheus.WithDisableServer(true), 
			prometheus.WithRegistry(mtl.Registry)),
		),
	}

	// 服务注册抽象为公共部分
	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		log.Fatal(err)
	}

	opts = append(opts, server.WithRegistry(r))
	return opts
}