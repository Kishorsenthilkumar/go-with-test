package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const sleep = "sleep"
const write = "write"

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")

}

func main() {
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
