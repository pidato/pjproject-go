package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"sync"
)

type Audio struct {
	call     *call             // Associated call if available.
	media    pjsua2.AudioMedia // Media
	inbound  []*Audio          // Receiving list of senders.
	outbound []*Audio          // Sending list of receivers.

	transmitting bool

	wg     sync.WaitGroup
	closed bool
	mu     sync.Mutex
}

func NewAudio(media pjsua2.AudioMedia) *Audio {
	audio := &Audio{
		media:    media,
		inbound:  nil,
		outbound: nil,
		wg:       sync.WaitGroup{},
		closed:   false,
		mu:       sync.Mutex{},
	}
	audio.wg.Add(1)
	return audio
}

// Ensures all transmissions are stopped.
func (p *Audio) Close() error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		p.wg.Wait()
		return nil
	}

	p.closed = true
	receiving := p.inbound
	p.inbound = nil
	sending := p.outbound
	p.outbound = nil
	p.mu.Unlock()

	for _, recv := range sending {
		recv.removeReceiving(p)
		p.media.StopTransmit(recv.media)
		recv.media.StopTransmit(p.media)
	}

	for _, send := range receiving {
		send.removeSending(p)
	}
	return nil
}

func (p *Audio) Stop() {
	defer func() {
		_ = recover()
	}()

	p.mu.Lock()
	defer p.mu.Unlock()
}

func (p *Audio) StopSending(recv *Audio) bool {
	defer func() {
		_ = recover()
	}()

	p.mu.Lock()
	if p.removeSending(recv) {
		recv.removeReceiving(p)
		p.mu.Unlock()
		p.media.StopTransmit(recv.media)
		return true
	}
	p.mu.Unlock()
	return false
}

func (p *Audio) StopReceiving(sender *Audio) {
	defer func() {
		_ = recover()
	}()
	sender.StopSending(p)
}

func (p *Audio) removeSending(recv *Audio) bool {
	var ok bool
	p.outbound, ok = removeFromSlice(p.inbound, recv)
	return ok
}

func (p *Audio) removeReceiving(sender *Audio) bool {
	var ok bool
	p.inbound, ok = removeFromSlice(p.inbound, sender)
	return ok
}

func removeFromSlice(slice []*Audio, value *Audio) ([]*Audio, bool) {
	var i int
	var receiver *Audio
	for i, receiver = range slice {
		if receiver == value {
			break
		}
	}

	if receiver == nil {
		return slice, false
	}

	switch len(slice) {
	case 0:
	case 1:
		slice = nil
	case 2:
		if i == 0 {
			slice[0] = slice[1]
		}
		slice[1] = nil
		slice = slice[0:1]
	default:
		for x := 0; x < len(slice); x++ {
			// Skip removal index.
			if x == i {
				continue
			}
			// Downshift.
			if x > i {
				slice[x-1] = slice[x]
			}
		}
		slice[len(slice)-1] = nil
		slice = slice[0 : len(slice)-1]
	}
	return slice, true
}
