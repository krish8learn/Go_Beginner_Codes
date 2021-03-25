package main

import "testing"

func TestMysum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 3, 4}, 9},
		test{[]int{2, 1, 1}, 4},
		test{[]int{1, 4, 1}, 6},
	}

	for _, v := range tests {
		x := Mysum(v.data...)
		if x != v.answer {
			t.Error("Expected:", v.answer, "go:", x)
		}
	}
}
