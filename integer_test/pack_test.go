package integertest

import "testing"

func TestAdder(t *testing.T) {

	t.Run("test example", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 5

		if sum != expected {
			t.Errorf("expected '%d' but got '%d' ", expected, sum)
		}

	})

	t.Run("test example 2+3", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5

		if sum != expected {
			t.Errorf("expected '%d' but got '%d' ", expected, sum)
		}

	})

}
