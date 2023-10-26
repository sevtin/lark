package xalipay

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"lark/domain/po"
	"lark/pkg/conf"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

var (
	pay *Alipay
)

type Alipay struct {
	client *alipay.Client
	cfg    *conf.Alipay
}

func NewAlipay(cfg *conf.Alipay) *Alipay {
	var (
		client *alipay.Client
		buf    []byte
		err    error
	)
	if buf, err = os.ReadFile(cfg.AppPrivateKey); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	cfg.AppPrivateKey = string(buf)
	if client, err = alipay.New(cfg.Appid, cfg.AppPrivateKey, false); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	if err = client.LoadAppCertPublicKeyFromFile(cfg.AppCertPublicKey); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	if err = client.LoadAliPayRootCertFromFile(cfg.AlipayRootCert); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	if err = client.LoadAlipayCertPublicKeyFromFile(cfg.AlipayCertPublicKey); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	//if err = client.SetEncryptKey(cfg.EncryptKey); err != nil {
	//	slog.Warn(err.Error())
	//	return nil
	//}
	pay = &Alipay{
		client: client,
		cfg:    cfg,
	}
	return pay
}

func TradePagePay(order *po.Order) (url *url.URL, err error) {
	if pay == nil {
		return
	}
	var tpp = alipay.TradePagePay{}
	tpp.NotifyURL = pay.cfg.Server + pay.cfg.NotifyURL
	tpp.ReturnURL = pay.cfg.Server + pay.cfg.ReturnURL
	tpp.Subject = order.Subject
	tpp.OutTradeNo = order.OrderSn
	tpp.TotalAmount = fmt.Sprintf("%.2f", float64(order.Amount)/100)
	tpp.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err = pay.client.TradePagePay(tpp)
	return
}

func TradePreCreate(tradeNo, amount, subject string) (result *alipay.TradePreCreateRsp, err error) {
	if pay == nil {
		return
	}
	var tpc = alipay.TradePreCreate{
		Trade: alipay.Trade{
			Subject:     subject,
			OutTradeNo:  tradeNo,
			TotalAmount: amount,
			NotifyURL:   pay.cfg.Server + pay.cfg.NotifyURL,
			ReturnURL:   pay.cfg.Server + pay.cfg.ReturnURL,
		},
	}
	return pay.client.TradePreCreate(tpc)
}

func TradeQuery(tradeNo string) (result *alipay.TradeQueryRsp, err error) {
	tq := alipay.TradeQuery{
		OutTradeNo: tradeNo,
	}
	return pay.client.TradeQuery(tq)
}

func DecodeNotify(writer http.ResponseWriter, request *http.Request) (notify *alipay.Notification, err error) {
	notify, err = pay.client.DecodeNotification(request.Form)
	if err != nil {
		// 错误处理
		slog.Error(err.Error())
		return
	}
	alipay.ACKNotification(writer)
	return
}
