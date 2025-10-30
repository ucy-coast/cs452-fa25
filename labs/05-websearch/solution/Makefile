.PHONY: all clean test docker docker-indexserver docker-webserver docker-push

# === Build binaries ===
all: webserver

webserver:
	go build -o bin/webserver ./cmd/webserver

partition_index:
	go build -o bin/partition_index ./tools/partition_index

# === Test ===

test:
	go test ./...

# === Clean ===

clean:
	rm -f bin/indexserver bin/webserver bin/partition_index
