package testutil

import (
	"os"
	"testing"
)

// helper to write lines to a file
func WriteLines(t *testing.T, filename string, lines []string) {
	t.Helper()
	f, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create file %s: %v", filename, err)
	}
	defer f.Close()

	for _, line := range lines {
		if _, err := f.WriteString(line + "\n"); err != nil {
			t.Fatalf("failed to write to file %s: %v", filename, err)
		}
	}
}
