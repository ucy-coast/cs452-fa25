
package indexserver

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/ucy-coast/websearch/internal/rpc_api"
	"github.com/ucy-coast/websearch/internal/invertedindex"
)

// helper to write lines to a file
func writeLines(t *testing.T, filename string, lines []string) {
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

func TestSearch_MatchesPerDocument(t *testing.T) {
	tmpDir := t.TempDir()

	// Prepare index files with sample data
	file1 := filepath.Join(tmpDir, "index1.txt")

	data1 := []string{
		"apple: 2 doc1,doc2",
		"banana: 1 doc1",
	}

	writeLines(t, file1, data1)

	// Create index, which preloads files into memory
	idx, err := invertedindex.NewInvertedIndex([]string{file1})
	if err != nil {
		t.Fatalf("failed to create index: %v", err)
	}

	s := &IndexServer{index: idx}
	args := &rpc_api.SearchArgs{
		Query: "apple banana cherry",
		TopK:  10,
	}
	var reply []rpc_api.SearchResult
	err = s.Search(args, &reply)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	// Convert to map for easy assertions
	results := make(map[string][]string)
	for _, r := range reply {
		sort.Strings(r.Matches)
		results[r.Document] = r.Matches
	}

	expected := map[string][]string{
		"doc1": {"apple", "banana"},
		"doc2": {"apple"},
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected results: %v, got: %v", expected, results)
	}
}
