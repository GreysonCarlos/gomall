package service

import (
	"context"

	home "github.com/GreysonCarlos/projects/Gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type Method1Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMethod1Service(Context context.Context, RequestContext *app.RequestContext) *Method1Service {
	return &Method1Service{RequestContext: RequestContext, Context: Context}
}

func (h *Method1Service) Run(req *home.Empty) (resp *home.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	
	return
}
