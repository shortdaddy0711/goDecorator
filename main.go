package main

import (
	"log"
	"net/http"
	"time"

	"github.com/shortdaddy0711/decorator/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time: ", time.Since(start).Milliseconds())
}

// NewHandler - empty handler
func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	return mux
}

func main() {
	mux := NewHandler()

	log.Fatal(http.ListenAndServe(":3000", mux))
}
