package oa

import (
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/mztlive/wechat/cache"
	wechatPkg "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

// NewClient 获得一个微信公众号客户端
func NewClient() (*officialaccount.OfficialAccount, error) {
	var (
		appID          string
		appSecret      string
		token          string
		encodingAESKey string
	)

	appID = config.String("Wechat.OfficialAccount.AppID")
	appSecret = config.String("Wechat.OfficialAccount.AppSecret")
	token = config.String("Wechat.OfficialAccount.Token")
	encodingAESKey = config.String("Wechat.OfficialAccount.EncodingAESKey")

	if appID == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform AppID is empty")
	}

	if appSecret == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform AppSecret is empty")
	}

	if token == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform Token is empty")
	}

	if encodingAESKey == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform EncodingAESKey is empty")
	}

	cfg := &offConfig.Config{
		AppID:          appID,
		AppSecret:      appSecret,
		Token:          token,
		EncodingAESKey: encodingAESKey,
		Cache:          cache.RedisCache(),
	}

	wc := wechatPkg.NewWechat()
	return wc.GetOfficialAccount(cfg), nil
}
