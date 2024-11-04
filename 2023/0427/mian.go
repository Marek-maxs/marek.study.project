package main

import (
	"context"
	"log"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

/**
*
* Author: Marek
* Date: 2023-04-27 16:03
* Email: 364021318@qq.com
*
 */

func main() {
	var (
		mchID                      string = "1610335870"                               // 商户号
		mchCertificateSerialNumber string = "7D7B619EE080796B8E1CDEB7DD7DBE601845BDCB" // 商户证书序列号
		mchAPIv3Key                string = "0ac87e957ac71305ad4b51f2cba0ddd3"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:         core.String("wx00d9ed378c568c28"),
			Mchid:         core.String("1610335870"),
			Description:   core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:    core.String("1217752501201407033233368012"),
			TimeExpire:    core.Time(time.Now()),
			NotifyUrl:     core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			Amount: &jsapi.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(100),
			},
			Payer: &jsapi.Payer{
				Openid: core.String("o3PmI5fVUg2-JxlY58uFshZVJbrU"),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, *resp.PrepayId)
	}
}