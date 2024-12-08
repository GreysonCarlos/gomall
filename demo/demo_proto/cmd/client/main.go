package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GreysonCarlos/demo/demo_proto/kitex_gen/pbapi"
	"github.com/GreysonCarlos/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r), client.WithRPCTimeout(time.Second*3))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", resp)
	time.Sleep(time.Second)
}
