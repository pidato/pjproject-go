package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"io"
	"os"
	"sync"
	"unsafe"
)

type RecorderListener interface {
	io.Closer

	// Invoked from a C++ worker thread. AudioFrame must be copied.
	OnFrame(frame *pjsua2.AudioFrame)
}

// Records PCM and Opus.
type recorder struct {
	base pjsua2.PiRecorder

	listeners []RecorderListener
	closed    bool
	mu        sync.Mutex
}

func newRecorder(ln ...RecorderListener) (*recorder, error) {
	r := &recorder{
		listeners: ln,
	}
	err := exec(func() {
		r.base = pjsua2.NewDirectorPiRecorder(r)
		r.base.Create()
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *recorder) AddListener(ln RecorderListener) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return os.ErrClosed
	}
	if ln == nil {
		return os.ErrInvalid
	}
	r.listeners = append(r.listeners, ln)
	return nil
}

func (r *recorder) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return os.ErrClosed
	}
	r.closed = true

	err := exec(func() {
		pjsua2.DeleteDirectorPiRecorder(r.base)
		r.base = nil
	})

	return err
}

func (r *recorder) OnHeartbeat() {
	//fmt.Println("onHeartbeat")
}

func (r *recorder) OnError(err pjsua2.Error) {
	Errorf("Err: %s\n", err.GetReason())
}

func (r *recorder) OnFrame(framePtr uintptr, prevExternNanos uint64) {
	// Cast
	frame := (*pjsua2.AudioFrame)(unsafe.Pointer(framePtr))

	r.mu.Lock()
	defer r.mu.Unlock()

	_ = exec(func() {
		if r.listeners != nil {
			for _, ln := range r.listeners {
				_ = exec(func() {
					ln.OnFrame(frame)
				})
			}
		}
	})
}
