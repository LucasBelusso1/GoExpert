package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// // Solution using Mutex
		// m.Lock()
		// number++
		// m.Unlock()

		// // Solution using atomic
		atomic.AddUint64(&number, 1) // add 1 to variable number
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})
	http.ListenAndServe(":8000", nil)
}
