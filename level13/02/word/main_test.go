package word

import (
	"fmt"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(string(i))
	}
}

func ExampleCount() {
	s := "a word another"
	c := Count(s)
	fmt.Println(c)
	// Output:
	// 3
}

func TestCount(t *testing.T) {

	s := "a word another"
	c := Count(s)

	if c != 3 {
		t.Error("expected", 3, "got", c)
	}
}

func TestUseCount(t *testing.T) {

	s := "aa ab ac aa"
	m := UseCount(s)

	if m["aa"] != 2 {
		t.Error("expected aa", 2, "got", m["aa"])
	}

	if m["ab"] != 1 {
		t.Error("expected aa", 1, "got", m["ab"])
	}
	if m["ac"] != 1 {
		t.Error("expected aa", 1, "got", m["ac"])
	}
}
