package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"sync"
)

type Account struct {
	account pjsua2.Account

	calls map[string]*Call
	mu    sync.RWMutex
}

func (p *Account) removeCall(id string, call *Call) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.calls, id)
}
