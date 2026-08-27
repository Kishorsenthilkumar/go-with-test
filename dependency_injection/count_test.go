package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDown(buffer, &SpyCountDownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {

		spysleepprinter := &SpyCountDownOperations{}
		CountDown(spysleepprinter, spysleepprinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spysleepprinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spysleepprinter.Calls)
		}

	})
}
