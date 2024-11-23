package messaging

import (
	"sync"
)

type Promise struct {
	sync.Mutex
	err       error
	finished  bool
	callbacks []func(err error)
}

func NewPromise() *Promise {
	return &Promise{}
}

func (p *Promise) executeCallbacks() {
	if p.finished {
		return
	}
	for _, s := range p.callbacks {
		s(p.err)
	}
	p.finished = true
}

func (p *Promise) Then(callback func(err error)) *Promise {
	p.Lock()
	defer p.Unlock()

	if p.finished {
		callback(p.err)
	} else {
		p.callbacks = append(p.callbacks, callback)
	}
	return p
}

func (p *Promise) Finish(err error) *Promise {
	p.Lock()
	defer p.Unlock()

	p.err = err

	p.executeCallbacks()
	return p
}
