package pj

import "github.com/pidato/pjproject-go/pjsua2"

type SipRxData struct {
	Info       string
	WholeMsg   string
	SrcAddress string
}

func newSipRxData(from pjsua2.SipRxData) *SipRxData {
	return &SipRxData{
		Info:       from.GetInfo(),
		WholeMsg:   from.GetWholeMsg(),
		SrcAddress: from.GetSrcAddress(),
	}
}

type SipHeader struct {}
type SipHeaders []*SipHeader

type AccountConfig struct {
	Priority       int
	IdUri          string
	RegConfig      AccountRegConfig
	SipConfig      AccountSipConfig
	CallConfig     AccountCallConfig
	PresConfig     AccountPresConfig
	MwiConfig      AccountMwiConfig
	NatConfig      AccountNatConfig
	MediaConfig    AccountMediaConfig
	VideoConfig    AccountVideoConfig
	IpChangeConfig AccountIpChangeConfig
}

func (a *AccountConfig) ToPJ() pjsua2.AccountConfig {
	to := pjsua2.NewAccountConfig()
	a.Priority = to.GetPriority()
	a.IdUri = to.GetIdUri()

	return to
}

type AccountRegConfig struct {
	RegistrarUri string
	RegisterOnAdd bool

}

type AccountSipConfig struct {
}

type AccountCallConfig struct {
}

type AccountPresConfig struct {
}

type AccountMwiConfig struct {
}

type AccountNatConfig struct {
}

type AccountMediaConfig struct {
}

type AccountVideoConfig struct {
}

type AccountIpChangeConfig struct {
}
