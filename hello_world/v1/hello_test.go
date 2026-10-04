package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Adi")
	want := "Hello, Adi!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
