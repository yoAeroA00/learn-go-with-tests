package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func ExampleRepeat() {
	ret := Repeat("a")
	fmt.Println(ret)
	// Output: aaaaa
}
