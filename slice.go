package gosafe

import (
	"sync"
)

type Slice[T comparable] struct {
	sync.RWMutex
	slice []T
}

func (s *Slice[T]) LenUnsafe() int {
	return len(s.slice)
}

func (s *Slice[T]) Len() int {
	s.RLock()
	defer s.RUnlock()
	return s.LenUnsafe()
}

func (s *Slice[T]) GetUnsafe(i int) T {
	return s.slice[i]
}

func (s *Slice[T]) Get(i int) T {
	s.RLock()
	defer s.RUnlock()
	return s.GetUnsafe(i)
}

func (s *Slice[T]) AppendUnsafe(elems ...T) {
	s.slice = append(s.slice, elems...)
}

func (s *Slice[T]) Append(elems ...T) {
	s.Lock()
	defer s.Unlock()
	s.AppendUnsafe(elems...)
}

func (s *Slice[T]) SetUnsafe(i int, value T) {
	s.slice[i] = value
}

func (s *Slice[T]) Set(i int, value T) {
	s.Lock()
	defer s.Unlock()
	s.SetUnsafe(i, value)
}

// Preserves order
func (s *Slice[T]) DeleteStableUnsafe(i int) {
	s.slice = append(s.slice[:i], s.slice[i+1:]...)
}

// Preserves order
func (s *Slice[T]) DeleteStable(i int) {
	s.Lock()
	defer s.Unlock()
	s.DeleteStableUnsafe(i)
}

// Doesn't preserve order
func (s *Slice[T]) DeleteUnstableUnsafe(i int) {
	// swaps the ith element with the last one
	s.slice[i] = s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
}

// Doesn't preserve order
func (s *Slice[T]) DeleteUnstable(i int) {
	s.Lock()
	defer s.Unlock()
	s.DeleteStableUnsafe(i)
}

func (s *Slice[T]) InUnsafe(elem T) bool {
	for _, e := range s.slice {
		if e == elem {
			return true
		}
	}
	return false
}

func (s *Slice[T]) In(elem T) bool {
	s.RLock()
	defer s.RUnlock()
	return s.InUnsafe(elem)
}
