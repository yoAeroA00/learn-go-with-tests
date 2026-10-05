package countdown

import (
	"bytes"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int64
}

type DSleeper struct{}

type Sleeper interface {
	Sleep(ms time.Duration)
}

func (s *SpySleeper) Sleep(ms time.Duration) {
	(*s).Calls++
}

func (ds *DSleeper) Sleep(ms time.Duration) {
	time.Sleep(ms)
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := (*buffer).String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
