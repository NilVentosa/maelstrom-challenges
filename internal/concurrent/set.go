package concurrent

import (
	"maps"
	"slices"
	"sync"
)

type Set[T comparable] struct {
	m sync.RWMutex
	d map[T]struct{}
}

func NewConcurrentSet[T comparable]() Set[T] {
	d := make(map[T]struct{})
	return Set[T]{sync.RWMutex{}, d}
}

func (s *Set[T]) Contains(item T) bool {
	s.m.RLock()
	defer s.m.RUnlock()
	_, present := s.d[item]
	return present
}

func (s *Set[T]) Add(item T) {
	s.m.Lock()
	defer s.m.Unlock()
	s.d[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.d, item)
}

func (s *Set[T]) Values() []T {
	s.m.RLock()
	defer s.m.RUnlock()
	return slices.Collect(maps.Keys(s.d))
}
