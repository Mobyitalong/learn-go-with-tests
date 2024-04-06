package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Run once", func(t *testing.T) {
		actual := Repeat("a", 1)
		expected := "a"

		assertResult(t, actual, expected)

	})

	t.Run("Run 5 times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"

		assertResult(t, actual, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}

func assertResult(t testing.TB, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Result of test: %q, does not match expected value: %q", actual, expected)
	}
}

func ExampleRepeat() {
	var repeated string = Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
