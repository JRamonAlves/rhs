package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetServicesHandlerReturnsJSONFile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	want := []Service{{
		Name:        "api",
		Url:         "https://api.example.com",
		Port:        443,
		Description: "API service",
		Category:    Services,
	}}
	path := filepath.Join(t.TempDir(), "services.json")
	data, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal services: %v", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("write service file: %v", err)
	}
	t.Setenv("SERVICE_PATH", path)

	recorder, ctx := newServicesTestContext()
	GetServicesHandler()(ctx)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json; charset=utf-8" {
		t.Errorf("Content-Type = %q, want application/json; charset=utf-8", contentType)
	}
	if !json.Valid(recorder.Body.Bytes()) {
		t.Fatalf("response is not valid JSON: %q", recorder.Body.String())
	}
	var got []Service
	if err := json.Unmarshal(recorder.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(got) != len(want) || got[0] != want[0] {
		t.Errorf("services = %#v, want %#v", got, want)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		category Category
		want     bool
	}{
		{category: MoviesAndSeries, want: true},
		{category: Services, want: true},
		{category: Photos, want: true},
		{category: Comics, want: true},
		{category: Category("Other Media"), want: false},
		{category: Category(""), want: false},
	}

	for _, tt := range tests {
		t.Run(string(tt.category), func(t *testing.T) {
			if got := IsValid(tt.category); got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.category, got, tt.want)
			}
		})
	}
}

func TestGetServicesHandlerFailures(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name  string
		setup func(t *testing.T)
	}{
		{
			name: "missing service path",
			setup: func(t *testing.T) {
				t.Setenv("SERVICE_PATH", "")
			},
		},
		{
			name: "unreadable service path",
			setup: func(t *testing.T) {
				t.Setenv("SERVICE_PATH", filepath.Join(t.TempDir(), "missing.json"))
			},
		},
		{
			name: "invalid JSON",
			setup: func(t *testing.T) {
				path := filepath.Join(t.TempDir(), "services.json")
				if err := os.WriteFile(path, []byte(`{"services":`), 0o600); err != nil {
					t.Fatalf("write service file: %v", err)
				}
				t.Setenv("SERVICE_PATH", path)
			},
		},
		{
			name: "invalid category",
			setup: func(t *testing.T) {
				path := filepath.Join(t.TempDir(), "services.json")
				if err := os.WriteFile(path, []byte(`[{"name":"api","category":"Other Media"}]`), 0o600); err != nil {
					t.Fatalf("write service file: %v", err)
				}
				t.Setenv("SERVICE_PATH", path)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)
			recorder, ctx := newServicesTestContext()

			GetServicesHandler()(ctx)

			if recorder.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusInternalServerError, recorder.Body.String())
			}
			if !ctx.IsAborted() {
				t.Error("context was not aborted")
			}
			if len(ctx.Errors) != 1 {
				t.Fatalf("errors = %d, want 1", len(ctx.Errors))
			}
		})
	}
}

func newServicesTestContext() (*httptest.ResponseRecorder, *gin.Context) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/getServices", nil)
	return recorder, ctx
}
