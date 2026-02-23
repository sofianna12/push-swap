package stack

type Stack struct {
	data []int
}

func NewStack(values []int) *Stack {
	cp := make([]int, len(values))
	copy(cp, values)
	return &Stack{data: cp}
}

func (s *Stack) Push(v int) {
	s.data = append([]int{v}, s.data...)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	val := s.data[0]
	s.data = s.data[1:]
	return val, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[0], true
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Values() []int {
	cp := make([]int, len(s.data))
	copy(cp, s.data)
	return cp
}
