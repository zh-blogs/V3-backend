package set

type mapSet struct {
	set map[interface{}]struct{}
}

// NewMapSet ...
func NewMapSet(items ...interface{}) Set {
	s := new(mapSet)
	s.Clear()

	for _, item := range items {
		s.set[item] = struct{}{}
	}
	return s
}

// Len ...
func (s *mapSet) Len() int { return len(s.set) }

// Contains ...
func (s *mapSet) Contains(item interface{}) bool {
	_, exist := s.set[item]
	return exist
}

// Add ...
func (s *mapSet) Add(items ...interface{}) {
	for _, item := range items {
		s.set[item] = struct{}{}
	}
}

// Remove ...
func (s *mapSet) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.set, item)
	}
}

// Clear ...
func (s *mapSet) Clear() { s.set = make(map[interface{}]struct{}) }

// Equal compare two set is equal
func (s *mapSet) Equal(r Set) bool {
	if s.Len() != r.Len() {
		return false
	}
	for item := range s.set {
		if !r.Contains(item) {
			return false
		}
	}
	return true
}

// Copy ...
func (s *mapSet) Copy() Set {
	ns := NewMapSet()
	for item := range s.set {
		ns.Add(item)
	}
	return ns
}

// ForEach ...
func (s *mapSet) ForEach(callback func(item interface{}) error) error {
	for item := range s.set {
		if err := callback(item); err != nil {
			return err
		}
	}

	return nil
}

// Union ...
func (s *mapSet) Union(r Set) Set {
	ns := NewMapSet()

	s.ForEach(func(item interface{}) error {
		ns.Add(item)
		return nil
	})

	r.ForEach(func(item interface{}) error {
		ns.Add(item)
		return nil
	})

	return ns
}

// Intersect ...
func (s *mapSet) Intersect(r Set) Set {
	ns := NewMapSet()
	for item := range s.set {
		if r.Contains(item) {
			ns.Add(item)
		}
	}
	return ns
}

// Difference ...
func (s *mapSet) Difference(r Set) Set {
	ns := s.Copy()
	r.ForEach(func(item interface{}) error {
		ns.Remove(item)
		return nil
	})
	return ns
}
