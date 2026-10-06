package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUrls(t *testing.T) {
	t.Run("normal implementation", func(t *testing.T) {
		slowURL := makeDelayedServer(time.Millisecond * 20)
		defer slowURL.Close()

		fastURL := makeDelayedServer(time.Millisecond * 10)
		defer fastURL.Close()

		want := fastURL.URL
		got := RacerOld(slowURL.URL, fastURL.URL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowURL := makeDelayedServer(time.Millisecond * 20)
		defer slowURL.Close()

		fastURL := makeDelayedServer(time.Millisecond * 10)
		defer fastURL.Close()

		want := fastURL.URL
		got, err := Racer(slowURL.URL, fastURL.URL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowURL := makeDelayedServer(time.Second * 12)
		defer slowURL.Close()

		fastURL := makeDelayedServer(time.Second * 11)
		defer fastURL.Close()

		_, err := Racer(slowURL.URL, fastURL.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
