package main

import (
	"context"
	payment "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/payment"
	"github.com/GreysonCarlos/gomall/app/payment/biz/service"
)

// PaymenServiceImpl implements the last service interface defined in the IDL.
type PaymenServiceImpl struct{}

// Charge implements the PaymenServiceImpl interface.
func (s *PaymenServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
