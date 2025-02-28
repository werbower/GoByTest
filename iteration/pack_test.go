package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		assertM(t, expected, repeated)
	})

}

func assertM(t *testing.T, expected, got string) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
