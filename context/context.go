package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	// Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		// data := make(chan string, 1)

		// go func() {
		// 	data <- store.Fetch()
		// }()

		// select {
		// case <-ctx.Done():
		// 	store.Cancel()
		// case d := <-data:
		// 	fmt.Fprint(w, d)
		// }

		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
