package set

import "encoding/json"

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	s.Add(items...)
	return s
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) ContainsAll(items ...T) bool {
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}

	return true
}

func (s Set[T]) ContainsAny(items ...T) bool {
	for _, item := range items {
		if s.Contains(item) {
			return true
		}
	}

	return false
}

func (s Set[T]) Clone() Set[T] {
	new := make(Set[T], len(s))
	for item := range s {
		new.Add(item)
	}
	return new
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for item := range s {
		slice = append(slice, item)
	}
	return slice
}

func (s Set[T]) Equals(other Set[T]) bool {
	if s == nil && other == nil {
		return true
	}

	if len(s) != len(other) {
		return false
	}

	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}

	return true
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	slice := []T{}
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	*s = New(slice...)

	return nil
}
