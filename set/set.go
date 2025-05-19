package set

type Set interface {
	Insert(d any)
	Remove(d any)
	Contains(d any) bool
	Size() int
	IsEmpty() bool
	GetAll() []any
}

type SetImpl struct {
	data map[any]struct{}
}

func New() Set {
	return &SetImpl{
		data: make(map[any]struct{}),
	}
}

func (s *SetImpl) Insert(d any) {
	s.data[d] = struct{}{}
}

func (s *SetImpl) Remove(d any) {
	delete(s.data, d)
}

func (s *SetImpl) Contains(d any) bool {
	_, exists := s.data[d]
	return exists
}

func (s *SetImpl) Size() int {
	return len(s.data)
}

func (s *SetImpl) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *SetImpl) GetAll() []any {
	all := make([]any, 0, len(s.data))
	for k := range s.data {
		all = append(all, k)
	}
	return all
}
