package oa

import (
	"context"

	"github.com/mztlive/logger"
	"github.com/mztlive/wechat/op"
	"go.uber.org/zap"
)

// RefreshAuthorToken 刷新公众号的授权token
// 并不是每一次都刷新，如果已经存在则不刷新了。
// 返回最新可用的token， 如果刷新失败则返回空字符串，并包含具体的错误内容
func RefreshAuthorToken(ctx context.Context, appid string, refreshToken string) (string, error) {
	opClient, err := op.GetOpenPlatform()
	if err != nil {
		return "", err
	}

	authorToken, err := opClient.GetAuthrAccessToken(appid)
	if err != nil {
		logger.Logger().Error("get author token failed.", zap.Error(err), zap.String("appid", appid))
	}

	if authorToken == "" {
		token, err := opClient.RefreshAuthrToken(appid, refreshToken)
		if err != nil {
			return "", err
		}
		authorToken = token.AccessToken
	}

	return authorToken, nil
}
