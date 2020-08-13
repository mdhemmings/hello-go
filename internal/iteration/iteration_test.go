package iteration

import "testing"

// TestRepeat : Test the repeat statement works
func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("got %q wanted %q", repeated, expected)
		}
	}

	t.Run("1 repeat", func(t *testing.T) {
		repeats := 1
		gotRepeats := Repeat("a", repeats)
		wantRepeats := "a"
		assertCorrectMessage(t, gotRepeats, wantRepeats)
	})

	t.Run("5 repeat", func(t *testing.T) {
		repeats := 5
		gotRepeats := Repeat("a", repeats)
		wantRepeats := "aaaaa"
		assertCorrectMessage(t, gotRepeats, wantRepeats)
	})

	t.Run("5 numeric repeat", func(t *testing.T) {
		repeats := 5
		gotRepeats := Repeat("1", repeats)
		wantRepeats := "11111"
		assertCorrectMessage(t, gotRepeats, wantRepeats)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}
