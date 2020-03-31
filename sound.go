package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"os"
	"reflect"
	"sync"
	"unsafe"

	_ "github.com/pidato/audio/opus"
)

var emptyFrame = make([]int16, 10000)

type Playlist struct {
}

type Sound struct {
	Name  string
	Clips []*Clip
}

type Clip struct {
	Sound *Sound
	PCM   []int16
}

type PlaylistPlayer struct {
	port pjsua2.PiPort

	clip            *Clip
	frame           int
	samplesPerFrame int

	closed bool
	mu     sync.Mutex
}

func newPlaylistPlayer() (*PlaylistPlayer, error) {
	p := &PlaylistPlayer{}
	err := exec(func() {
		p.port = pjsua2.NewDirectorPiPort(p)
		p.port.Create()
		p.samplesPerFrame = int(p.port.GetSamplesPerFrame())
	})
	if err != nil {
		if p.port != nil {
			_ = exec(func() {
				pjsua2.DeleteDirectorPiPort(p.port)
			})
		}
		return nil, err
	}
	return p, nil
}

func (p *PlaylistPlayer) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return os.ErrClosed
	}

	p.closed = true
	return exec(func() {

	})
}

func (p *PlaylistPlayer) OnGetFrame(arg2 pjsua2.Pjmedia_frame_type, data uintptr, length int64, arg5 uint64, arg6 uint) {
	if arg2 != pjsua2.PJMEDIA_FRAME_TYPE_AUDIO {
		return
	}

	slice := &reflect.SliceHeader{
		Data: data,
		Len:  int(length),
		Cap:  int(length),
	}

	pjframe := *(*[]int16)(unsafe.Pointer(slice))

	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.clip == nil {
		if slice.Len > len(emptyFrame) {
			return
		}
		// Zero it out.
		copy(pjframe, emptyFrame[0:len(pjframe)])
		return
	}

	//
}
