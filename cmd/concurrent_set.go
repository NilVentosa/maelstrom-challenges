package main

import (
	"iter"
	"maps"
	"sync"
)

type ConcurrentSet[T comparable] struct {
	m sync.RWMutex
	d map[T]struct{}
}

func NewConcurrentSet[T comparable]() ConcurrentSet[T] {
	d := make(map[T]struct{})
	return ConcurrentSet[T]{sync.RWMutex{}, d}
}

func (s *ConcurrentSet[T]) Contains(item T) bool {
	s.m.RLock()
	defer s.m.RUnlock()
	_, present := s.d[item]
	return present
}

func (s *ConcurrentSet[T]) Add(item T) {
	s.m.Lock()
	defer s.m.Unlock()
	s.d[item] = struct{}{}
}

func (s *ConcurrentSet[T]) Values() iter.Seq[T] {
	s.m.RLock()
	defer s.m.RUnlock()
	return maps.Keys(s.d)
}
