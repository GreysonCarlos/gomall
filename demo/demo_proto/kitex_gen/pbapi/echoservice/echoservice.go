// Code generated by Kitex v0.9.1. DO NOT EDIT.

package echoservice

import (
	"context"
	"errors"
	pbapi "github.com/GreysonCarlos/demo/demo_proto/kitex_gen/pbapi"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Echo": kitex.NewMethodInfo(
		echoHandler,
		newEchoArgs,
		newEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	echoServiceServiceInfo                = NewServiceInfo()
	echoServiceServiceInfoForClient       = NewServiceInfoForClient()
	echoServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return echoServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return echoServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return echoServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "EchoService"
	handlerType := (*pbapi.EchoService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "pbapi",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(pbapi.Request)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(pbapi.EchoService).Echo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *EchoArgs:
		success, err := handler.(pbapi.EchoService).Echo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*EchoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newEchoArgs() interface{} {
	return &EchoArgs{}
}

func newEchoResult() interface{} {
	return &EchoResult{}
}

type EchoArgs struct {
	Req *pbapi.Request
}

func (p *EchoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(pbapi.Request)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *EchoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *EchoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *EchoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *EchoArgs) Unmarshal(in []byte) error {
	msg := new(pbapi.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var EchoArgs_Req_DEFAULT *pbapi.Request

func (p *EchoArgs) GetReq() *pbapi.Request {
	if !p.IsSetReq() {
		return EchoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EchoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EchoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type EchoResult struct {
	Success *pbapi.Response
}

var EchoResult_Success_DEFAULT *pbapi.Response

func (p *EchoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(pbapi.Response)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *EchoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *EchoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *EchoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *EchoResult) Unmarshal(in []byte) error {
	msg := new(pbapi.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EchoResult) GetSuccess() *pbapi.Response {
	if !p.IsSetSuccess() {
		return EchoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EchoResult) SetSuccess(x interface{}) {
	p.Success = x.(*pbapi.Response)
}

func (p *EchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, Req *pbapi.Request) (r *pbapi.Response, err error) {
	var _args EchoArgs
	_args.Req = Req
	var _result EchoResult
	if err = p.c.Call(ctx, "Echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
