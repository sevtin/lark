package xalipay

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"lark/domain/po"
	"lark/pkg/conf"
	"log/slog"
	"net/url"
)

type Alipay struct {
	client *alipay.Client
	cfg    *conf.Alipay
}

func NewAlipay(cfg *conf.Alipay) (pay *Alipay) {
	var (
		client *alipay.Client
		err    error
	)
	if client, err = alipay.New(cfg.Appid, cfg.AppPrivateKey, false); err != nil {
		slog.Warn(err.Error())
		return
	}
	if err = client.LoadAppCertPublicKeyFromFile(cfg.AppCertPublicKey); err != nil {
		slog.Warn(err.Error())
		return
	}
	if err = client.LoadAliPayRootCertFromFile(cfg.AlipayRootCert); err != nil {
		slog.Warn(err.Error())
		return
	}
	if err = client.LoadAlipayCertPublicKeyFromFile(cfg.AlipayCertPublicKey); err != nil {
		slog.Warn(err.Error())
		return
	}
	if err = client.SetEncryptKey(cfg.EncryptKey); err != nil {
		slog.Warn(err.Error())
		return
	}
	pay = &Alipay{
		client: client,
		cfg:    cfg,
	}
	return
}

func (p *Alipay) CreateOrder(order *po.Order) (url *url.URL, err error) {
	var pay = alipay.TradePagePay{}
	pay.NotifyURL = p.cfg.Server + p.cfg.NotifyURL
	pay.ReturnURL = p.cfg.Server + p.cfg.ReturnURL
	pay.Subject = order.Subject
	pay.OutTradeNo = order.OrderSn
	pay.TotalAmount = fmt.Sprintf("%.2f", float64(order.Amount)/100)
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err = p.client.TradePagePay(pay)
	return
}
