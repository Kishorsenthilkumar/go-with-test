package Sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter upto 3", func(t *testing.T) {

		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()
	})
	t.Run("concurrent counter", func(t *testing.T) {
		wantedcount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedcount)

		for i := 0; i < wantedcount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedcount)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
