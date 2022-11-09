package op

import (
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/mztlive/wechat/cache"
	wechatPkg "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/openplatform"
	openPlatformCfg "github.com/silenceper/wechat/v2/openplatform/config"
)

// GetOpenPlatform 获得一个OpenPlatform对象
func GetOpenPlatform() (*openplatform.OpenPlatform, error) {

	var (
		OPAppID          string
		OPAppSecret      string
		OPToken          string
		OPEncodingAESKey string
	)

	OPAppID = config.String("Wechat.OpenPlatform.AppID")
	OPAppSecret = config.String("Wechat.OpenPlatform.AppSecret")
	OPToken = config.String("Wechat.OpenPlatform.Token")
	OPEncodingAESKey = config.String("Wechat.OpenPlatform.EncodingAESKey")

	if OPAppID == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform AppID is empty")
	}

	if OPAppSecret == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform AppSecret is empty")
	}

	if OPToken == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform Token is empty")
	}

	if OPEncodingAESKey == "" {
		return nil, fmt.Errorf("Wechat OpenPlatform EncodingAESKey is empty")
	}

	cfg := &openPlatformCfg.Config{
		AppID:          OPAppID,
		AppSecret:      OPAppSecret,
		Token:          OPToken,
		EncodingAESKey: OPEncodingAESKey,
		Cache:          cache.RedisCache(),
	}

	wc := wechatPkg.NewWechat()
	openPlatform := wc.GetOpenPlatform(cfg)

	return openPlatform, nil
}
