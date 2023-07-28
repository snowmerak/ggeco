package bean

import "sync"

type Default struct {
	beans map[any]any
	lock  sync.RWMutex
}

func (d *Default) Get(key any) (value any, ok bool) {
	d.lock.RLock()
	value, ok = d.beans[key]
	d.lock.RUnlock()
	return
}

func (d *Default) Set(key any, value any) {
	d.lock.Lock()
	d.beans[key] = value
	d.lock.Unlock()
}

func (d *Default) Delete(key any) {
	d.lock.Lock()
	delete(d.beans, key)
	d.lock.Unlock()
}

func (d *Default) Keys() (keys []any) {
	d.lock.RLock()
	keys = make([]any, 0, len(d.beans))
	for key := range d.beans {
		keys = append(keys, key)
	}
	d.lock.RUnlock()
	return
}

type Container interface {
	Get(key any) (value any, ok bool)
	Set(key any, value any)
	Delete(key any)
	Keys() (keys []any)
}

func NewContainer() (container *Default) {
	container = &Default{
		beans: make(map[any]any),
	}
	return
}

func CopyBeanContainer(dst Container, src Container) {
	for _, key := range src.Keys() {
		value, ok := src.Get(key)
		if !ok {
			continue
		}
		dst.Set(key, value)
	}
}
