package sort

import (
	"math/rand"
	"testing"

	"push-swap/internal/stack"
)

func BenchmarkSort5(b *testing.B) {
	nums := []int{5, 1, 4, 2, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := stack.New("a", append([]int(nil), nums...))
		bStack := stack.New("b", nil)
		SortCollect(a, bStack)
	}
}

func BenchmarkSort6(b *testing.B) {
	nums := []int{2, 1, 3, 6, 5, 8}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := stack.New("a", append([]int(nil), nums...))
		bStack := stack.New("b", nil)
		SortCollect(a, bStack)
	}
}

func BenchmarkSort100(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	nums := r.Perm(1000)[:100]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := stack.New("a", append([]int(nil), nums...))
		bStack := stack.New("b", nil)
		SortCollect(a, bStack)
	}
}

func BenchmarkSort500(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	nums := r.Perm(10000)[:500]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := stack.New("a", append([]int(nil), nums...))
		bStack := stack.New("b", nil)
		SortCollect(a, bStack)
	}
}
