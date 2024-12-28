package service

import (
	"context"

	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/user"
	auth "github.com/GreysonCarlos/projects/Gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/GreysonCarlos/projects/Gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/GreysonCarlos/projects/Gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email: 				req.Email,
		Password: 			req.Password,
		PasswordConfirm: 	req.PasswordConfirm,
	})

	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
