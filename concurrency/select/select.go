package _select

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(slowurl, fasturl string) (winurl string, err error) {
	select {
	case <-ping(slowurl):
		return slowurl, nil
	case <-ping(fasturl):
		return fasturl, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timeout for request %s and %s ", slowurl, fasturl)
	}

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		resp, err := http.Get(url)

		if err == nil {
			resp.Body.Close()
		}
		close(ch)
	}()
	return ch
}
