package util

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// IsValidAddress checks if the given address has a valid host and port.
// If the address is missing a port, it returns an error.
func IsValidAddress(address string) (string, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", fmt.Errorf("invalid address format: %v", err)
	}

	if host == "" && port != "" {
		host = "0.0.0.0" // Default host for an empty host part (i.e., listening on all interfaces)
	}

	portInt, err := strconv.Atoi(port)
	if err != nil || portInt < 0 || portInt > 65535 {
		return "", fmt.Errorf("invalid port: %v", err)
	}

	return fmt.Sprintf("%s:%d", host, portInt), nil
}

// IsValidAddressWithDefaultPort checks if the given address has a valid host and port.
// If the address is missing a port, it adds the default port.
func IsValidAddressWithDefaultPort(addr string, defaultPort int) (string, error) {
	if !strings.Contains(addr, ":") {
		addr = addr + ":" + strconv.Itoa(defaultPort)
	} else if strings.HasSuffix(addr, ":") {
		addr = addr + strconv.Itoa(defaultPort)
	}
	return IsValidAddress(addr)
}
