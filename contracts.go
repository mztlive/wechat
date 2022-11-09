package wechat

import (
	"context"

	"github.com/mztlive/wechat/oa"
)

// OfficialAccountQR 公众号的二维码相关功能抽象
type OfficialAccountQR interface {
	// GenTemporary 生成多个公众号临时二维码, 有效期600s
	// 如果其中一个公众号生成失败，那么就不返回这个公众号的二维码
	//	- officialAccounts: 公众号列表
	// 	- scene: 场景值
	GenTemporary(ctx context.Context, officialAccounts []oa.OfficialAccountKey, sceneStr string) (oa.QrCodes, error)
}
