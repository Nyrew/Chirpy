package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) middlewareMetrianocsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1) // Increment the counter
		next.ServeHTTP(w, r)      // Call the next handler
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	count := cfg.fileserverHits.Load() // Safely read the counter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	html := fmt.Sprintf(`
		<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
		</html>`, count)
	w.Write([]byte(html))
}
