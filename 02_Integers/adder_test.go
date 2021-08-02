package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(4, 5)
	fmt.Println(sum)
	// Output: 9
}

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	}

	t.Run("Adds two numbers and return the result", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)
	})
}
