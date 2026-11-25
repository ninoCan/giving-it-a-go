package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Given a name, greet that name", func(t *testing.T) {
		expected := "Hello, Nino!"

		actual := Hello("Nino")

		assertEqual(t, actual, expected)
	})
	t.Run("Given empty string, greet the world", func(t *testing.T) {
		expected := "Hello, World!"

		actual := Hello("")

		assertEqual(t, actual, expected)

	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
