package main

import "testing"

func TestHello(t *testing.T) {
	languages := Languages()

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("chris")
		want := "hello, chris"

		assertMessages(t, got, want)
	})

	t.Run("saying world to empty", func(t *testing.T) {
		got := Hello("")
		want := "hello, world"

		assertMessages(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		name := "Tom"
		got := Hello(name, languages.spanish)
		want := "hello in Spanish, " + name

		assertMessages(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		name := "Tom"
		got := Hello(name, languages.french)
		want := "hello in French, " + name

		assertMessages(t, got, want)
	})

}

func assertMessages(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
