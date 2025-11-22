package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	// subtests: Describe different scenarios around a function
	t.Run("Saying hello to person", func(t *testing.T) {
		got := Hello2("Soma")
		want := "Hello, Soma"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	// subtests
	t.Run("Saying 'Hello World' when empty string is supplied", func(t *testing.T) {
		got := Hello2("")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}

func TestMaths(t *testing.T) {
	got := maths(10.5, 10)
	var want float32 = 20.5

	if got != want {
		t.Errorf("got %f but want %f", got, want)
	}
}
