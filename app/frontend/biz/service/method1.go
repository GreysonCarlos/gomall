package service

import (
	"context"

	common "github.com/GreysonCarlos/projects/Gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type Method1Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMethod1Service(Context context.Context, RequestContext *app.RequestContext) *Method1Service {
	return &Method1Service{RequestContext: RequestContext, Context: Context}
}

func (h *Method1Service) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-2.jpeg"},
		{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-3.jpeg"},
		{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/t-shirt-4.jpeg"},
		{"Name": "T-shirt-5", "Price": 140, "Picture": "/static/image/t-shirt-5.jpeg"},
		{"Name": "T-shirt-6", "Price": 150, "Picture": "/static/image/t-shirt-6.jpeg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Item"] = items
	return resp, nil
}
