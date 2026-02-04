package urlvalidator

import (
	"net"
	"testing"
)

func TestValidateURLFormat(t *testing.T) {
	if _, err := ValidateURLFormat("", false); err == nil {
		t.Fatalf("expected empty url to fail")
	}
	if _, err := ValidateURLFormat("://bad", false); err == nil {
		t.Fatalf("expected invalid url to fail")
	}
	if _, err := ValidateURLFormat("http://example.com", false); err == nil {
		t.Fatalf("expected http to fail when allow_insecure_http is false")
	}
	if _, err := ValidateURLFormat("https://example.com", false); err != nil {
		t.Fatalf("expected https to pass, got %v", err)
	}
	if _, err := ValidateURLFormat("http://example.com", true); err != nil {
		t.Fatalf("expected http to pass when allow_insecure_http is true, got %v", err)
	}
	if _, err := ValidateURLFormat("https://example.com:bad", true); err == nil {
		t.Fatalf("expected invalid port to fail")
	}

	// 验证末尾斜杠被移除
	normalized, err := ValidateURLFormat("https://example.com/", false)
	if err != nil {
		t.Fatalf("expected trailing slash url to pass, got %v", err)
	}
	if normalized != "https://example.com" {
		t.Fatalf("expected trailing slash to be removed, got %s", normalized)
	}

	// 验证多个末尾斜杠被移除
	normalized, err = ValidateURLFormat("https://example.com///", false)
	if err != nil {
		t.Fatalf("expected multiple trailing slashes to pass, got %v", err)
	}
	if normalized != "https://example.com" {
		t.Fatalf("expected all trailing slashes to be removed, got %s", normalized)
	}

	// 验证带路径的 URL 末尾斜杠被移除
	normalized, err = ValidateURLFormat("https://example.com/api/v1/", false)
	if err != nil {
		t.Fatalf("expected trailing slash url with path to pass, got %v", err)
	}
	if normalized != "https://example.com/api/v1" {
		t.Fatalf("expected trailing slash to be removed from path, got %s", normalized)
	}
}

func TestIsPrivateIP(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		// Standard Loopback
		{"127.0.0.1", true},
		{"127.0.0.2", true},
		{"::1", true},
		// IPv4-mapped IPv6 Loopback
		{"::ffff:127.0.0.1", true},

		// RFC 1918 Private
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"192.168.1.1", true},
		// IPv4-mapped IPv6 Private
		{"::ffff:10.0.0.1", true},

		// Link-local
		{"169.254.169.254", true},
		{"fe80::1", true},
		// IPv4-mapped IPv6 Link-local
		{"::ffff:169.254.169.254", true},

		// Unspecified / Broadcast / 0.0.0.0/8
		{"0.0.0.0", true},
		{"0.0.0.1", true},
		{"0.255.255.255", true},
		{"1.0.0.0", false}, // Just outside
		{"::", true},
		{"255.255.255.255", true},

		// Carrier-grade NAT (100.64.0.0/10)
		{"100.64.0.1", true},
		{"100.127.255.255", true},
		{"100.63.255.255", false}, // Just outside
		{"100.128.0.0", false},    // Just outside

		// Benchmarking (198.18.0.0/15)
		{"198.18.0.1", true},
		{"198.19.255.255", true},

		// Reserved / Documentation
		{"192.0.2.1", true},    // TEST-NET-1
		{"198.51.100.1", true}, // TEST-NET-2
		{"203.0.113.1", true},  // TEST-NET-3
		{"240.0.0.1", true},    // Reserved
		{"254.0.0.1", true},    // Reserved

		// Multicast
		{"224.0.0.1", true},
		{"239.255.255.255", true},
		{"ff02::1", true},
		{"ff00::1", true},

		// Transition Mechanisms
		{"2001:0000:4136:e378:8000:63bf:3fff:fdd2", true}, // Teredo
		{"2002:c0a8:0101::1", true},                       // 6to4

		// Cloud IMDS
		{"169.254.169.254", true},
		{"fd00:ec2::254", true}, // AWS IMDSv2

		// Public addresses
		{"8.8.8.8", false},
		{"1.1.1.1", false},
		{"20.20.20.20", false},
		{"2606:4700:4700::1111", false},
	}

	for _, tt := range tests {
		if res := IsPrivateIP(net.ParseIP(tt.ip)); res != tt.expected {
			t.Errorf("IsPrivateIP(%s) failed, expected %t, got %t", tt.ip, tt.expected, res)
		}
	}
}
