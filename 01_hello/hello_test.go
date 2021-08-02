package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMesage := func(t testing.TB, got, want string) {
		// Helper marks the calling function as a test helper function.
		// When printing file and line information, that function will be skipped.
		// Helper may be called simultaneously from multiple goroutines.
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Juan", "")
		want := "Hello, Juan"
		assertCorrectMesage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMesage(t, got, want)
	})

	t.Run("Saying Hello in spanish", func(t *testing.T) {
		got := Hello("Juan", "ES")
		want := "Hola, Juan"
		assertCorrectMesage(t, got, want)
	})

	t.Run("Saying Hello in french", func(t *testing.T) {
		got := Hello("Juan", "FR")
		want := "Bonjour, Juan"
		assertCorrectMesage(t, got, want)
	})

	t.Run("Saying Hello in italian", func(t *testing.T) {
		got := Hello("Juan", "IT")
		want := "Ciao, Juan"
		assertCorrectMesage(t, got, want)
	})
}
