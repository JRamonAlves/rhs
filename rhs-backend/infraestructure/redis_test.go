package infraestructure

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestConfig(t *testing.T) {
	originalMode := gin.Mode()
	t.Cleanup(func() { gin.SetMode(originalMode) })

	tests := []struct {
		name     string
		mode     string
		password string
		wantAddr string
	}{
		{name: "debug uses local Redis", mode: gin.DebugMode, wantAddr: "localhost:6379"},
		{name: "test uses service Redis", mode: gin.TestMode, password: "test-secret", wantAddr: "redis:6379"},
		{name: "release uses service Redis", mode: gin.ReleaseMode, password: "release-secret", wantAddr: "redis:6379"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(tt.mode)
			t.Setenv("REDIS_PASSWORD", tt.password)

			got := RedisConfig()

			if got.Addr != tt.wantAddr {
				t.Errorf("Addr = %q, want %q", got.Addr, tt.wantAddr)
			}
			if got.Password != tt.password {
				t.Errorf("Password = %q, want %q", got.Password, tt.password)
			}
			if got.DB != 0 {
				t.Errorf("DB = %d, want 0", got.DB)
			}
		})
	}
}
