package main

import (
	"fmt"
	"net/http"
)

func (c *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	h1 := fmt.Sprintf("<h1>Welcome, Chirpy Admin</h1>")
	p := fmt.Sprintf("<p>Chirpy has been visited %d times!</p>", c.fileServerHits)
	// w.Write([]byte(fmt.Sprintf("Hits: %d", c.fileServerHits)))
	w.Write([]byte(h1))
	w.Write([]byte(p))
}

func (c * apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.fileServerHits++
		fmt.Println("Hits %s", c.fileServerHits)
		next.ServeHTTP(w, r)
	})
}
