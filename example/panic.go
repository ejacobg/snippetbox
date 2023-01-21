package example

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	// ...

	// Goroutine to do some background work:
	go func() {
		// In case this goroutine starts to panic:
		defer func() {
			if err := recover(); err != nil {
				log.Println(fmt.Errorf("%s\n%s", err, debug.Stack()))
			}
		}()

		// Do your background work...
	}()

	w.Write([]byte("OK"))
}
