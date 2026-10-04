package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Adi", "")
		want := "Hello, Adi!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world when empty string given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world in hindi", func(t *testing.T) {
		got := Hello("", "Hindi")
		want := "Namaste, Duniya!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
