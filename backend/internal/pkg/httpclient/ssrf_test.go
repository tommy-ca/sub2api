package httpclient

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewDialer_Control(t *testing.T) {
	dialer := NewDialer(true, false) // validateResolvedIP=true, allowPrivate=false

	blockedIPs := []struct {
		addr string
		name string
	}{
		{"127.0.0.1:8080", "IPv4 Loopback"},
		{"[::1]:8080", "IPv6 Loopback"},
		{"169.254.169.254:80", "Cloud IMDS"},
		{"[fd00:ec2::254]:80", "AWS IMDSv2"},
		{"224.0.0.1:5353", "Multicast"},
		{"0.0.0.0:22", "Unspecified IPv4"},
		{"[::]:22", "Unspecified IPv6"},
		{"0.1.2.3:80", "0.0.0.0/8 range"},
		{"10.0.0.1:443", "RFC 1918 (10.x)"},
		{"172.16.0.1:443", "RFC 1918 (172.x)"},
		{"192.168.1.1:443", "RFC 1918 (192.x)"},
		{"100.64.0.1:80", "CGNAT"},
		{"198.18.0.1:80", "Benchmarking"},
		{"2001:0000:4136:e378:8000:63bf:3fff:fdd2:443", "Teredo"},
		{"[2002:c0a8:0101::1]:443", "6to4"},
	}

	for _, tt := range blockedIPs {
		err := dialer.Control("tcp", tt.addr, nil)
		if err == nil {
			t.Errorf("Expected block for %s (%s), but connection allowed", tt.addr, tt.name)
		} else if !strings.Contains(err.Error(), "SSRF protection") {
			t.Errorf("Expected SSRF protection error for %s, got: %v", tt.addr, err)
		}
	}

	// Allowed IPs
	allowedIPs := []string{
		"8.8.8.8:53",
		"1.1.1.1:443",
		"20.20.20.20:80",
		"93.184.216.34:443", // example.com
	}

	for _, addr := range allowedIPs {
		err := dialer.Control("tcp", addr, nil)
		if err != nil {
			t.Errorf("Expected connection allowed for %s, but got error: %v", addr, err)
		}
	}
}

func TestNewDialer_AllowPrivateHosts(t *testing.T) {
	// Dialer with SSRF protection enabled BUT private hosts allowed
	dialer := NewDialer(true, true)

	// In this case, Control should be nil or not block
	if dialer.Control != nil {
		err := dialer.Control("tcp", "127.0.0.1:80", nil)
		if err != nil {
			t.Errorf("Expected 127.0.0.1 to be allowed when allowPrivateHosts is true, got: %v", err)
		}
	}
}

func TestNewDialer_Integration_Localhost(t *testing.T) {
	// Create a local test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	// Create a hardened client
	client, err := GetClient(Options{
		ValidateResolvedIP: true,
		AllowPrivateHosts:  false,
	})
	if err != nil {
		t.Fatalf("Failed to get client: %v", err)
	}

	// Attempt to request the local server
	_, err = client.Get(ts.URL)
	if err == nil {
		t.Errorf("Expected request to local server %s to be blocked, but it succeeded", ts.URL)
	} else if !strings.Contains(err.Error(), "SSRF protection") {
		t.Errorf("Expected SSRF protection error, got: %v", err)
	}
}

func TestNewDialer_DNSRebinding(t *testing.T) {
	// This test mocks the dialer's resolver to simulate a domain resolving to a private IP.
	// Since Dialer.Control is called AFTER resolution, it should catch the private IP.

	privateIP := net.ParseIP("127.0.0.1")

	dialer := NewDialer(true, false)

	// Directly simulate what net.Dialer would pass to Control after resolving a "malicious" domain
	addr := net.JoinHostPort(privateIP.String(), "80")
	err := dialer.Control("tcp", addr, nil)

	if err == nil {
		t.Errorf("Expected DNS rebinding attempt to %s to be blocked, but it was allowed", addr)
	} else if !strings.Contains(err.Error(), "SSRF protection") {
		t.Errorf("Expected SSRF protection error, got: %v", err)
	}
}

func BenchmarkSSRFControlFunction(b *testing.B) {
	dialer := NewDialer(true, false)
	targetAddr := "93.184.216.34:443" // example.com IP

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := dialer.Control("tcp", targetAddr, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}
