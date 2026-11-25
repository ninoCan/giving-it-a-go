package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Given a name, greet that name", func(t *testing.T) {
		got := Hello("Nino")
		want := "Hello, Nino!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Given empty string, greet the world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
