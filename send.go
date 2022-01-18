package webhookForward

import (
	"context"
	"net/http"
	"time"

	"github.com/go-sdk/lib/bot/dingtalk"
	"github.com/go-sdk/lib/bot/telegram"
	"github.com/go-sdk/lib/bot/weixin"
	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/errx"
	"github.com/go-sdk/lib/httpx"
	"github.com/go-sdk/lib/log"
	"github.com/go-sdk/lib/pool"
	"github.com/go-sdk/lib/srv"
)

var (
	debug = conf.Get("debug").Bool()
	proxy = conf.Get("proxy").String()

	hc = httpx.New(httpx.WithDebug(debug)).SetProxy(proxy)

	dingtalkEnable = conf.Get("dingtalk.enable").Bool()
	dingtalkToken  = conf.Get("dingtalk.token").String()
	dingtalkSecret = conf.Get("dingtalk.secret").String()

	telegramEnable = conf.Get("telegram.enable").Bool()
	telegramToken  = conf.Get("telegram.token").String()
	telegramTo     = conf.Get("telegram.to").String()

	weixinEnable = conf.Get("weixin.enable").Bool()
	weixinId     = conf.Get("weixin.id").String()
	weixinSecret = conf.Get("weixin.secret").String()
	weixinAid    = conf.Get("weixin.aid").Int64()
	weixinTo     = conf.Get("weixin.to").String()
)

func Send(c *srv.Context) {
	message := c.Query("message")
	if message == "" {
		c.JSON(http.StatusBadRequest, errx.BadRequest("message is empty"))
		return
	}
	log.WithContext(c).Debug(message)
	p := pool.New(c)
	p.Add(func(ctx context.Context) error {
		if dingtalkEnable {
			return dingtalk.SendMessage(hc, dingtalkToken, dingtalkSecret, message)
		}
		return nil
	})
	p.Add(func(ctx context.Context) error {
		if telegramEnable {
			return telegram.SendMessage(hc, telegramToken, telegramTo, message)
		}
		return nil
	})
	p.Add(func(ctx context.Context) error {
		if weixinEnable {
			at, err := mc.GetOrFetch(
				"webhook:forward:weixin:"+weixinId,
				func() (interface{}, time.Duration, error) {
					resp, err := weixin.GetToken(hc, weixinId, weixinSecret)
					if err != nil {
						return nil, 0, err
					}
					return resp.AccessToken, time.Duration(resp.ExpiresIn)*time.Second - time.Minute, nil
				},
				func(v interface{}) (interface{}, error) {
					return v, nil
				},
			)
			if err != nil {
				return err
			}
			return weixin.SendMessage(hc, weixinAid, at.(string), weixinTo, message)
		}
		return nil
	})
	err := p.Run(3)
	if err != nil {
		log.WithContext(c).Error(err)
	}
	c.JSON(http.StatusOK, errx.OK("ok"))
}
