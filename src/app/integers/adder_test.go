package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}

func ExampleAdd() {
	var sum int = Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
