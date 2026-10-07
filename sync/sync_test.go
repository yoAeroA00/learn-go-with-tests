package sync

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {
	t.Run("increamenting the counter 3 times leaves it at 3", func(t *testing.T) {
		// counter := Counter{}
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		want := 1000
		// counter := Counter{}
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, want)
	})

	t.Run("it runs safely concurrently using atomic", func(t *testing.T) {
		var want int64 = 1000
		counter := ACounter{}

		var wg sync.WaitGroup

		for i := int64(0); i < want; i++ {
			wg.Add(1)
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertACounter(t, &counter, want)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	if (*counter).Value() != want {
		t.Errorf("got %d, want %d", (*counter).Value(), want)
	}
}

func assertACounter(t testing.TB, counter *ACounter, want int64) {
	t.Helper()
	if (*counter).Value() != want {
		t.Errorf("got %d, want %d", (*counter).Value(), want)
	}
}
