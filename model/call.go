package model

import "github.com/pidato/pjproject-go/pjsua2"

type OnCallTsxStateParam struct {

}

type SipEvent struct {
	Xtype pjsua2.Pjsip_event_id_e
	Body SipEventBody
}

type SipEventBody struct {

}

type TimerEvent struct {}

type TsxStateEvent struct {
}

type TxMsgEvent struct {}

type TxErrorEvent struct {}

type RxMsgEvent struct {}

type UserEvent struct {}