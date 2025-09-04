
package webserver

import (
	"reflect"
	"testing"

	"github.com/ucy-coast/websearch/internal/rpc_api"

)

func TestAggregateShardResults(t *testing.T) {
	shard1 := []rpc_api.SearchResult{
		{Document: "doc1", Matches: []string{"apple", "banana"}},
		{Document: "doc2", Matches: []string{"apple"}},
	}
	shard2 := []rpc_api.SearchResult{
		{Document: "doc1", Matches: []string{"cherry"}},
		{Document: "doc2", Matches: []string{"banana"}},
	}

	results := aggregateShardResults([][]rpc_api.SearchResult{shard1, shard2}, 10)

	// Convert to map for assertions
	got := make(map[string]int)
	for _, r := range results {
		got[r.Document] = len(r.Matches)
	}

	expected := map[string]int{
		"doc1": 3,
		"doc2": 2,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
