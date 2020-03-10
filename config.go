package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"github.com/rs/zerolog"
	"runtime"
)

type Config struct {
	MaxCalls      uint
	Threads       uint
	BindAddress   string
	PublicAddress string
	Port          uint
	MediaPorts    uint
	UserAgent     string

	LogLevel        zerolog.Level
	NoVad           bool
	ClockRate       uint
	Ptime           uint
	AudioFramePtime uint
	EcTailLen       uint
	EcOptions       uint
}

func NewConfig() Config {
	return Config{
		LogLevel: zerolog.DebugLevel,
		MaxCalls: 1024,
		//Threads:  uint(halfCPUs()),
		Threads:  1,

		BindAddress:   "",
		PublicAddress: "",
		Port:          0,
		UserAgent:     "pjsip",

		MediaPorts:      4096,
		NoVad:           false,
		ClockRate:       16000,
		Ptime:           20,
		AudioFramePtime: 20,
		EcTailLen:       100,
		EcOptions: uint(pjsua2.PJMEDIA_ECHO_WEBRTC) |
			uint(pjsua2.PJMEDIA_ECHO_USE_NOISE_SUPPRESSOR) |
			uint(pjsua2.PJMEDIA_ECHO_AGGRESSIVENESS_AGGRESSIVE),
	}
}

func halfCPUs() int {
	c := runtime.NumCPU()
	if c <= 1 {
		return 1
	}
	return c / 2
}
