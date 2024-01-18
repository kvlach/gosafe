package gosafe

import (
	"sync"
)

type Value[T any] struct {
	sync.RWMutex
	value T
}

func (v *Value[T]) SetUnsafe(value T) {
	v.value = value
}

func (v *Value[T]) Set(value T) {
	v.Lock()
	defer v.Unlock()
	v.SetUnsafe(value)
}

func (v *Value[T]) GetUnsafe() T {
	return v.value
}

func (v *Value[T]) Get() T {
	v.RLock()
	defer v.RUnlock()
	return v.GetUnsafe()
}
