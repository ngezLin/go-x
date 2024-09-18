package sync

import "sync"

type Map[K, V any] interface {
	Load(K) (V, bool)
	Store(K, V)
	LoadOrStore(K, V) (actual V, loaded bool)
	Delete(K)
	Range(func(K, V) bool)
}

type SyncMap[K any, V any] struct {
	m sync.Map
}

var _ Map[any, any] = (*SyncMap[any, any])(nil)

func (sm *SyncMap[K, V]) Load(key K) (v V, ok bool) {
	_v, ok := sm.m.Load(key)
	if !ok {
		return
	}
	v = _v.(V)
	return
}

func (sm *SyncMap[K, V]) Store(k K, v V) {
	sm.m.Store(k, v)
}

func (sm *SyncMap[K, V]) LoadOrStore(k K, v V) (actual V, loaded bool) {
	_v, loaded := sm.m.LoadOrStore(k, v)
	if !loaded {
		actual = v
		return
	}
	actual = _v.(V)
	return
}

func (sm *SyncMap[K, V]) LoadAndDelete(k K) (actual V, loaded bool) {
	_v, loaded := sm.m.LoadAndDelete(k)
	if !loaded {
		return
	}
	actual = _v.(V)
	return
}

func (sm *SyncMap[K, V]) Delete(k K) {
	sm.m.Delete(k)
}

func (sm *SyncMap[K, V]) Range(f func(K, V) bool) {
	sm.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

type MutexMap[K comparable, V any] struct {
	m    map[K]V
	mutx sync.RWMutex
}

var _ Map[int, any] = (*MutexMap[int, any])(nil)

func (mm *MutexMap[K, V]) Load(k K) (v V, ok bool) {
	mm.mutx.RLock()
	defer mm.mutx.Unlock()
	v, ok = mm.m[k]
	return
}

func (mm *MutexMap[K, V]) Store(k K, v V) {
	mm.mutx.Lock()
	defer mm.mutx.Unlock()
	mm.m[k] = v
}

func (mm *MutexMap[K, V]) LoadOrStore(k K, v V) (V, bool) {
	mm.mutx.RLock()
	if ev, ok := mm.m[k]; ok {
		mm.mutx.Unlock()
		return ev, true
	}
	mm.mutx.Unlock()
	mm.mutx.Lock()
	mm.m[k] = v
	mm.mutx.Unlock()
	return v, false
}

func (mm *MutexMap[K, V]) LoadAndDelete(k K) (actual V, loaded bool) {
	mm.mutx.RLock()
	ev, ok := mm.m[k]
	if !ok {
		mm.mutx.Unlock()
		loaded = ok
		return
	}
	// key exist
	actual = ev
	mm.mutx.Unlock()
	mm.Delete(k)
	return
}

func (mm *MutexMap[K, V]) Delete(k K) {
	mm.mutx.Lock()
	delete(mm.m, k)
	mm.mutx.Unlock()
}

func (mm *MutexMap[K, V]) Range(f func(K, V) bool) {
	mm.mutx.RLock()
	m := mm.m
	mm.mutx.Unlock()
	for k, v := range m {
		if _continue := f(k, v); !_continue {
			return
		}
	}
}
