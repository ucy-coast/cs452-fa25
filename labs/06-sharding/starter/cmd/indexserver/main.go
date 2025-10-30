
package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"path/filepath"
	"strconv"
	"strings"
	"os"

	"github.com/ucy-coast/websearch/internal/indexserver"
	"github.com/ucy-coast/websearch/internal/invertedindex"
)

// parseFlags parses and validates CLI flags; returns plain values.
func parseFlags() (rpcAddr string, indexDir string, indexFiles string, shardID int, numShards int, partitioned bool, useMemory bool) {
	rpcAddrFlag := flag.String("rpc_addr", ":9090", "Address to listen on for RPC")
	indexDirFlag := flag.String("index_dir", "", "Directory containing index-*.txt files")
	indexFilesFlag := flag.String("index_files", "", "Comma-separated list of index files to load")
	shardNameFlag := flag.String("shard_name", "", "Shard hostname in the form 'name-id'; extracts shard ID from the suffix")
	shardIDFlag := flag.Int("shard_id", -1, "Explicit shard ID (0-based). Overrides shard_name if set")
	numShardsFlag := flag.Int("num_shards", -1, "Total number of shards")
	partitionedFlag := flag.Bool("partitioned", false, "Enable shard partitioning")
    useMemoryFlag := flag.Bool("use_memory", false, "Use in-memory index instead of scanning disk files")

	flag.Parse()

	if *indexDirFlag != "" && *indexFilesFlag != "" {
		log.Fatal("Cannot specify both -index_dir and -index_files")
	}
	if *indexDirFlag == "" && *indexFilesFlag == "" {
		log.Fatal("Must specify either -index_dir or -index_files")
	}

	finalShardID := *shardIDFlag

	// If shard_id is not explicitly set, try to extract from shard_name
	if finalShardID < 0 && *shardNameFlag != "" {
		parts := strings.Split(*shardNameFlag, "-")
		if len(parts) < 2 {
			log.Fatalf("Invalid shard_name format: %s. Expected format: name-id", *shardNameFlag)
		}
		idStr := parts[len(parts)-1]
		parsedID, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("Failed to extract shard ID from shard_name '%s': %v", *shardNameFlag, err)
		}
		finalShardID = parsedID
	}

	if *partitionedFlag {
		if finalShardID < 0 || *numShardsFlag <= 0 {
			log.Fatal("Must specify a valid -shard_id or -shard_name and a positive -num_shards when using -partitioned")
		}
	}

	return *rpcAddrFlag, *indexDirFlag, *indexFilesFlag, finalShardID, *numShardsFlag, *partitionedFlag, *useMemoryFlag
}

func listAllFiles(indexDir string) ([]string, error) {
	entries, err := os.ReadDir(indexDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(indexDir, entry.Name()))
		}
	}
	return files, nil
}

func selectFilesForShard(files []string, shardID, numShards int) []string {
	var selected []string
	for _, f := range files {
		base := filepath.Base(f)

		// Find last dash to isolate the index part
		dashIdx := strings.LastIndex(base, "-")
		if dashIdx == -1 || !strings.HasSuffix(base, ".txt") {
			continue // Skip malformed filenames
		}

		// Extract the number after the last dash, before ".txt"
		indexStr := strings.TrimSuffix(base[dashIdx+1:], ".txt")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			continue // Skip if not a valid number
		}
		if index%numShards == shardID {
			selected = append(selected, f)
		}
		
	}
	return selected
}

func main() {
	rpcAddr, indexDir, indexFiles, shardID, numShards, partitioned, useMemory := parseFlags()
	
	var files []string
	var err error

	if indexFiles != "" {
		files = strings.Split(indexFiles, ",")
	} else {
		files, err = listAllFiles(indexDir)
		if err != nil {
			log.Fatalf("Failed to read files from %s: %v", indexDir, err)
		}
		if partitioned {
			files = selectFilesForShard(files, shardID, numShards)
		}
	}

	if len(files) == 0 {
		log.Fatal("No index files found to load")
	}

	index, err := invertedindex.NewInvertedIndex(files)
	if err != nil {
		log.Fatalf("Failed to load index: %v", err)
	}

	server := indexserver.NewIndexServer(index, useMemory)
	err = rpc.Register(server)
	if err != nil {
		log.Fatalf("Failed to register RPC server: %v", err)
	}

	listener, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", rpcAddr, err)
	}

	log.Printf("Index RPC server listening on %s\n", rpcAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
