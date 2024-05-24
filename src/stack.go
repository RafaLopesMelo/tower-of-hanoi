package main

type Stack struct {
	data []int
	limit int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, true
	}

	index := len(s.data) - 1

	element := s.data[index]
	s.data = s.data[:index]

	return element, false
}

func (s *Stack) Size() int {
	size := 0

	for _, element := range s.data {
		if s.limit > element {
			size = size + 1
		}
	}

	return size
}

func (s *Stack) Decrease() {
	s.limit = s.limit - 1
}

func (s *Stack) Increase() {
	s.limit = s.limit + 1
}

func NewStack(limit int) *Stack {
	s := Stack{
		data: []int{},
		limit: limit,
	}
	return &s
}
