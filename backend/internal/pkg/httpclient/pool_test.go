package httpclient

import (
	"testing"
	"time"
)

func TestGetClient_CacheMetrics(t *testing.T) {
	opts := Options{
		Timeout: 5 * time.Second,
	}

	// First call should be a miss
	hitsBefore, missesBefore := GetCacheStats()
	_, err := GetClient(opts)
	if err != nil {
		t.Fatalf("Failed to get client: %v", err)
	}
	hitsAfter, missesAfter := GetCacheStats()

	if missesAfter != missesBefore+1 {
		t.Errorf("Expected miss count to increment, got %d -> %d", missesBefore, missesAfter)
	}
	if hitsAfter != hitsBefore {
		t.Errorf("Expected hit count to remain same, got %d -> %d", hitsBefore, hitsAfter)
	}

	// Second call with same opts should be a hit
	hitsBefore, missesBefore = hitsAfter, missesAfter
	_, err = GetClient(opts)
	if err != nil {
		t.Fatalf("Failed to get client: %v", err)
	}
	hitsAfter, missesAfter = GetCacheStats()

	if hitsAfter != hitsBefore+1 {
		t.Errorf("Expected hit count to increment, got %d -> %d", hitsBefore, hitsAfter)
	}
	if missesAfter != missesBefore {
		t.Errorf("Expected miss count to remain same, got %d -> %d", missesBefore, missesAfter)
	}
}
