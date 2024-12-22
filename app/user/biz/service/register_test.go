package service

import (
	"context"
	"testing"

	"github.com/GreysonCarlos/gomall/app/user/biz/dal/mysql"
	user "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email: "1demo@admin.com",
		Password: "jadoiajf",
		PasswordConfirm: "jadoiajf",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
