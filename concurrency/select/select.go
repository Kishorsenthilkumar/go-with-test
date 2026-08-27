package _select

import (
	"net/http"
	"time"
)

func Racer(slowurl, fasturl string) (winurl string) {
	slowurl_dur:=measureResponseTime(slowurl)
	fasturl_dur:=measureResponseTime(fasturl)

	if fasturl_dur<slowurl_dur{
		return fasturl
	}
	return slowurl
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	resp, err := http.Get(url)
	if err == nil {
		resp.Body.Close()
	}
	return time.Since(start)
}