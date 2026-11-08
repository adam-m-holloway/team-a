package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		target     string
		wantStatus int
		wantBody   map[string]string
	}{
		{
			name:       "healthz ok",
			method:     http.MethodGet,
			target:     "/healthz",
			wantStatus: http.StatusOK,
			wantBody:   map[string]string{"status": "ok"},
		},
		{
			name:       "greet with name",
			method:     http.MethodGet,
			target:     "/greet?name=Ada",
			wantStatus: http.StatusOK,
			wantBody:   map[string]string{"message": "Hello, Ada"},
		},
		{
			name:       "greet with url-encoded name",
			method:     http.MethodGet,
			target:     "/greet?name=Ada%20Lovelace",
			wantStatus: http.StatusOK,
			wantBody:   map[string]string{"message": "Hello, Ada Lovelace"},
		},
		{
			name:       "greet missing name",
			method:     http.MethodGet,
			target:     "/greet",
			wantStatus: http.StatusBadRequest,
			wantBody:   map[string]string{"error": "name is required"},
		},
		{
			name:       "greet empty name",
			method:     http.MethodGet,
			target:     "/greet?name=",
			wantStatus: http.StatusBadRequest,
			wantBody:   map[string]string{"error": "name is required"},
		},
		{
			name:       "greet whitespace name",
			method:     http.MethodGet,
			target:     "/greet?name=%20%20",
			wantStatus: http.StatusBadRequest,
			wantBody:   map[string]string{"error": "name is required"},
		},
		{
			name:       "wrong method on healthz",
			method:     http.MethodPost,
			target:     "/healthz",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "unknown path",
			method:     http.MethodGet,
			target:     "/nope",
			wantStatus: http.StatusNotFound,
		},
	}

	h := NewHandler()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.target, nil)
			rec := httptest.NewRecorder()

			h.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d (body %q)", rec.Code, tt.wantStatus, rec.Body.String())
			}
			if tt.wantBody == nil {
				return
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("Content-Type = %q, want application/json", ct)
			}
			var got map[string]string
			if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
				t.Fatalf("decode body %q: %v", rec.Body.String(), err)
			}
			if len(got) != len(tt.wantBody) {
				t.Fatalf("body = %v, want %v", got, tt.wantBody)
			}
			for k, v := range tt.wantBody {
				if got[k] != v {
					t.Errorf("body[%q] = %q, want %q", k, got[k], v)
				}
			}
		})
	}
}
