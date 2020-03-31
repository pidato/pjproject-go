package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"sync"
)

type PeerAccountConfig struct {

}

type UserAccountConfig struct {

}

type account struct {
	account pjsua2.Account

	fnIncoming func(id int, data interface{})
	calls      map[int]*call
	mu         sync.RWMutex
}

func newPeerAccount(config *PeerAccountConfig) *account {
	a := &account{
		account: nil,
		calls:   make(map[int]*call, 128),
		mu:      sync.RWMutex{},
	}
	a.account = pjsua2.NewDirectorAccount(a)
	pjsua2.NewAccountConfig()
	pjsua2.DeleteAccountConfig()
	return a
}

func (a *account) removeCall(id int, call *call) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.calls, id)
}

func (a *account) OnIncomingCall(arg2 pjsua2.OnIncomingCallParam) {
	err := exec(func() {
		call := newIncomingCall(a, arg2.GetCallId(), newSipRxData(arg2.GetRdata()))

		param := pjsua2.NewCallOpParam()
		param.GetOpt().SetAudioCount(1)
		param.GetOpt().SetVideoCount(0)
		param.SetStatusCode(pjsua2.PJSIP_SC_OK)

		a.mu.Lock()
		a.calls[call.call.GetId()] = call
		a.mu.Unlock()

		call.call.Answer(param)
	})
	if err != nil {
		Errorf("Account::OnIncomingCall %s", err.Error())
	}
}

func (a *account) OnRegStarted(arg2 pjsua2.OnRegStartedParam) {

}

func (a *account) OnRegState(arg2 pjsua2.OnRegStateParam) {

}

func (a *account) OnIncomingSubscribe(arg2 pjsua2.OnIncomingSubscribeParam) {

}

func (a *account) OnInstantMessage(arg2 pjsua2.OnInstantMessageParam) {

}

func (a *account) OnInstantMessageStatus(arg2 pjsua2.OnInstantMessageStatusParam) {

}

func (a *account) OnTypingIndication(arg2 pjsua2.OnTypingIndicationParam) {

}

func (a *account) OnMwiInfo(arg2 pjsua2.OnMwiInfoParam) {

}
