package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
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
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("GET /api/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", func(w http.ResponseWriter, r *http.Request) {

		// decoding json
		type parameters struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}

		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)

		if err != nil {
			log.Printf("Error decoding parameters: %s", err)
			w.WriteHeader(500)
			return
		}

		// encoding json
		type returnVals struct {
			CreateAt time.Time `json:"created_at"`
			ID int `json:"id"`
		}

		respBody := returnVals {
			CreateAt: time.Now(),
			ID: 123,
		}

		dat, err := json.Marshal(respBody)

		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "appication/json")
		w.WriteHeader(200)
		w.Write(dat)
	})

	// http server
	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Server on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

