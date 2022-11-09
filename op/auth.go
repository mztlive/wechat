package op

import (
	"context"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/openplatform"
	"github.com/skip2/go-qrcode"
)

const (
	// WechatAuthURL 微信授权链接
	WechatAuthURL = "https://open.weixin.qq.com/wxaopen/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d"
)

// GetPreQRCode 获得预授权二维码
func GetPreQRCode(ctx context.Context, serverHost string, OPAppID string) ([]byte, error) {
	var (
		opClient *openplatform.OpenPlatform
		err      error

		preCode string
		qrCode  []byte
	)

	if opClient, err = GetOpenPlatform(); err != nil {
		return nil, err
	}

	if preCode, err = opClient.GetPreCode(); err != nil {
		return nil, err
	}

	authURL := fmt.Sprintf(
		WechatAuthURL,
		OPAppID, preCode, url.QueryEscape(serverHost+"/wxop/authcallback"), 3,
	)

	if qrCode, err = qrcode.Encode(authURL, qrcode.Medium, 256); err != nil {
		return nil, err
	}

	return qrCode, nil
}
