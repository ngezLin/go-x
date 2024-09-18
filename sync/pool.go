package sync

import (
	"sync"
)

// Reseter define a contract to reset an instance of a type.
// if the type that the pool hold can fulfill the contract
// it the Reset method will be called before the instance is
// being put back to the internal Pool
type Reseter interface {
	Reset()
}

// Pool is a wrapper to sync.Poll. it use generic feature
// available starting from go 1.18 to limit what the content of the pool
// T is the type the pool is storing while R is the result;
// result will be described later.
//
// The main difference bwtween this and sync.Pool is the user
// does not need to handle object management, borrowing and returning
// it to the pool. Thus preventing the
type Pool[T, R any] struct {
	p sync.Pool
}

func New[T, R any](producer func() T) *Pool[T, R] {
	return &Pool[T, R]{
		p: sync.Pool{
			New: func() any {
				return producer()
			},
		},
	}
}

func (p *Pool[T, R]) DoWith(fn func(T) R) R {
	v := p.p.Get()
	t := v.(T)
	defer func() {
		if r, ok := v.(Reseter); ok {
			r.Reset()
		}
		p.p.Put(v)
	}()
	return fn(t)
}
