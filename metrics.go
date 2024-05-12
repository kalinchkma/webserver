package main

import (
	"fmt"
	"net/http"
)

// metrics functions

func (c *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", c.fileServerHits)))
}

func (c * apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.fileServerHits++
		fmt.Println("Hits %s", c.fileServerHits)
		next.ServeHTTP(w, r)
	})
}
