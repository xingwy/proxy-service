package set

type Set struct {
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		data: make(map[string]struct{}),
	}
}

func (s *Set) Add(items ...string) {
	for _, item := range items {
		s.data[item] = struct{}{}
	}
}

func (s *Set) Remove(item string) {
	delete(s.data, item)
}

func (s *Set) Exist(item string) bool {
	_, ok := s.data[item]
	return ok
}
