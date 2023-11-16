package main

import "testing"

// TestFactorialPart tests the factorialPart function for correctness
func TestFactorialPart(t *testing.T) {
	tests := []struct {
		start, end, expected int
	}{
		{1, 5, 120}, // 5! = 120
		{6, 10, 30240},
	}

	for _, test := range tests {
		ch := make(chan int)
		go factorialPart(test.start, test.end, ch)
		result := <-ch
		if result != test.expected {
			t.Errorf("factorialPart(%d, %d) = %d; expected %d", test.start, test.end, result, test.expected)
		}
	}
}

func BenchmarkFactorialPart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go factorialPart(1, 10, ch)
		<-ch
	}
}
