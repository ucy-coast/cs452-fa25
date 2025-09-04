package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// roundRobinRand deterministically cycles through 0..n-1
type roundRobinRand struct {
	count int
}

func (r *roundRobinRand) Intn(n int) int {
	val := r.count % n
	r.count++
	return val
}

// normalizeInputContent parses original input file content into map[keyword][]docs (sorted)
func normalizeInputContent(input string) map[string][]string {
	m := make(map[string][]string)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		keyword := strings.TrimSpace(parts[0])
		rest := strings.TrimSpace(parts[1])

		spaceIdx := strings.Index(rest, " ")
		if spaceIdx == -1 {
			continue
		}
		docListStr := rest[spaceIdx+1:]
		docs := strings.Split(docListStr, ",")
		for i := range docs {
			docs[i] = strings.TrimSpace(docs[i])
		}
		sort.Strings(docs)
		m[keyword] = docs
	}
	return m
}

// comparePartitionsWithOriginal merges partitions and compares them to original input content
func comparePartitionsWithOriginal(t *testing.T, outputPattern string, numPartitions int, originalContent string) {
	merged := make(map[string][]string)
	for i := 0; i < numPartitions; i++ {
		fname := fmt.Sprintf(outputPattern, i)
		file, err := os.Open(fname)
		if err != nil {
			t.Fatalf("Failed to open partition file %q: %v", fname, err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue
			}
			keyword := strings.TrimSpace(parts[0])
			rest := strings.TrimSpace(parts[1])

			spaceIdx := strings.Index(rest, " ")
			if spaceIdx == -1 {
				continue
			}
			docListStr := rest[spaceIdx+1:]
			docs := strings.Split(docListStr, ",")
			for i := range docs {
				docs[i] = strings.TrimSpace(docs[i])
			}
			merged[keyword] = append(merged[keyword], docs...)
		}
		file.Close()
	}

	// Sort docs for each keyword
	for k := range merged {
		sort.Strings(merged[k])
	}

	original := normalizeInputContent(originalContent)

	if len(original) != len(merged) {
		t.Fatalf("Keyword count mismatch: original %d, merged %d", len(original), len(merged))
	}

	for kw, origDocs := range original {
		mergedDocs, ok := merged[kw]
		if !ok {
			t.Errorf("Keyword %q missing in merged partitions", kw)
			continue
		}
		if len(origDocs) != len(mergedDocs) {
			t.Errorf("Document count mismatch for keyword %q: original %d, merged %d", kw, len(origDocs), len(mergedDocs))
			continue
		}
		for i := range origDocs {
			if origDocs[i] != mergedDocs[i] {
				t.Errorf("Document mismatch for keyword %q at pos %d: original %q, merged %q", kw, i, origDocs[i], mergedDocs[i])
			}
		}
	}
}

func assertPartitionsNotEmpty(t *testing.T, outputPattern string, numPartitions int) {
	t.Helper()
	for i := 0; i < numPartitions; i++ {
		fname := fmt.Sprintf(outputPattern, i)
		info, err := os.Stat(fname)
		if err != nil {
			t.Fatalf("Failed to stat partition file %q: %v", fname, err)
		}
		if info.Size() == 0 {
			t.Errorf("Partition file %q is empty", fname)
		}
	}
}

func TestPartitionIndexFile(t *testing.T) {
	tmpDir := t.TempDir()

	inputContent := `apple: 3 doc1, doc2, doc3
banana: 2 doc4, doc5
cherry: 1 doc6`

	inputFile := filepath.Join(tmpDir, "index.txt")
	if err := os.WriteFile(inputFile, []byte(inputContent), 0644); err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	outputPattern := filepath.Join(tmpDir, "partition-%d.txt")
	numPartitions := 3

	rr := &roundRobinRand{}

	err := PartitionIndexFile(inputFile, numPartitions, outputPattern, rr)
	if err != nil {
		t.Fatalf("PartitionIndexFile failed: %v", err)
	}

	assertPartitionsNotEmpty(t, outputPattern, numPartitions)
	comparePartitionsWithOriginal(t, outputPattern, numPartitions, inputContent)
}
