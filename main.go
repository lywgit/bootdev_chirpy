package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		log.Printf("current hit cnt %d", cfg.fileserverHits.Load())
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) showMetricsHandlerFunc() http.HandlerFunc {
	writeHitCnt := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("current hit cnt %d", cfg.fileserverHits.Load())
		s := fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())
		w.Write([]byte(s))
		w.WriteHeader(http.StatusOK)
	}
	return writeHitCnt
}

func (cfg *apiConfig) resetMetricsHandlerFunc() http.HandlerFunc {
	resetHitCnt := func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Store(0)
		log.Printf("reset hit cnt to %d", cfg.fileserverHits.Load())
		w.WriteHeader(http.StatusOK)
	}
	return resetHitCnt
}

func main() {
	const filepathRoot = "."
	const port = "8080"
	apiCfg := &apiConfig{} // remember to use ptr

	mux := http.NewServeMux()

	fileHandler := http.FileServer(http.Dir(filepathRoot))
	fileHandler = http.StripPrefix("/app", fileHandler)
	fileHandler = apiCfg.middlewareMetricsInc(fileHandler)
	mux.Handle("/app/", fileHandler)

	mux.HandleFunc("/healthz", handlerReadiness)

	mux.HandleFunc("/metrics", apiCfg.showMetricsHandlerFunc())

	mux.HandleFunc("/reset", apiCfg.resetMetricsHandlerFunc())

	// start server
	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	log.Printf("serving files from %s on port %s", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
