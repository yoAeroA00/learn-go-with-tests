package racer

import (
	"fmt"
	"net/http"
	"time"
)

func RacerOld(slowUrl, fastUrl string) string {
	durationS := measureResponseTime(slowUrl)

	durationF := measureResponseTime(fastUrl)

	if durationS > durationF {
		return fastUrl
	}

	return slowUrl
}

func Racer(slowUrl, fastUrl string) (string, error) {
	select {
	case <-ping(slowUrl):
		return slowUrl, nil
	case <-ping(fastUrl):
		return fastUrl, nil
	case <-time.After(time.Second * 10):
		return "", fmt.Errorf("timed out waiting for %s and %s", slowUrl, fastUrl)
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	resp, err := http.Get(url)
	if err == nil {
		resp.Body.Close()
	}
	return time.Since(start)
}

func ping(url string) chan struct{} {
	result := make(chan struct{})

	go func() {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
		}
		close(result)
	}()

	return result
}
