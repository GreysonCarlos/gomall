package rpc

import (
	"context"
	"testing"

	user "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/user"
)


func Test_initUserClient(t *testing.T) {
	initUserClient()
	resp, err := UserClient.Register(context.Background(), &user.RegisterReq{
		Email:    "2demo@damin.com",
		Password: "Jfoajsfoji",
		PasswordConfirm: "Jfoajsfoji",
	})
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	t.Logf("resp: %v", resp)
}