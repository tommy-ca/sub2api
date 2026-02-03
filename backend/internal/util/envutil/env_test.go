package envutil

import (
	"os"
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
		{"1", false, true},
		{"yes", false, true},
		{"false", true, false},
		{"0", true, false},
		{"no", true, false},
		{"invalid", true, true},
		{"", false, false},
	}

	for _, tt := range tests {
		if tt.val != "" {
			os.Setenv("TEST_BOOL", tt.val)
		} else {
			os.Unsetenv("TEST_BOOL")
		}

		if GetBool("TEST_BOOL", tt.def) != tt.expected {
			t.Errorf("GetBool(%s, %t) failed, expected %t", tt.val, tt.def, tt.expected)
		}
	}
}
