package invertedindex

import (
	"path/filepath"
	"testing"

	"github.com/ucy-coast/websearch/internal/testutil"
)

// compareResults compares two slices of Result ignoring the order of results
// and the order of keywords in the Matches slice. It fails the test if there
// is any mismatch.
func compareResults(t *testing.T, testName string, got, want []Result) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: expected %d results, got %d", testName, len(want), len(got))
		return
	}

	// Create a map from document to matches slice for got results
	gotMap := make(map[string][]string)
	for _, r := range got {
		gotMap[r.Document] = r.Matches
	}

	// Helper to check if two string slices are equal as sets (ignoring order)
	matchesEqual := func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		counts := make(map[string]int)
		for _, s := range a {
			counts[s]++
		}
		for _, s := range b {
			if counts[s] == 0 {
				return false
			}
			counts[s]--
		}
		return true
	}

	for _, w := range want {
		gotMatches, ok := gotMap[w.Document]
		if !ok {
			t.Errorf("%s: expected document %q missing in results", testName, w.Document)
		} else if !matchesEqual(gotMatches, w.Matches) {
			t.Errorf("%s: expected matches for %q: %v, got %v", testName, w.Document, w.Matches, gotMatches)
		}
	}
}

func TestSearchAllMethodsSingleFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Prepare index files with sample data
	file1 := filepath.Join(tmpDir, "index1.txt")

	data1 := []string{
		"apple: 3 doc1,doc2,doc3",
		"banana: 2 doc1,doc3",
	}

	testutil.WriteLines(t, file1, data1)

	// Create index, which preloads files into memory
	idx, err := NewInvertedIndex([]string{file1})
	if err != nil {
		t.Fatalf("failed to create index: %v", err)
	}

	keywords := []string{"apple", "banana"}

	expectedResults := []Result{
		{Document: "doc1", Matches: []string{"apple", "banana"}},
		{Document: "doc3", Matches: []string{"apple", "banana"}},
		{Document: "doc2", Matches: []string{"apple"}},
	}

	// Test SearchScanFiles (scans index files on every call)
	resultsScanFiles, err := idx.SearchScanFiles(keywords)
	if err != nil {
		t.Fatalf("SearchScanFiles failed: %v", err)
	}
	compareResults(t, "SearchScanFiles", resultsScanFiles, expectedResults)

	// Test SearchInMemory (searches in-memory index)
	resultsInMemory, err := idx.SearchInMemory(keywords)
	if err != nil {
		t.Fatalf("SearchInMemory failed: %v", err)
	}
	compareResults(t, "SearchInMemory", resultsInMemory, expectedResults)

	// Test SearchInMemoryTopK (search top 2 results in-memory)
	expectedTop2 := expectedResults[:2]
	resultsTopK, err := idx.SearchInMemoryTopK(keywords, 2)
	if err != nil {
		t.Fatalf("SearchInMemoryTopK failed: %v", err)
	}
	compareResults(t, "SearchInMemoryTopK", resultsTopK, expectedTop2)
}

func TestSearchAllMethodsMultipleFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Prepare index files with sample data
	file1 := filepath.Join(tmpDir, "index1.txt")
	file2 := filepath.Join(tmpDir, "index2.txt")

	data1 := []string{
		"apple: 3 doc1,doc2,doc3",
		"banana: 2 doc1,doc3",
	}
	data2 := []string{
		"cherry: 1 doc2",
		"banana: 1 doc2",
	}

	testutil.WriteLines(t, file1, data1)
	testutil.WriteLines(t, file2, data2)

	// Create index, which preloads files into memory
	idx, err := NewInvertedIndex([]string{file1, file2})
	if err != nil {
		t.Fatalf("failed to create index: %v", err)
	}

	keywords := []string{"apple", "banana"}

	expectedResults := []Result{
		{Document: "doc1", Matches: []string{"apple", "banana"}},
		{Document: "doc2", Matches: []string{"apple", "banana"}},
		{Document: "doc3", Matches: []string{"apple", "banana"}},
	}

	// Test SearchScanFiles (scans index files on every call)
	resultsScanFiles, err := idx.SearchScanFiles(keywords)
	if err != nil {
		t.Fatalf("SearchScanFiles failed: %v", err)
	}
	compareResults(t, "SearchScanFiles", resultsScanFiles, expectedResults)

	// Test SearchInMemory (searches in-memory index)
	resultsInMemory, err := idx.SearchInMemory(keywords)
	if err != nil {
		t.Fatalf("SearchInMemory failed: %v", err)
	}
	compareResults(t, "SearchInMemory", resultsInMemory, expectedResults)

	// Test SearchInMemoryTopK (search top 2 results in-memory)
	expectedTop2 := expectedResults[:2]
	resultsTopK, err := idx.SearchInMemoryTopK(keywords, 2)
	if err != nil {
		t.Fatalf("SearchInMemoryTopK failed: %v", err)
	}
	compareResults(t, "SearchInMemoryTopK", resultsTopK, expectedTop2)
}
