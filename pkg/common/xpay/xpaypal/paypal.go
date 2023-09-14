package xpaypal

import (
	"fmt"
	"github.com/smartwalle/paypal"
	"lark/domain/po"
	"lark/pkg/conf"
	"log/slog"
)

var (
	pay *Paypal
)

type Paypal struct {
	client *paypal.Client
	cfg    *conf.Paypal
}

func NewPaypal(cfg *conf.Paypal) *Paypal {
	// 沙盒模式
	client := paypal.New(cfg.ClientID, cfg.Secret, false)
	pay = &Paypal{
		client: client,
		cfg:    cfg,
	}
	token, err := client.GetAccessToken()
	if err != nil {
		slog.Warn(err.Error())
	}
	slog.Info(token.AccessToken)

	return pay
}

func CreatePayment(order *po.Order) (result *paypal.Payment, err error) {
	var p = &paypal.Payment{}
	p.Intent = paypal.PaymentIntentSale
	p.Payer = &paypal.Payer{}
	p.Payer.PaymentMethod = paypal.PaymentMethodPayPal
	p.RedirectURLs = &paypal.RedirectURLs{}
	p.RedirectURLs.CancelURL = pay.cfg.Server + pay.cfg.CancelURL
	p.RedirectURLs.ReturnURL = pay.cfg.Server + pay.cfg.ReturnURL

	var transaction = &paypal.Transaction{}
	transaction.InvoiceNumber = order.OrderSn
	p.Transactions = []*paypal.Transaction{transaction}

	transaction.Amount = &paypal.Amount{}
	transaction.Amount.Total = fmt.Sprintf("%.2f", float64(order.Amount)/100)
	transaction.Amount.Currency = "USD"

	var item = &paypal.Item{}
	item.Name = order.Subject
	item.Description = ""
	item.Quantity = "1"
	item.Price = transaction.Amount.Total
	item.Tax = "0"
	item.SKU = "0"
	item.Currency = transaction.Amount.Currency
	transaction.ItemList = &paypal.ItemList{}
	transaction.ItemList.Items = []*paypal.Item{item}

	result, err = pay.client.CreatePayment(p)
	return result, err
}

func ExecuteApprovedPayment(paymentId string, payerId string) (result *paypal.Payment, err error) {
	return pay.client.ExecuteApprovedPayment(paymentId, payerId)
}
