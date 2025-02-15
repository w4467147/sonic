package vo

import "github.com/go-sonic/sonic/consts"

type MFAFactorAuth struct {
	QRImage    string         `json:"qrImage"`
	OptAuthUrl string         `json:"optAuthUrl"`
	MFAKey     string         `json:"mfaKey"`
	MFAType    consts.MFAType `json:"mfaType"`
}
