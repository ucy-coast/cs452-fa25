package invertedindex

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// Result represents a search result with the document name and the list of keyword matches.
type Result struct {
	Document string
	Matches []string 
}

// InvertedIndexEntry stores the number of occurrences and the list of documents for a word.
type InvertedIndexEntry struct {
	Count int
	Documents []string
}

// InvertedIndex manages the in-memory inverted index.
type InvertedIndex struct {
	indexFiles []string
	data map[string]InvertedIndexEntry
}

// NewInvertedIndex creates and returns a new empty InvertedIndex.
func NewInvertedIndex(files []string) (*InvertedIndex, error) {
	return &InvertedIndex{
		indexFiles: files,
		data: nil,
	}, nil
}

// SearchScanFiles scans all index files on each search and returns matched keywords per document,
// sorted by descending number of matches and then alphabetically by document name.
func (idx *InvertedIndex) SearchScanFiles(keywords []string) ([]Result, error) {
	docMatches := make(map[string]map[string]struct{}) 

	for _, filename := range idx.indexFiles {
		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open index file %s: %w", filename, err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			word, entry := parseInvertedIndexEntry(line)
			if word == "" {
				continue
			}
			for _, kw := range keywords {
				if word == kw {
					for _, doc := range entry.Documents {
						if docMatches[doc] == nil {
							docMatches[doc] = make(map[string]struct{})
						}
						docMatches[doc][kw] = struct{}{}
					}
				}
			}
		}
		file.Close()
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading index file %s: %w", filename, err)
		}
	}

	docMatchesSlices := convertSetMapToSliceMap(docMatches)

	return toSortedResults(docMatchesSlices), nil
}

// SearchInMemory searches the preloaded index and returns matched keywords per 
// document, sorted by descending number of matches and then alphabetically by 
// document name.
func (idx *InvertedIndex) SearchInMemory(keywords []string) ([]Result, error) {
	if err := idx.ensureIndexIsLoaded(); err != nil {
		return nil, err
	}
	docMatches := collectDocMatches(idx.data, keywords)
	docMatchesSlices := convertSetMapToSliceMap(docMatches)
	return toSortedResults(docMatchesSlices), nil
}

// searchInMemoryTopK searches the preloaded index and returns the top K results after
// sorting the document frequencies.
func (idx *InvertedIndex) SearchInMemoryTopK(keywords []string, k int) ([]Result, error) {
	// TODO: Implement me
	return nil, errors.New("SearchInMemoryTopK not implemented")
}

func (idx *InvertedIndex) ensureIndexIsLoaded() error {
	if idx.data != nil {
		return nil 
	}

	idx.data = make(map[string]InvertedIndexEntry)
	if err := idx.loadInvertedIndexFiles(idx.indexFiles); err != nil {
		return err
	}
	return nil
}

// parseInvertedIndexEntry parses a line from the inverted index file.
// Each line is expected to be in the format: "word: count doc1,doc2,..."
// It returns the word and the corresponding InvertedIndexEntry.
// If the line is malformed, it logs an error and returns an empty entry.
func parseInvertedIndexEntry(line string) (string, InvertedIndexEntry) {
	// Define the regular expression to capture word, file count, and file list
	re := regexp.MustCompile(`^(\S+): (\d+) (.+)$`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 4 {
		log.Printf("invalid index line: %s", line)
		return "", InvertedIndexEntry{}
	}

	word := matches[1]
	count := 0
	_, err := fmt.Sscanf(matches[2], "%d", &count)
	if err != nil {
		log.Printf("invalid count for word '%s': %v", word, err)
		return "", InvertedIndexEntry{}
	}

	docs := strings.Split(matches[3], ",")
	for i := range docs {
		docs[i] = strings.TrimSpace(docs[i])
	}
	return word, InvertedIndexEntry{
		Count: count,
		Documents: docs,
	}
}

// toSortedResults converts a document-to-keywords map into a sorted slice of Results.
// Sort order: descending number of matched keywords, then alphabetical document name.
// Each Result.Matches slice is sorted alphabetically.
func toSortedResults(docMatches map[string][]string) []Result {
	results := make([]Result, 0, len(docMatches))
	for doc, matches := range docMatches {
		// Clean the document name: remove leading/trailing whitespace, including non-breaking spaces.
		cleanedDoc := strings.TrimFunc(doc, func(r rune) bool {
			return unicode.IsSpace(r) || r == '\u00A0'
		})

		// Sort matches alphabetically
		sortedMatches := make([]string, len(matches))
		copy(sortedMatches, matches)
		sort.Strings(sortedMatches)

		results = append(results, Result{Document: cleanedDoc, Matches: sortedMatches})
	}

	sort.Slice(results, func(i, j int) bool {
		if len(results[i].Matches) != len(results[j].Matches) {
			return len(results[i].Matches) > len(results[j].Matches) // more matches first
		}
		return results[i].Document < results[j].Document // tie-break alphabetically
	})

	return results
}

// mergeIntoIndex parses a file and merges its contents into the given index map.
// The file should contain lines in the format: "word: count doc1,doc2,..."
func mergeIntoIndex(index map[string]InvertedIndexEntry, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open index file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		word, newEntry := parseInvertedIndexEntry(line)
		if word == "" {
			continue
		}

		existingEntry, found := index[word]
		if found {
			// Merge unique documents
			docSet := make(map[string]struct{})
			for _, d := range existingEntry.Documents {
				docSet[d] = struct{}{}
			}
			for _, d := range newEntry.Documents {
				if _, exists := docSet[d]; !exists {
					existingEntry.Documents = append(existingEntry.Documents, d)
					docSet[d] = struct{}{}
				}
			}
			existingEntry.Count = len(docSet)
			index[word] = existingEntry
		} else {
			index[word] = newEntry
		}
	}

	return scanner.Err()
}

// LoadInvertedIndexFiles loads multiple inverted index files into the in-memory index.
func (idx *InvertedIndex) loadInvertedIndexFiles(filenames []string) error {
	for _, filename := range filenames {
		if err := mergeIntoIndex(idx.data, filename); err != nil {
			return err
		}
		log.Printf("Loaded index file: %s\n", filename)
	}
	return nil
}

// collectDocMatches loops through each keyword and gathers the documents that contain it
func collectDocMatches(data map[string]InvertedIndexEntry, keywords []string) map[string]map[string]struct{} {
	docMatches := make(map[string]map[string]struct{})

	for _, kw := range keywords {
		entry, found := data[kw]
		if !found {
			continue
		}
		for _, doc := range entry.Documents {
			if docMatches[doc] == nil {
				docMatches[doc] = make(map[string]struct{})
			}
			docMatches[doc][kw] = struct{}{}
		}
	}
	return docMatches
}

// convertSetMapToSliceMap converts a map of sets (map[string]map[string]struct{})
// into a map of slices (map[string][]string).
func convertSetMapToSliceMap(setMap map[string]map[string]struct{}) map[string][]string {
	sliceMap := make(map[string][]string, len(setMap))
	for key, set := range setMap {
		slice := make([]string, 0, len(set))
		for item := range set {
			slice = append(slice, item)
		}
		sliceMap[key] = slice
	}
	return sliceMap
}
