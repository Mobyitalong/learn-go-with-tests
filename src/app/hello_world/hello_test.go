package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("English", func(t *testing.T) {
		actual := Hello("Khaleel", "English")
		expected := "Hello, Khaleel"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Empty Strings", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, world"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Spanish", func(t *testing.T) {
		actual := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("French", func(t *testing.T) {
		actual := Hello("Pierre", "French")
		expected := "Bonjour, Pierre"

		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}
