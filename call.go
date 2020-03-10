package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"sync"
)

type OnMedia struct {
	Call *Call
}

type OnCallState struct{}

type Call struct {
	call     pjsua2.Call
	incoming bool

	closed bool

	account *Account

	mu sync.RWMutex

	hasMedia bool
	media    pjsua2.AudioMedia
	audio    *Audio
}

func NewCall(account *Account) *Call {
	call := &Call{
		incoming: false,
		closed:   false,
		account:  account,
		mu:       sync.RWMutex{},
		media:    nil,
	}

	call.call = pjsua2.NewDirectorCall(call, account.account)

	return call
}

func (p *Call) IsActive() bool {
	p.mu.RLock()
	if p.closed {
		p.mu.RUnlock()
		return false
	}
	active := p.call.IsActive()
	p.mu.RUnlock()
	return active
}

func (p *Call) close() {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return
	}
	p.closed = true
	id := p.call.GetInfo().GetCallIdString()
	p.mu.Unlock()

	p.account.removeCall(id, p)

	// Delete the call
	pjsua2.DeleteDirectorCall(p.call)
}

func (p *Call) SafeRun(fn func(call pjsua2.Call) interface{}) interface{} {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil
	}
	if fn != nil {
		return fn(p.call)
	}
	return nil
}

func (p *Call) OnCallState(prm pjsua2.OnCallStateParam) {
	ci := p.call.GetInfo()

	Debugf("[GetCall] onCallState %v, aor = %v", ci.GetStateText(), ci.GetRemoteUri())

	switch ci.GetState() {
	case pjsua2.PJSIP_INV_STATE_CALLING:
	case pjsua2.PJSIP_INV_STATE_CONFIRMED:
	case pjsua2.PJSIP_INV_STATE_DISCONNECTED:
	case pjsua2.PJSIP_INV_STATE_CONNECTING:
	case pjsua2.PJSIP_INV_STATE_EARLY:
	case pjsua2.PJSIP_INV_STATE_INCOMING:
	case pjsua2.PJSIP_INV_STATE_NULL:
	default:
	}

	if ci.GetState() == pjsua2.PJSIP_INV_STATE_DISCONNECTED {
		Debugf("[GetCall] GetCall Closed, CallId=%v, AOR=%v, reason=%v, lastStatusCode=%v",
			ci.GetCallIdString(), ci.GetRemoteUri(),
			ci.GetLastReason(), ci.GetLastStatusCode())

		p.close()
	}
}

func (p *Call) OnCallTsxState(prm pjsua2.OnCallTsxStateParam) {

}

func (p *Call) OnCallSdpCreated(arg2 pjsua2.OnCallSdpCreatedParam) {

}

func (p *Call) OnStreamCreated(arg2 pjsua2.OnStreamCreatedParam) {

}

func (p *Call) OnStreamDestroyed(arg2 pjsua2.OnStreamDestroyedParam) {

}

func (p *Call) OnDtmfDigit(arg2 pjsua2.OnDtmfDigitParam) {

}

func (p *Call) OnCallTransferRequest(arg2 pjsua2.OnCallTransferRequestParam) {

}

func (p *Call) OnCallTransferStatus(arg2 pjsua2.OnCallTransferStatusParam) {

}

func (p *Call) OnCallReplaceRequest(arg2 pjsua2.OnCallReplaceRequestParam) {

}

func (p *Call) OnCallReplaced(arg2 pjsua2.OnCallReplacedParam) {

}

func (p *Call) OnCallRxOffer(arg2 pjsua2.OnCallRxOfferParam) {

}

func (p *Call) OnCallRxReinvite(arg2 pjsua2.OnCallRxReinviteParam) {

}

func (p *Call) OnCallTxOffer(arg2 pjsua2.OnCallTxOfferParam) {}

func (p *Call) OnInstantMessage(arg2 pjsua2.OnInstantMessageParam) {}

func (p *Call) OnInstantMessageStatus(arg2 pjsua2.OnInstantMessageStatusParam) {}

func (p *Call) OnTypingIndication(arg2 pjsua2.OnTypingIndicationParam) {}

func (p *Call) OnCallRedirected(arg2 pjsua2.OnCallRedirectedParam) (_swig_ret pjsua2.Pjsip_redirect_op) {
	return _swig_ret
}
func (p *Call) OnCallMediaTransportState(arg2 pjsua2.OnCallMediaTransportStateParam) {
	switch arg2.GetState() {
	case pjsua2.PJSUA_MED_TP_CREATING:

	case pjsua2.PJSUA_MED_TP_DISABLED:
	case pjsua2.PJSUA_MED_TP_IDLE:
	case pjsua2.PJSUA_MED_TP_INIT:
	case pjsua2.PJSUA_MED_TP_RUNNING:
	case pjsua2.PJSUA_MED_TP_NULL:
	default:
	}
}
func (p *Call) OnCallMediaEvent(arg2 pjsua2.OnCallMediaEventParam) {

}

func (p *Call) OnCreateMediaTransport(arg2 pjsua2.OnCreateMediaTransportParam)         {}

func (p *Call) OnCreateMediaTransportSrtp(arg2 pjsua2.OnCreateMediaTransportSrtpParam) {}

func (p *Call) OnCallMediaState(arg2 pjsua2.OnCallMediaStateParam) {
	p.mu.Lock()
	defer p.mu.Unlock()

	Debugf("[GetCall] onCallMediaState")

	ci := p.call.GetInfo()

	medias := ci.GetMedia()
	for i := int64(0); i < medias.Size(); i++ {
		media := p.call.GetMedia(uint(i))
		if media.GetType() == pjsua2.PJMEDIA_TYPE_AUDIO {
			if p.hasMedia {
				break
			}

			//sc.hasMedia = true
			p.media = p.call.GetAudioMedia(int(i))
			p.onMedia()
		}
	}
}

func (p *Call) onMedia() {
	println("**** GOT MEDIA ****")
	if !p.incoming {
		//p.aud_med.StartTransmit(p.sipService.Endpoint.AudDevManager().GetPlaybackDevMedia())
		return
	}

	//p.player.SetPos(0)
	//p.player.StartTransmit(p.aud_med)
	//p.sipService.Endpoint.AudDevManager().GetPlaybackDevMedia().StartTransmit(p.aud_med)

	//p.aud_med.StartTransmit(p.player)

	//p.player.StartTransmit(p.sipService.Endpoint.AudDevManager().GetPlaybackDevMedia())
}
