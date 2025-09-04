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
# === Docker Builds ===

docker: docker-indexserver docker-webserver

docker-indexserver:
	docker build -f deploy/docker/indexserver.Dockerfile -t websearch_indexserver:latest .

docker-webserver:
	docker build -f deploy/docker/webserver.Dockerfile -t websearch_webserver:latest .

# === Docker Tag + Push ===

docker-push: docker-push-indexserver docker-push-webserver

docker-push-indexserver:
	@REGISTRY=$${REGISTRY:-}; \
	TAG=$${REGISTRY:+$${REGISTRY}/}websearch_indexserver:latest; \
	echo "Pushing $$TAG"; \
	docker tag websearch_indexserver:latest $$TAG; \
	docker push $$TAG

docker-push-webserver:
	@REGISTRY=$${REGISTRY:-}; \
	TAG=$${REGISTRY:+$${REGISTRY}/}websearch_webserver:latest; \
	echo "Pushing $$TAG"; \
	docker tag websearch_webserver:latest $$TAG; \
	docker push $$TAG
