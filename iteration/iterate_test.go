package iteration

import (
	"fmt"
	"testing"
)

func Test_Iterate(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}

func Test_Again(t *testing.T) {
	got := Repeat("df", 2)
	want := "dfdf"

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	output := Repeat("Hi!", 4)
	fmt.Println(output)
	// Output: Hi!Hi!Hi!Hi!
}
