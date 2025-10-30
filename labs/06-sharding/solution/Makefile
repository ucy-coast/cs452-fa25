.PHONY: all clean test docker docker-indexserver docker-webserver docker-push

# === Build binaries ===
all: indexserver webserver partition_index

webserver:
	go build -o bin/webserver ./cmd/webserver
indexserver:
	go build -o bin/indexserver ./cmd/indexserver

partition_index:
	go build -o bin/partition_index ./tools/partition_index

# === Test ===

test:
	go test ./...

# === Clean ===

clean:
	rm -f bin/indexserver bin/webserver bin/partition_index
