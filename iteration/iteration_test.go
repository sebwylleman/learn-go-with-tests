package iteration

import "testing"

func TestRepeatFiveTimes(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"
	if repeated != expected {
		t.Errorf("want %q, got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
