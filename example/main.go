package main

import "C"
import (
	"fmt"
	"github.com/pidato/audio/opus"
	"github.com/pidato/pjproject-go"
	"github.com/pidato/pjproject-go/pjsua2"
	"github.com/rs/zerolog"
	"runtime"
	"time"
	"unsafe"

	_ "github.com/pidato/audio/opus"
	_ "github.com/pidato/vad-go"
)

func fibrec(n int) int {
	if n <= 1 {
		return n
	}
	return fibrec(n-1) + fibrec(n-2)
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			err := e.(pjsua2.Error)
			if err != nil {
				fmt.Println(err.GetReason())
			}
		}
	}()

	pj.SetLogLevel(zerolog.InfoLevel)
	pj.Infof("sizeof(AudioFrame) = %d -> %d", unsafe.Sizeof(pjsua2.AudioFrame{}), pjsua2.SizeofPiAudioFrame())

	cfg := pj.NewConfig()
	cfg.Ptime = 20
	cfg.AudioFramePtime = 20
	service, err := pj.Start(cfg)
	if err != nil {
		panic(err)
	}
	defer service.Close()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if !pjsua2.EndpointInstance().LibIsThreadRegistered() {
		pjsua2.EndpointInstance().LibRegisterThread("main")
	}

	player := pjsua2.NewAudioMediaPlayer()
	//player.CreatePlayer("recording.wav", uint(pjsua2.PJMEDIA_FILE_NO_LOOP))
	player.CreatePlayer("recording.wav")

	encoder := &Recorder{
		ch: make(chan uint64),
	}
	encoder.director = pjsua2.NewDirectorPiRecorder(encoder)
	encoder.director.Create()

	encoder.multiple = 48000 / int(encoder.director.GetClockRate())
	//out, err := opus.CreateFile("out.opus", uint32(encoder.director.GetClockRate()), 1)
	//if err != nil {
	//	panic(err)
	//}
	//encoder.writer = out
	//pjsua2.PiEncoderAddEncoderThreads(3)

	go func() {
		for msg := range encoder.ch {
			fmt.Printf("\tDirector CPU: %d\n", msg)
		}
	}()

	player.StartTransmit(encoder.director)

	time.Sleep(time.Minute * 60)
}

//export pidato_create
func pidato_create(path string) *C.char {
	return nil
}

type Recorder struct {
	//pjsua2.Recorder
	director pjsua2.PiRecorder

	directorCPU uint64
	speaking    bool

	writer *opus.OggWriter
	multiple int

	ch chan uint64
}

func (enc *Recorder) OnHeartbeat() {
	fmt.Println("onHeartbeat")
}

func (enc *Recorder) OnError(err pjsua2.Error) {
	fmt.Printf("Err: %s\n", err.GetReason())
}

func (enc *Recorder) OnFrame(framePointer uintptr, prevExternCPU uint64) {
	//fmt.Println(framePointer)
	frame := (*pjsua2.AudioFrame)(unsafe.Pointer(framePointer))

	if frame.Vad == 0 {
		if enc.speaking {
			fmt.Println("Not Speaking")
			enc.speaking = false
		}
	} else if frame.Vad == 1 {
		if !enc.speaking {
			fmt.Println("Speaking")
			enc.speaking = true
		}
	}

	//err := enc.writer.Write(frame.Opus[0:frame.OpusSize], int(frame.PcmSamples) * enc.multiple)
	//if err != nil {
	//	pj.Errorf("failed to write Opus file: %s", err.Error())
	//}
	//fmt.Println(frame.Vad)

	//if frame.FrameNum > 0 && frame.FrameNum%50 == 0 {
	//	fmt.Printf("Frame: %d Vad: %d  Opus: %d\n", frame.FrameNum, frame.Vad, frame.OpusSize)
	//	fmt.Printf("\tDirector CPU: %d\n", enc.directorCPU)
	//	//enc.ch <- enc.directorCPU
	//	enc.directorCPU = 0
	//}
	//enc.directorCPU += prevExternCPU
}
