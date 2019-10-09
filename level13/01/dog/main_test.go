package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func TestYears(t *testing.T) {

	y := Years(1)

	if y != 7 {
		t.Error("expected", 7, "got", y)
	}
}

func ExampleYears() {

	y := Years(1)
	fmt.Println(y)

}

func BenchmarkYearsTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}
func TestYearsTwo(t *testing.T) {

	y := YearsTwo(1)

	if y != 7 {
		t.Error("expected", 7, "got", y)
	}
}

func ExampleYearsTwo() {

	y := YearsTwo(1)
	fmt.Println(y)

}
