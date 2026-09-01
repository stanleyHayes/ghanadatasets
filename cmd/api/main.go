package main

import (
	"encoding/json"
	"github.com/stanleyHayes/ghanadatasets/internal/catalog"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type server struct {
	dataset       *catalog.Dataset
	allowedOrigin string
}

func main() {
	dataPath := os.Getenv("DATASETS_DATA_PATH")
	if dataPath == "" {
		dataPath = "data/datasets.json"
	}
	raw, e := os.ReadFile(dataPath)
	if e != nil {
		log.Fatal(e)
	}
	d, e := catalog.Parse(raw)
	if e != nil {
		log.Fatal(e)
	}
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		origin = "https://datasets.digitalghana.dev"
	}
	s := &server{d, origin}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /v1/datasets", s.datasets)
	mux.HandleFunc("GET /v1/datasets/{id}", s.datasetOne)
	mux.HandleFunc("POST /graphql", s.graphql)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("GhanaDataset API :%s dataset=%s", port, d.DatasetVersion)
	log.Fatal(http.ListenAndServe(":"+port, s.middleware(mux)))
}
func (s *server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Cache-Control", "public, max-age=60")
		if origin := r.Header.Get("Origin"); origin != "" && origin == s.allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(204)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	review := 0
	for _, r := range s.dataset.Datasets {
		if r.Status != "REACHABLE" {
			review++
		}
	}
	writeJSON(w, 200, map[string]any{"status": "ok", "datasetVersion": s.dataset.DatasetVersion, "coverage": s.dataset.Coverage, "datasets": len(s.dataset.Datasets), "review": review, "checkedAt": time.Now().UTC().Format(time.RFC3339)})
}
func (s *server) datasets(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	result := s.dataset.Search(catalog.Query{Search: q.Get("q"), Publisher: q.Get("publisher"), Topic: q.Get("topic"), Format: q.Get("format")})
	writeJSON(w, 200, map[string]any{"data": result, "count": len(result), "datasetVersion": s.dataset.DatasetVersion, "coverage": s.dataset.Coverage})
}
func (s *server) datasetOne(w http.ResponseWriter, r *http.Request) {
	item, ok := s.dataset.Get(r.PathValue("id"))
	if !ok {
		writeJSON(w, 404, map[string]string{"code": "DATASET_NOT_FOUND", "message": "No dataset metadata record exists for this stable ID."})
		return
	}
	writeJSON(w, 200, map[string]any{"data": item, "datasetVersion": s.dataset.DatasetVersion})
}
func (s *server) graphql(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables"`
	}
	if e := json.NewDecoder(http.MaxBytesReader(w, r.Body, 32<<10)).Decode(&body); e != nil {
		writeJSON(w, 400, map[string]any{"errors": []map[string]string{{"message": "Invalid JSON request"}}})
		return
	}
	if strings.Contains(body.Query, "datasets") {
		q, _ := body.Variables["q"].(string)
		topic, _ := body.Variables["topic"].(string)
		result := s.dataset.Search(catalog.Query{Search: q, Topic: topic})
		writeJSON(w, 200, map[string]any{"data": map[string]any{"datasets": result}, "extensions": map[string]string{"datasetVersion": s.dataset.DatasetVersion}})
		return
	}
	writeJSON(w, 400, map[string]any{"errors": []map[string]string{{"message": "Only the documented datasets query is supported in beta."}}})
}
