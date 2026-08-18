package infoexchange

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func TestSetValuesHandlerValidation(t *testing.T) {
	db, server := newTestRedis(t)
	handler := SetValuesHandler(*db)

	tests := []struct {
		name  string
		query url.Values
	}{
		{name: "missing key", query: url.Values{"value": {"value"}}},
		{name: "empty key", query: url.Values{"key": {""}, "value": {"value"}}},
		{name: "missing value", query: url.Values{"key": {"key"}}},
		{name: "empty value", query: url.Values{"key": {"key"}, "value": {""}}},
		{name: "both missing", query: url.Values{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := performRequest(handler, http.MethodPost, tt.query)
			if recorder.Code != http.StatusNotAcceptable {
				t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusNotAcceptable, recorder.Body.String())
			}
			if len(server.Keys()) != 0 {
				t.Fatalf("validation failure wrote Redis keys: %v", server.Keys())
			}
		})
	}
}

func TestSetValuesHandlerStoresValues(t *testing.T) {
	db, server := newTestRedis(t)
	handler := SetValuesHandler(*db)

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{name: "plain text", key: "greeting", value: "hello world"},
		{name: "function source", key: "handler", value: "func add(a, b int) int { return a + b }"},
		{name: "URL special characters", key: "path/a+b&c", value: "?first=one&second=two#fragment"},
		{name: "Unicode", key: "saudacao", value: "olá 世界"},
		{name: "JSON text", key: "document", value: `{"enabled":true,"count":3}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := performRequest(handler, http.MethodPost, url.Values{"key": {tt.key}, "value": {tt.value}})
			if recorder.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusOK, recorder.Body.String())
			}
			if recorder.Body.Len() != 0 {
				t.Errorf("body = %q, want empty", recorder.Body.String())
			}

			got, err := server.Get(tt.key)
			if err != nil {
				t.Fatalf("get stored value: %v", err)
			}
			if got != tt.value {
				t.Errorf("stored value = %q, want %q", got, tt.value)
			}
			if ttl := server.TTL(tt.key); ttl != 0 {
				t.Errorf("TTL = %v, want no expiration", ttl)
			}
		})
	}
}

func TestSetValuesHandlerOverwritesExistingValue(t *testing.T) {
	db, server := newTestRedis(t)
	handler := SetValuesHandler(*db)
	server.Set("key", "old")

	recorder := performRequest(handler, http.MethodPost, url.Values{"key": {"key"}, "value": {"new"}})

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	got, err := server.Get("key")
	if err != nil {
		t.Fatalf("get stored value: %v", err)
	}
	if got != "new" {
		t.Errorf("stored value = %q, want new", got)
	}
}

func TestSetValuesHandlerRedisFailure(t *testing.T) {
	db := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:1",
		DialTimeout:  20 * time.Millisecond,
		ReadTimeout:  20 * time.Millisecond,
		WriteTimeout: 20 * time.Millisecond,
		MaxRetries:   -1,
	})
	t.Cleanup(func() { _ = db.Close() })

	recorder := performRequest(SetValuesHandler(*db), http.MethodPost, url.Values{"key": {"key"}, "value": {"value"}})

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusInternalServerError, recorder.Body.String())
	}
}

func TestGetValuesHandlerValidation(t *testing.T) {
	db, _ := newTestRedis(t)
	handler := GetValuesHandler(*db)

	for _, query := range []url.Values{{}, {"key": {""}}} {
		recorder := performRequest(handler, http.MethodGet, query)
		if recorder.Code != http.StatusNotAcceptable {
			t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusNotAcceptable, recorder.Body.String())
		}
	}
}

func TestGetValuesHandlerReturnsStoredValues(t *testing.T) {
	db, server := newTestRedis(t)
	handler := GetValuesHandler(*db)

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{name: "plain text", key: "plain", value: "hello world"},
		{name: "function source", key: "function", value: "func() string { return `worked` }"},
		{name: "escaped text", key: "escaped", value: "line one\n\"line two\"\\end"},
		{name: "Unicode", key: "unicode", value: "ação 日本語"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.Set(tt.key, tt.value)

			recorder := performRequest(handler, http.MethodGet, url.Values{"key": {tt.key}})
			if recorder.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusOK, recorder.Body.String())
			}
			if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json; charset=utf-8" {
				t.Errorf("Content-Type = %q, want application/json; charset=utf-8", contentType)
			}

			var got string
			if err := json.Unmarshal(recorder.Body.Bytes(), &got); err != nil {
				t.Fatalf("decode response: %v", err)
			}
			if got != tt.value {
				t.Errorf("response = %q, want %q", got, tt.value)
			}
		})
	}
}

func TestGetValuesHandlerMissingKey(t *testing.T) {
	db, _ := newTestRedis(t)

	recorder := performRequest(GetValuesHandler(*db), http.MethodGet, url.Values{"key": {"missing"}})

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	var got string
	if err := json.Unmarshal(recorder.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got != "" {
		t.Errorf("response = %q, want empty string", got)
	}
}

func TestSetThenGetRoundTrip(t *testing.T) {
	db, _ := newTestRedis(t)
	key := "calculator/function"
	value := "func multiply(left, right int) int { return left * right }"

	setRecorder := performRequest(SetValuesHandler(*db), http.MethodPost, url.Values{"key": {key}, "value": {value}})
	if setRecorder.Code != http.StatusOK {
		t.Fatalf("set status = %d, want %d; body = %q", setRecorder.Code, http.StatusOK, setRecorder.Body.String())
	}

	getRecorder := performRequest(GetValuesHandler(*db), http.MethodGet, url.Values{"key": {key}})
	if getRecorder.Code != http.StatusOK {
		t.Fatalf("get status = %d, want %d; body = %q", getRecorder.Code, http.StatusOK, getRecorder.Body.String())
	}
	var got string
	if err := json.Unmarshal(getRecorder.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got != value {
		t.Errorf("round-trip value = %q, want %q", got, value)
	}
}

func newTestRedis(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	server := miniredis.RunT(t)
	db := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { _ = db.Close() })
	if err := db.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("ping test Redis: %v", err)
	}
	return db, server
}

func performRequest(handler gin.HandlerFunc, method string, query url.Values) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	request := httptest.NewRequest(method, "/?"+query.Encode(), nil)
	ctx.Request = request
	handler(ctx)
	return recorder
}
