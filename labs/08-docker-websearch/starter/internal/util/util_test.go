package util

import (
	"testing"
)

func TestIsValidAddress(t *testing.T) {
	tests := []struct {
		address   string
		expected  string
		expectErr bool
	}{
		// Valid address with host and port
		{"127.0.0.1:8080", "127.0.0.1:8080", false},
		// Valid address with host only (defaults to "0.0.0.0")
		{"::8080", "", true},
		// Invalid port format (non-numeric port)
		{"127.0.0.1:abcd", "", true},
		// Invalid port range (negative)
		{"127.0.0.1:-1", "", true},
		// Invalid port range (too high)
		{"127.0.0.1:70000", "", true},
		// Address with missing port
		{"127.0.0.1", "", true},
		// Valid address with port 0
		{"127.0.0.1:0", "127.0.0.1:0", false},
	}

	for _, tt := range tests {
		t.Run(tt.address, func(t *testing.T) {
			result, err := IsValidAddress(tt.address)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestIsValidAddressWithDefaultPort(t *testing.T) {
	tests := []struct {
		address     string
		defaultPort int
		expected    string
		expectErr   bool
	}{
		// Valid address with host and port
		{"127.0.0.1:8080", 3000, "127.0.0.1:8080", false},
		// Valid address with host only, should add the default port
		{"127.0.0.1", 3000, "127.0.0.1:3000", false},
		// Address with missing port but ends with a colon
		{"127.0.0.1:", 3000, "127.0.0.1:3000", false},
		// Address with missing host and port, should add the default port
		{"", 3000, "0.0.0.0:3000", false},
		// Invalid address with wrong port
		{"127.0.0.1:abcd", 3000, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.address, func(t *testing.T) {
			result, err := IsValidAddressWithDefaultPort(tt.address, tt.defaultPort)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
