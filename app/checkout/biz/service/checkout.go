package service

import (
	"context"

	"github.com/GreysonCarlos/gomall/app/checkout/infrac/mq"
	"github.com/GreysonCarlos/gomall/app/checkout/infrac/rpc"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/email"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/order"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/payment"
	"github.com/GreysonCarlos/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	// "github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}

	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var total float32
	var oi []*order.OrderItem
	for _, cartItem := range cartResult.Items {
		// 真实开发环境中避免在for循环内使用rpc调用，会影响性能
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity: cartItem.Quantity,
			},
			Cost: cost,
		})
	}
	var orderId string
	// 虚拟订单，模拟订单创建
	// u, _ := uuid.NewRandom()
	// orderId := u.String()

	// 更改:与orderservice进行交互，通过rpc调用获取orderId
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		UserCurrency: "USA",
		Email: req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City: req.Address.City,
			State: req.Address.State,
			Country: req.Address.Country,
			ZipCode: req.Address.ZipCode,
		},
		Items: oi,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}
	
	if orderResp != nil && orderResp.Result != nil {
		orderId = orderResp.Result.OrderId
	}
	payReq := &payment.ChargeReq{
		UserId: req.UserId,
		OrderId: orderId,
		Num: total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber: req.CreditCard.CreditCardNumber,
			CreditCardCvv: req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear: req.CreditCard.CreditCardExpirationYear,
		},
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)

	if err != nil {
		return nil, err
	}

	data, _ := proto.Marshal(&email.EmailReq{
		From: "from@example.com",
		To: req.Email,
		ContentType: "text/plain",
		Subject: "You have just created an order in the Phony Shop",
		Content: "You have just created an order in the Phony Shop",
	})
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	
	msg := &nats.Msg{Subject: "email", Data: data}

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentResult.TransactionId,
	}
	
	return
}
