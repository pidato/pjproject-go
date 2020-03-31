package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"sync"
)

type OnMedia struct {
	Call *call
}

type OnCallState struct{}

type call struct {
	call     pjsua2.Call
	id       int
	idStr    string
	incoming bool
	rxData   SipRxData

	closed bool

	account *account

	mu sync.RWMutex

	hasMedia bool
	media    pjsua2.AudioMedia
	audio    *Audio
}

func newIncomingCall(account *account, id int, rxData *SipRxData) *call {
	call := &call{
		id:       id,
		rxData:   *rxData,
		incoming: false,
		closed:   false,
		account:  account,
		mu:       sync.RWMutex{},
		media:    nil,
	}

	call.call = pjsua2.NewDirectorCall(call, account.account)

	return call
}

func (c *call) IsActive() bool {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return false
	}
	active := c.call.IsActive()
	c.mu.RUnlock()
	return active
}

func (c *call) close() {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return
	}
	c.closed = true
	id := c.id
	c.mu.Unlock()

	c.account.removeCall(id, c)

	// Delete the call
	pjsua2.DeleteDirectorCall(c.call)
}

func (c *call) Hangup() {
	c.exec(func() {

	})
}

func (c *call) SafeRun(fn func(call pjsua2.Call) interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil
	}
	if fn != nil {
		return fn(c.call)
	}
	return nil
}

func (c *call) exec(fn func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return
	}
	if fn == nil {
		return
	}
	fn()
}

func (c *call) OnCallState(prm pjsua2.OnCallStateParam) {
	c.exec(func() {
		ci := c.call.GetInfo()

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

			c.close()
		}
	})
}

func (c *call) OnCallTsxState(prm pjsua2.OnCallTsxStateParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallSdpCreated(arg2 pjsua2.OnCallSdpCreatedParam) {
	c.exec(func() {

	})
}

func (c *call) OnStreamCreated(arg2 pjsua2.OnStreamCreatedParam) {
	c.exec(func() {

	})
}

func (c *call) OnStreamDestroyed(arg2 pjsua2.OnStreamDestroyedParam) {
	c.exec(func() {

	})
}

func (c *call) OnDtmfDigit(arg2 pjsua2.OnDtmfDigitParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallTransferRequest(arg2 pjsua2.OnCallTransferRequestParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallTransferStatus(arg2 pjsua2.OnCallTransferStatusParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallReplaceRequest(arg2 pjsua2.OnCallReplaceRequestParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallReplaced(arg2 pjsua2.OnCallReplacedParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallRxOffer(arg2 pjsua2.OnCallRxOfferParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallRxReinvite(arg2 pjsua2.OnCallRxReinviteParam) {
	c.exec(func() {

	})
}

func (c *call) OnCallTxOffer(arg2 pjsua2.OnCallTxOfferParam) {
	c.exec(func() {

	})
}

func (c *call) OnInstantMessage(arg2 pjsua2.OnInstantMessageParam) {
	c.exec(func() {

	})
}

func (c *call) OnInstantMessageStatus(arg2 pjsua2.OnInstantMessageStatusParam) {
	c.exec(func() {

	})
}

func (c *call) OnTypingIndication(arg2 pjsua2.OnTypingIndicationParam) {}

func (c *call) OnCallRedirected(arg2 pjsua2.OnCallRedirectedParam) (_swig_ret pjsua2.Pjsip_redirect_op) {
	return _swig_ret
}

func (c *call) OnCallMediaTransportState(arg2 pjsua2.OnCallMediaTransportStateParam) {
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

func (c *call) OnCallMediaEvent(arg2 pjsua2.OnCallMediaEventParam) {
}

func (c *call) OnCreateMediaTransport(arg2 pjsua2.OnCreateMediaTransportParam) {

}

func (c *call) OnCreateMediaTransportSrtp(arg2 pjsua2.OnCreateMediaTransportSrtpParam) {}

func (c *call) getAudioMedia() (media pjsua2.AudioMedia, err error) {
	err = exec(func() {
		media = c.call.GetAudioMedia(-1)
	})
	return
}

func (c *call) OnCallMediaState(arg2 pjsua2.OnCallMediaStateParam) {
	c.exec(func() {
		Debugf("[Call::OnCallMediaState]")

		var err error
		var media pjsua2.AudioMedia
		media, err = c.getAudioMedia()
		if err != nil {
			return
		}

		c.media = media
		if c.media != media {
			if c.media == nil {
				c.media = media
				c.onMedia()
			} else {
				if media == nil {
					c.media = nil
					c.onLostMedia()
				} else {
					c.media = media
					c.onMediaChanged()
				}
			}
		}
	})
}

func (c *call) onLostMedia() {

}

func (c *call) onMediaChanged() {

}

func (c *call) onMedia() {
	//Debugf("**** GOT MEDIA ****")
	if !c.incoming {
		//p.aud_med.StartTransmit(p.sipService.Endpoint.AudDevManager().GetPlaybackDevMedia())
		return
	}

	//p.player.SetPos(0)
	//p.player.StartTransmit(p.aud_med)
	//p.sipService.Endpoint.AudDevManager().GetPlaybackDevMedia().StartTransmit(p.aud_med)

	//p.aud_med.StartTransmit(p.player)

	//p.player.StartTransmit(p.sipService.Endpoint.AudDevManager().GetPlaybackDevMedia())
}
