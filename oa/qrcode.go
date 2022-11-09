package oa

import (
	"context"
	"errors"
	"fmt"

	"github.com/mztlive/logger"
	"github.com/mztlive/wechat/op"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	"github.com/silenceper/wechat/v2/openplatform"
	"go.uber.org/zap"
)

// QrCode 公众号二维码相关功能
type QrCode struct{}

// GenTemporary 获得多个公众号的订阅二维码（临时）
// 如果部分公众号获取不到二维码则对应值是空字符串
func (qr *QrCode) GenTemporary(ctx context.Context, officialAccounts []OfficialAccountKey, sceneStr string) (QrCodes, error) {
	var (
		err      error
		opClient *openplatform.OpenPlatform
	)

	if len(officialAccounts) == 0 {
		return nil, errors.New("appids is empty")
	}

	opClient, err = op.GetOpenPlatform()
	if err != nil {
		return nil, fmt.Errorf("get openplatform client failed. %w", err)
	}

	// 1. 获得每个公众号的订阅二维码
	qrCodes := make(map[string]string, len(officialAccounts))
	for _, item := range officialAccounts {
		var (
			qrCode string
			ticket *basic.Ticket
		)
		_, err := RefreshAuthorToken(ctx, item.AppID, item.RefreshToken)
		if err != nil {
			logger.Logger().Warn("refresh author token failed", zap.Error(err), zap.String("appid", item.AppID))
			continue
		}

		officialAccount := opClient.GetOfficialAccount(item.AppID)
		ticket, err = officialAccount.GetBasic().GetQRTicket(basic.NewTmpQrRequest(600, sceneStr))
		if err != nil {
			logger.Logger().Error("get qr code ticket failed.", zap.Error(err), zap.String("appid", item.AppID), zap.String("key", sceneStr))
			continue
		}

		qrCode = basic.ShowQRCode(ticket)
		qrCodes[item.AppID] = qrCode
	}

	return qrCodes, nil
}
