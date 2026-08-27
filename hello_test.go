package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {

		got := Hello("Chris")
		want := "Hello Chriss!"

		assertCorrectMessage(t, want, got)

	})

	t.Run("test 2", func(t *testing.T) {

		got := Hello("")
		want := "Hello xxx!"

		assertCorrectMessage(t, want, got)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
