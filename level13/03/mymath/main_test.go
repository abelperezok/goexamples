package mymath

import (
	"fmt"
	"testing"
)

type dataset struct {
	output float64
	input  []int
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 4, 6, 8, 100}
	c := CenteredAvg(xi)
	fmt.Println(c)
	// Output:
	// 6
}

func TestCenteredAvg(t *testing.T) {
	data := []dataset{
		{2, []int{0, 1, 2, 3, 100}},
		{6, []int{1, 4, 6, 8, 100}},
		{9, []int{0, 8, 10, 1000}},
		{9, []int{9000, 4, 10, 8, 6, 12}},
		{170, []int{123, 744, 140, 200}},
	}

	for i := 0; i < len(data); i++ {
		a := CenteredAvg(data[i].input)
		if a != data[i].output {
			t.Error("expected", data[i].output, "got", a)
		}
	}
}
