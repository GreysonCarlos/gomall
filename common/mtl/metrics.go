package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info){
	Registry = prometheus.NewRegistry()
	// 注册go运行时相关指标
	Registry.MustRegister(collectors.NewGoCollector())
	// 注册进程相关指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	r, _ := consul.NewConsulRegister(registryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)	// 将metricsPort转换为tcp地址
	// 构造注册信息
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr: addr,
		Weight: 1,
		Tags: map[string]string{"service": serviceName},
	}

	_ = r.Register(registryInfo)
	// 注册shutdow hook，在服务关闭时下线注册信息
	server.RegisterShutdownHook(func() {
		_ = r.Deregister(registryInfo)
	})
	// 启动metrics server
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	// 异步启动server供Prometheus拉取指标
	go func () {
		_ = http.ListenAndServe(metricsPort, nil)
	}()
	return r, registryInfo
}