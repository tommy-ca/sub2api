package envutil

import (
	"os"
	"strings"
	"testing"
)

func TestGetString(t *testing.T) {
	os.Setenv("TEST_STRING", "hello")
	defer os.Unsetenv("TEST_STRING")

	if GetString("TEST_STRING", "default") != "hello" {
		t.Errorf("GetString failed, expected hello")
	}
	if GetString("NON_EXISTENT", "default") != "default" {
		t.Errorf("GetString failed, expected default")
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("TEST_INT", "123")
	defer os.Unsetenv("TEST_INT")

	if GetInt("TEST_INT", 0) != 123 {
		t.Errorf("GetInt failed, expected 123")
	}
	if GetInt("NON_EXISTENT", 456) != 456 {
		t.Errorf("GetInt failed, expected 456")
	}
	os.Setenv("TEST_INT_INVALID", "abc")
	if GetInt("TEST_INT_INVALID", 789) != 789 {
		t.Errorf("GetInt failed for invalid value, expected default 789")
	}
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		val      string
		def      bool
		expected bool
	}{
		{"true", false, true},
		{" TRUE ", false, true},
		{"1", false, true},
		{" 1 ", false, true},
		{"false", true, false},
		{" FALSE ", true, false},
		{"0", true, false},
		{" 0 ", true, false},
		{"yes", false, false}, // No longer supported (opinionated)
		{"on", false, false},  // No longer supported (opinionated)
		{"no", true, true},    // No longer supported (opinionated)
		{"off", true, true},   // No longer supported (opinionated)
		{"invalid", true, true},
		{"invalid", false, false},
		{"", false, false},
		{"", true, true},
	}

	for _, tt := range tests {
		key := "TEST_BOOL_" + strings.ReplaceAll(strings.TrimSpace(tt.val), " ", "_")
		if tt.val != "" {
			os.Setenv(key, tt.val)
			defer os.Unsetenv(key)
		}

		if res := GetBool(key, tt.def); res != tt.expected {
			t.Errorf("GetBool(%s [raw: %q], default: %t) failed, expected %t, got %t", key, tt.val, tt.def, tt.expected, res)
		}
	}
}

func FuzzGetBool(f *testing.F) {
	f.Add("true")
	f.Add("false")
	f.Add("1")
	f.Add("0")
	f.Add("on")
	f.Add("off")
	f.Add("yes")
	f.Add("no")
	f.Add(" ")
	f.Add("")

	f.Fuzz(func(t *testing.T, val string) {
		key := "FUZZ_TEST_BOOL"
		os.Setenv(key, val)
		defer os.Unsetenv(key)

		// GetBool should never crash
		_ = GetBool(key, true)
		_ = GetBool(key, false)
	})
}
