/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./pjsua2.i

#ifndef SWIG_pjsua2_WRAP_H_
#define SWIG_pjsua2_WRAP_H_

class Swig_memory;

class SwigDirector_AudioMediaPlayer : public pj::AudioMediaPlayer
{
 public:
  SwigDirector_AudioMediaPlayer(int swig_p);
  virtual ~SwigDirector_AudioMediaPlayer();
  void _swig_upcall_onEof2() {
    pj::AudioMediaPlayer::onEof2();
  }
  virtual void onEof2();
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_Buddy : public pj::Buddy
{
 public:
  SwigDirector_Buddy(int swig_p);
  virtual ~SwigDirector_Buddy();
  void _swig_upcall_onBuddyState() {
    pj::Buddy::onBuddyState();
  }
  virtual void onBuddyState();
  void _swig_upcall_onBuddyEvSubState(pj::OnBuddyEvSubStateParam &prm) {
    pj::Buddy::onBuddyEvSubState(prm);
  }
  virtual void onBuddyEvSubState(pj::OnBuddyEvSubStateParam &prm);
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_FindBuddyMatch : public pj::FindBuddyMatch
{
 public:
  SwigDirector_FindBuddyMatch(int swig_p);
  bool _swig_upcall_match(pj::string const &token, pj::Buddy const &buddy) {
    return pj::FindBuddyMatch::match(token,buddy);
  }
  virtual bool match(pj::string const &token, pj::Buddy const &buddy);
  virtual ~SwigDirector_FindBuddyMatch();
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_Account : public pj::Account
{
 public:
  SwigDirector_Account(int swig_p);
  virtual ~SwigDirector_Account();
  void _swig_upcall_onIncomingCall(pj::OnIncomingCallParam &prm) {
    pj::Account::onIncomingCall(prm);
  }
  virtual void onIncomingCall(pj::OnIncomingCallParam &prm);
  void _swig_upcall_onRegStarted(pj::OnRegStartedParam &prm) {
    pj::Account::onRegStarted(prm);
  }
  virtual void onRegStarted(pj::OnRegStartedParam &prm);
  void _swig_upcall_onRegState(pj::OnRegStateParam &prm) {
    pj::Account::onRegState(prm);
  }
  virtual void onRegState(pj::OnRegStateParam &prm);
  void _swig_upcall_onIncomingSubscribe(pj::OnIncomingSubscribeParam &prm) {
    pj::Account::onIncomingSubscribe(prm);
  }
  virtual void onIncomingSubscribe(pj::OnIncomingSubscribeParam &prm);
  void _swig_upcall_onInstantMessage(pj::OnInstantMessageParam &prm) {
    pj::Account::onInstantMessage(prm);
  }
  virtual void onInstantMessage(pj::OnInstantMessageParam &prm);
  void _swig_upcall_onInstantMessageStatus(pj::OnInstantMessageStatusParam &prm) {
    pj::Account::onInstantMessageStatus(prm);
  }
  virtual void onInstantMessageStatus(pj::OnInstantMessageStatusParam &prm);
  void _swig_upcall_onTypingIndication(pj::OnTypingIndicationParam &prm) {
    pj::Account::onTypingIndication(prm);
  }
  virtual void onTypingIndication(pj::OnTypingIndicationParam &prm);
  void _swig_upcall_onMwiInfo(pj::OnMwiInfoParam &prm) {
    pj::Account::onMwiInfo(prm);
  }
  virtual void onMwiInfo(pj::OnMwiInfoParam &prm);
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_Call : public pj::Call
{
 public:
  SwigDirector_Call(int swig_p, pj::Account &acc, int call_id);
  SwigDirector_Call(int swig_p, pj::Account &acc);
  virtual ~SwigDirector_Call();
  void _swig_upcall_onCallState(pj::OnCallStateParam &prm) {
    pj::Call::onCallState(prm);
  }
  virtual void onCallState(pj::OnCallStateParam &prm);
  void _swig_upcall_onCallTsxState(pj::OnCallTsxStateParam &prm) {
    pj::Call::onCallTsxState(prm);
  }
  virtual void onCallTsxState(pj::OnCallTsxStateParam &prm);
  void _swig_upcall_onCallMediaState(pj::OnCallMediaStateParam &prm) {
    pj::Call::onCallMediaState(prm);
  }
  virtual void onCallMediaState(pj::OnCallMediaStateParam &prm);
  void _swig_upcall_onCallSdpCreated(pj::OnCallSdpCreatedParam &prm) {
    pj::Call::onCallSdpCreated(prm);
  }
  virtual void onCallSdpCreated(pj::OnCallSdpCreatedParam &prm);
  void _swig_upcall_onStreamCreated(pj::OnStreamCreatedParam &prm) {
    pj::Call::onStreamCreated(prm);
  }
  virtual void onStreamCreated(pj::OnStreamCreatedParam &prm);
  void _swig_upcall_onStreamDestroyed(pj::OnStreamDestroyedParam &prm) {
    pj::Call::onStreamDestroyed(prm);
  }
  virtual void onStreamDestroyed(pj::OnStreamDestroyedParam &prm);
  void _swig_upcall_onDtmfDigit(pj::OnDtmfDigitParam &prm) {
    pj::Call::onDtmfDigit(prm);
  }
  virtual void onDtmfDigit(pj::OnDtmfDigitParam &prm);
  void _swig_upcall_onCallTransferRequest(pj::OnCallTransferRequestParam &prm) {
    pj::Call::onCallTransferRequest(prm);
  }
  virtual void onCallTransferRequest(pj::OnCallTransferRequestParam &prm);
  void _swig_upcall_onCallTransferStatus(pj::OnCallTransferStatusParam &prm) {
    pj::Call::onCallTransferStatus(prm);
  }
  virtual void onCallTransferStatus(pj::OnCallTransferStatusParam &prm);
  void _swig_upcall_onCallReplaceRequest(pj::OnCallReplaceRequestParam &prm) {
    pj::Call::onCallReplaceRequest(prm);
  }
  virtual void onCallReplaceRequest(pj::OnCallReplaceRequestParam &prm);
  void _swig_upcall_onCallReplaced(pj::OnCallReplacedParam &prm) {
    pj::Call::onCallReplaced(prm);
  }
  virtual void onCallReplaced(pj::OnCallReplacedParam &prm);
  void _swig_upcall_onCallRxOffer(pj::OnCallRxOfferParam &prm) {
    pj::Call::onCallRxOffer(prm);
  }
  virtual void onCallRxOffer(pj::OnCallRxOfferParam &prm);
  void _swig_upcall_onCallRxReinvite(pj::OnCallRxReinviteParam &prm) {
    pj::Call::onCallRxReinvite(prm);
  }
  virtual void onCallRxReinvite(pj::OnCallRxReinviteParam &prm);
  void _swig_upcall_onCallTxOffer(pj::OnCallTxOfferParam &prm) {
    pj::Call::onCallTxOffer(prm);
  }
  virtual void onCallTxOffer(pj::OnCallTxOfferParam &prm);
  void _swig_upcall_onInstantMessage(pj::OnInstantMessageParam &prm) {
    pj::Call::onInstantMessage(prm);
  }
  virtual void onInstantMessage(pj::OnInstantMessageParam &prm);
  void _swig_upcall_onInstantMessageStatus(pj::OnInstantMessageStatusParam &prm) {
    pj::Call::onInstantMessageStatus(prm);
  }
  virtual void onInstantMessageStatus(pj::OnInstantMessageStatusParam &prm);
  void _swig_upcall_onTypingIndication(pj::OnTypingIndicationParam &prm) {
    pj::Call::onTypingIndication(prm);
  }
  virtual void onTypingIndication(pj::OnTypingIndicationParam &prm);
  pjsip_redirect_op _swig_upcall_onCallRedirected(pj::OnCallRedirectedParam &prm) {
    return pj::Call::onCallRedirected(prm);
  }
  virtual pjsip_redirect_op onCallRedirected(pj::OnCallRedirectedParam &prm);
  void _swig_upcall_onCallMediaTransportState(pj::OnCallMediaTransportStateParam &prm) {
    pj::Call::onCallMediaTransportState(prm);
  }
  virtual void onCallMediaTransportState(pj::OnCallMediaTransportStateParam &prm);
  void _swig_upcall_onCallMediaEvent(pj::OnCallMediaEventParam &prm) {
    pj::Call::onCallMediaEvent(prm);
  }
  virtual void onCallMediaEvent(pj::OnCallMediaEventParam &prm);
  void _swig_upcall_onCreateMediaTransport(pj::OnCreateMediaTransportParam &prm) {
    pj::Call::onCreateMediaTransport(prm);
  }
  virtual void onCreateMediaTransport(pj::OnCreateMediaTransportParam &prm);
  void _swig_upcall_onCreateMediaTransportSrtp(pj::OnCreateMediaTransportSrtpParam &prm) {
    pj::Call::onCreateMediaTransportSrtp(prm);
  }
  virtual void onCreateMediaTransportSrtp(pj::OnCreateMediaTransportSrtpParam &prm);
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_PiRecorder : public PiRecorder
{
 public:
  SwigDirector_PiRecorder(int swig_p);
  virtual ~SwigDirector_PiRecorder();
  void _swig_upcall_onHeartbeat() {
    PiRecorder::onHeartbeat();
  }
  virtual void onHeartbeat();
  void _swig_upcall_onError(pj::Error &e) {
    PiRecorder::onError(e);
  }
  virtual void onError(pj::Error &e);
  void _swig_upcall_onFrameDTX(void *frame, pj_uint64_t prevExternCPU) {
    PiRecorder::onFrameDTX(frame,prevExternCPU);
  }
  virtual void onFrameDTX(void *frame, pj_uint64_t prevExternCPU);
  void _swig_upcall_onFrame(void *frame, pj_uint64_t prevExternCPU) {
    PiRecorder::onFrame(frame,prevExternCPU);
  }
  virtual void onFrame(void *frame, pj_uint64_t prevExternCPU);
  void _swig_upcall_onDestroy() {
    PiRecorder::onDestroy();
  }
  virtual void onDestroy();
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_PiPort : public PiPort
{
 public:
  SwigDirector_PiPort(int swig_p);
  virtual ~SwigDirector_PiPort();
  void _swig_upcall_onPutFrame(pjmedia_frame_type frameType, void *pcm, pj_size_t size, pj_uint64_t timestamp, pj_uint32_t bit_info) {
    PiPort::onPutFrame(frameType,pcm,size,timestamp,bit_info);
  }
  virtual void onPutFrame(pjmedia_frame_type frameType, void *pcm, pj_size_t size, pj_uint64_t timestamp, pj_uint32_t bit_info);
  void _swig_upcall_onGetFrame(pjmedia_frame_type frameType, void *pcm, pj_size_t size, pj_uint64_t timestamp, pj_uint32_t bit_info) {
    PiPort::onGetFrame(frameType,pcm,size,timestamp,bit_info);
  }
  virtual void onGetFrame(pjmedia_frame_type frameType, void *pcm, pj_size_t size, pj_uint64_t timestamp, pj_uint32_t bit_info);
  void _swig_upcall_onDestroy() {
    PiPort::onDestroy();
  }
  virtual void onDestroy();
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_LogWriter : public pj::LogWriter
{
 public:
  SwigDirector_LogWriter(int swig_p);
  virtual ~SwigDirector_LogWriter();
  virtual void write(pj::LogEntry const &entry);
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

class SwigDirector_Endpoint : public pj::Endpoint
{
 public:
  SwigDirector_Endpoint(int swig_p);
  virtual ~SwigDirector_Endpoint();
  void _swig_upcall_onNatDetectionComplete(pj::OnNatDetectionCompleteParam const &prm) {
    pj::Endpoint::onNatDetectionComplete(prm);
  }
  virtual void onNatDetectionComplete(pj::OnNatDetectionCompleteParam const &prm);
  void _swig_upcall_onNatCheckStunServersComplete(pj::OnNatCheckStunServersCompleteParam const &prm) {
    pj::Endpoint::onNatCheckStunServersComplete(prm);
  }
  virtual void onNatCheckStunServersComplete(pj::OnNatCheckStunServersCompleteParam const &prm);
  void _swig_upcall_onTransportState(pj::OnTransportStateParam const &prm) {
    pj::Endpoint::onTransportState(prm);
  }
  virtual void onTransportState(pj::OnTransportStateParam const &prm);
  void _swig_upcall_onTimer(pj::OnTimerParam const &prm) {
    pj::Endpoint::onTimer(prm);
  }
  virtual void onTimer(pj::OnTimerParam const &prm);
  void _swig_upcall_onSelectAccount(pj::OnSelectAccountParam &prm) {
    pj::Endpoint::onSelectAccount(prm);
  }
  virtual void onSelectAccount(pj::OnSelectAccountParam &prm);
  void _swig_upcall_onIpChangeProgress(pj::OnIpChangeProgressParam &prm) {
    pj::Endpoint::onIpChangeProgress(prm);
  }
  virtual void onIpChangeProgress(pj::OnIpChangeProgressParam &prm);
  void _swig_upcall_onMediaEvent(pj::OnMediaEventParam &prm) {
    pj::Endpoint::onMediaEvent(prm);
  }
  virtual void onMediaEvent(pj::OnMediaEventParam &prm);
 private:
  intgo go_val;
  Swig_memory *swig_mem;
};

#endif
