package oa

// QrCodes 公众号二维码集合
// key 是公众号的 appID
// value 是公众号的二维码url
type QrCodes map[string]string

// OfficialAccountKey 公众号的关键Key
type OfficialAccountKey struct {
	AppID        string
	RefreshToken string
}
