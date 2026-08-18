package services

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCountServicesReturnsServiceCount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	path := filepath.Join(t.TempDir(), "services.json")
	if err := os.WriteFile(path, []byte(`[{"name":"api"},{"name":"web"}]`), 0o600); err != nil {
		t.Fatalf("write service file: %v", err)
	}
	t.Setenv("SERVICE_PATH", path)

	recorder, ctx := newServicesTestContext()
	CountServices()(ctx)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	if got, want := recorder.Body.String(), "2"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestCountServicesFailures(t *testing.T) {
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
				if err := os.WriteFile(path, []byte(`[{"name":`), 0o600); err != nil {
					t.Fatalf("write service file: %v", err)
				}
				t.Setenv("SERVICE_PATH", path)
			},
		},
		{
			name: "non-array JSON",
			setup: func(t *testing.T) {
				path := filepath.Join(t.TempDir(), "services.json")
				if err := os.WriteFile(path, []byte(`{"name":"api"}`), 0o600); err != nil {
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

			CountServices()(ctx)

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
