package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assert := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("Should return a string default length which is 5", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assert(t, repeated, expected)
	})

	t.Run("Should return a string with length equal to times parameter", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assert(t, repeated, expected)
	})
}

// Benchmarking feature
/*
	goos: linux
	goarch: amd64
	pkg: iteration
	cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	BenchmarkRepeat-8        9551298               111.5 ns/op
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Printf("%q", repeated)
	// Output: "aaaaaaaaaa"
}
