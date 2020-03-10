package main

import "C"
import (
	"fmt"
	"github.com/pidato/pjproject-go"
	"github.com/pidato/pjproject-go/pjsua2"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			err := e.(pjsua2.Error)
			if err != nil {
				fmt.Println(err.GetReason())
			}
		}
	}()

	fmt.Println(unsafe.Sizeof(pjsua2.AudioFrame{}))

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
		pjsua2.EndpointInstance().LibRegisterThread("")
	}

	player := pjsua2.NewAudioMediaPlayer()
	//player.CreatePlayer("recording.wav", uint(pjsua2.PJMEDIA_FILE_NO_LOOP))
	player.CreatePlayer("recording.wav")

	encoder := &PiEncoder{
		ch: make(chan uint64),
	}
	encoder.director = pjsua2.NewDirectorPiEncoder(encoder)
	encoder.director.Create()
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

type PiEncoder struct {
	//pjsua2.PiEncoder
	director pjsua2.PiEncoder

	directorCPU uint64
	speaking    bool

	ch chan uint64
}

func (enc *PiEncoder) OnHeartbeat() {
	fmt.Println("onHeartbeat")
}

func (enc *PiEncoder) OnError(err pjsua2.Error) {
	fmt.Printf("Err: %s\n", err.GetReason())
}

func (enc *PiEncoder) OnFrameDTX(framePointer uintptr, prevExternCPU uint64) {
	//frame := (*pjsua2.AudioFrame)(unsafe.Pointer(framePointer))
	fmt.Println(framePointer)
	//if frameNum > 0 && frameNum % 50 == 0 {
	//	fmt.Printf("DTX Frame: %d\n", frameNum)
	//	fmt.Printf("\tDirector CPU: %d\n", prevExternCPU)
	//}
	//enc.directorCPU += prevExternCPU
}

func (enc *PiEncoder) OnFrame(framePointer uintptr, prevExternCPU uint64) {
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
	//fmt.Println(frame.Vad)

	//if frame.FrameNum > 0 && frame.FrameNum%50 == 0 {
	//	fmt.Printf("Frame: %d Vad: %d  Opus: %d\n", frame.FrameNum, frame.Vad, frame.OpusSize)
	//	fmt.Printf("\tDirector CPU: %d\n", enc.directorCPU)
	//	//enc.ch <- enc.directorCPU
	//	enc.directorCPU = 0
	//}
	//enc.directorCPU += prevExternCPU
}
