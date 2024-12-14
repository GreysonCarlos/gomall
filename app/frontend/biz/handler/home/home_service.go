package home

import (
	"context"

	"github.com/GreysonCarlos/projects/Gomall/app/frontend/biz/service"
	"github.com/GreysonCarlos/projects/Gomall/app/frontend/biz/utils"
	home "github.com/GreysonCarlos/projects/Gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Method1 .
// @router / [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMethod1Service(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// 加载特定模板
	c.HTML(consts.StatusOK, "home", resp)
}
