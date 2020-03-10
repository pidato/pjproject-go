package pj

import (
	"github.com/pidato/chanx"
	"github.com/pidato/pjproject-go/pjsua2"
	"os"
	"sync"

	_ "github.com/pidato/pjproject-go/vad"
)

type Message interface {
}

//
type Service struct {
	pjsua2.Endpoint

	epConfig pjsua2.EpConfig
	calls    map[string]*Call
	counter  uint32

	closed bool
	ch     chanx.C

	mu sync.RWMutex
}

var (
	logWriter = pjsua2.NewDirectorLogWriter(new(LogWriter))
)

func Start(config Config) (*Service, error) {
	s := &Service{
		calls: make(map[string]*Call),
		ch:    chanx.Make(5),
	}

	s.Endpoint = pjsua2.NewDirectorEndpoint(s)

	s.Endpoint.LibCreate()

	// SetLogLevel library
	epConfig := pjsua2.NewEpConfig()
	epConfig.GetLogConfig().SetLevel(5)

	//switch config.LogLevel {
	//case zerolog.TraceLevel:
	//	epConfig.GetLogConfig().SetLevel(7)
	//case zerolog.DebugLevel:
	//	epConfig.GetLogConfig().SetLevel(6)
	//case zerolog.InfoLevel:
	//	epConfig.GetLogConfig().SetLevel(4)
	//case zerolog.WarnLevel:
	//	epConfig.GetLogConfig().SetLevel(3)
	//case zerolog.ErrorLevel:
	//	epConfig.GetLogConfig().SetLevel(2)
	//case zerolog.FatalLevel:
	//	epConfig.GetLogConfig().SetLevel(1)
	//case zerolog.PanicLevel:
	//	epConfig.GetLogConfig().SetLevel(0)
	//}
	SetLogLevel(config.LogLevel)

	epConfig.GetLogConfig().SetWriter(logWriter)

	if config.MaxCalls == 0 {
		config.MaxCalls = 1024
	}
	epConfig.GetUaConfig().SetMaxCalls(config.MaxCalls)
	epConfig.GetUaConfig().SetThreadCnt(config.Threads)
	if len(config.UserAgent) > 0 {
		epConfig.GetUaConfig().SetUserAgent(config.UserAgent)
	}

	// Hardcode to 16KHz
	epConfig.GetMedConfig().SetSndClockRate(config.ClockRate)
	epConfig.GetMedConfig().SetClockRate(config.ClockRate)
	// Hardcode to 10ms ptime.
	epConfig.GetMedConfig().SetPtime(config.Ptime)
	epConfig.GetMedConfig().SetSndRecLatency(config.Ptime)
	epConfig.GetMedConfig().SetSndPlayLatency(config.Ptime)
	epConfig.GetMedConfig().SetAudioFramePtime(config.AudioFramePtime)
	epConfig.GetMedConfig().SetChannelCount(1)
	epConfig.GetMedConfig().SetNoVad(config.NoVad)
	epConfig.GetMedConfig().SetEcTailLen(config.EcTailLen)
	epConfig.GetMedConfig().SetEcOptions(config.EcOptions)
	epConfig.GetMedConfig().SetHasIoqueue(true)
	if config.MediaPorts > 0 {
		epConfig.GetMedConfig().SetMaxMediaPorts(config.MediaPorts)
	}

	s.epConfig = epConfig

	s.Endpoint.LibInit(epConfig)

	//s.Endpoint.CodecSetPriority("opus/48000/2", 254)

	// Transport
	transportConfig := pjsua2.NewTransportConfig()
	if len(config.BindAddress) > 0 {
		transportConfig.SetBoundAddress(config.BindAddress)
	}
	if len(config.PublicAddress) > 0 {
		transportConfig.SetPublicAddress(config.PublicAddress)
	}
	if config.Port > 0 {
		transportConfig.SetPort(config.Port)
	}
	s.Endpoint.TransportCreate(pjsua2.PJSIP_TRANSPORT_UDP, transportConfig)

	Infof("[PJSIP] Available codecs:")
	for i := 0; i < int(s.Endpoint.CodecEnum2().Size()); i++ {
		c := s.Endpoint.CodecEnum2().Get(i)
		Infof("\t - %s (priority: %d)", c.GetCodecId(), c.GetPriority())
	}

	s.Endpoint.LibStart()

	s.AudDevManager().SetNullDev()

	var wg sync.WaitGroup
	wg.Add(1)
	go s.run(&wg)
	wg.Wait()

	Infof("[PJSIP] pjsua2 started!")

	return s, nil
}

func (ps *Service) Close() error {
	ps.CheckThread()
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.closed {
		return os.ErrClosed
	}
	ps.Endpoint.LibDestroy()
	return nil
}

func (ps *Service) run(wg *sync.WaitGroup) {
	wg.Done()
	ps.CheckThread()

	for {
		msg, ok := ps.ch.Recv()
		if !ok {
			break
		}

		switch m := msg.(type) {
		case *MsgStopAudio:
			m.Audio.Stop()
		default:
		}
	}
}

func (ps *Service) CheckThread() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.Endpoint.LibIsThreadRegistered() {
		ps.Endpoint.LibRegisterThread("")
	}
}

func (ps *Service) OnTimer(arg2 pjsua2.OnTimerParam) {

}

func (ps *Service) OnSelectAccount(arg2 pjsua2.OnSelectAccountParam) {

}
