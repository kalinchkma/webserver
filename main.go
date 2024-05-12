package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileServerHits int
}


func main() {
	const port = "8080"
	const filePath = "."
	const assets = "./assets"

	apiCfg := apiConfig{
		fileServerHits: 0,
	}

	mux := http.NewServeMux()

	mux.Handle("GET /app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePath)))))
	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("/reset", apiCfg.handlerReset)

	// http server
	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Server on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

