
#!/bin/bash

set -e

# Resolve the directory this script is in
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Assume project root is two levels up from script location
PROJECT_HOME="$(cd "$SCRIPT_DIR/../.." && pwd)"

# Paths to binaries and data
INDEXSERVER_BIN="$PROJECT_HOME/bin/indexserver"
WEBSERVER_BIN="$PROJECT_HOME/bin/webserver"
INDEX_FILE="$PROJECT_HOME/test/data/invertedindex-medium.txt"
HTML_PATH="$PROJECT_HOME/web/static/index.html"

# Start indexserver shard in background
echo "Starting indexserver shard on 127.0.0.1:9090..."
"$INDEXSERVER_BIN" -rpc_addr="127.0.0.1:9090" -index_files="$INDEX_FILE" &
INDEX_PID=$!

# Wait a moment to ensure indexserver is up
sleep 2

# Start webserver in background
echo "Starting webserver on 0.0.0.0:8080..."
"$WEBSERVER_BIN" -addr="0.0.0.0:8080" -shards="127.0.0.1:9090" -htmlPath="$HTML_PATH" -topk=10 &
WEB_PID=$!

# Wait a moment to ensure webserver is up
sleep 2

# Test query
echo "Testing search API..."
curl -s "http://localhost:8080/api/search?q=adventure"

# Cleanup: kill background servers
echo "Cleaning up..."
kill $INDEX_PID $WEB_PID
wait $INDEX_PID $WEB_PID 2>/dev/null

echo "Done."
