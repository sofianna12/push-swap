package stack

func Sa(a *Stack) string {
	if a.Len() < 2 {
		return ""
	}
	a.data[0], a.data[1] = a.data[1], a.data[0]
	return "sa"
}

func Sb(b *Stack) string {
	if b.Len() < 2 {
		return ""
	}
	b.data[0], b.data[1] = b.data[1], b.data[0]
	return "sb"
}

func Ss(a, b *Stack) string {
	Sa(a)
	Sb(b)
	return "ss"
}

func Pa(a, b *Stack) string {
	val, ok := b.Pop()
	if !ok {
		return ""
	}
	a.Push(val)
	return "pa"
}

func Pb(a, b *Stack) string {
	val, ok := a.Pop()
	if !ok {
		return ""
	}
	b.Push(val)
	return "pb"
}

func Ra(a *Stack) string {
	if a.Len() < 2 {
		return ""
	}
	first := a.data[0]
	a.data = append(a.data[1:], first)
	return "ra"
}

func Rb(b *Stack) string {
	if b.Len() < 2 {
		return ""
	}
	first := b.data[0]
	b.data = append(b.data[1:], first)
	return "rb"
}

func Rr(a, b *Stack) string {
	Ra(a)
	Rb(b)
	return "rr"
}

func Rra(a *Stack) string {
	if a.Len() < 2 {
		return ""
	}
	last := a.data[len(a.data)-1]
	a.data = append([]int{last}, a.data[:len(a.data)-1]...)
	return "rra"
}

func Rrb(b *Stack) string {
	if b.Len() < 2 {
		return ""
	}
	last := b.data[len(b.data)-1]
	b.data = append([]int{last}, b.data[:len(b.data)-1]...)
	return "rrb"
}

func Rrr(a, b *Stack) string {
	Rra(a)
	Rrb(b)
	return "rrr"
}
