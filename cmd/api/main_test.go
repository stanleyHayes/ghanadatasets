package main

import (
	"github.com/stanleyHayes/ghanadatasets/internal/catalog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func testServer(t *testing.T) *server {
	t.Helper()
	raw, e := os.ReadFile("../../data/datasets.json")
	if e != nil {
		t.Fatal(e)
	}
	d, e := catalog.Parse(raw)
	if e != nil {
		t.Fatal(e)
	}
	return &server{d, "https://datasets.digitalghana.dev"}
}
func TestRESTAndCORS(t *testing.T) {
	s := testServer(t)
	req := httptest.NewRequest(http.MethodGet, "/v1/datasets?q=GLSS7", nil)
	req.Header.Set("Origin", s.allowedOrigin)
	w := httptest.NewRecorder()
	s.middleware(http.HandlerFunc(s.datasets)).ServeHTTP(w, req)
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"id":"gss-glss7"`) {
		t.Fatalf("%d %s", w.Code, w.Body.String())
	}
	if w.Header().Get("Access-Control-Allow-Origin") != s.allowedOrigin {
		t.Fatal("cors")
	}
}
func TestGraphQLParity(t *testing.T) {
	s := testServer(t)
	req := httptest.NewRequest(http.MethodPost, "/graphql", strings.NewReader(`{"query":"query($q:String){datasets(q:$q){id}}","variables":{"q":"GLSS7"}}`))
	w := httptest.NewRecorder()
	s.graphql(w, req)
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"id":"gss-glss7"`) {
		t.Fatalf("%d %s", w.Code, w.Body.String())
	}
}
