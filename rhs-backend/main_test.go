package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/ping", nil)

	pingHandler(ctx)

	assertJSONResponse(t, recorder, http.StatusOK, map[string]string{"message": "pong"})
}

func TestHelloHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)

	HelloHandler(ctx)

	assertJSONResponse(t, recorder, http.StatusOK, map[string]string{"message": "Hello World!"})
}

func assertJSONResponse(t *testing.T, recorder *httptest.ResponseRecorder, wantStatus int, wantBody map[string]string) {
	t.Helper()
	if recorder.Code != wantStatus {
		t.Fatalf("status = %d, want %d; body = %q", recorder.Code, wantStatus, recorder.Body.String())
	}
	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json; charset=utf-8" {
		t.Errorf("Content-Type = %q, want application/json; charset=utf-8", contentType)
	}

	var got map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	for key, want := range wantBody {
		if got[key] != want {
			t.Errorf("response[%q] = %q, want %q", key, got[key], want)
		}
	}
}
