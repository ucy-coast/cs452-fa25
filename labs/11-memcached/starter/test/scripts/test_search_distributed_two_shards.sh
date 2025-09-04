
#!/bin/bash

set -e

# Resolve the directory this script is in
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Assume project root is two levels up from script location
PROJECT_HOME="$(cd "$SCRIPT_DIR/../.." && pwd)"

# Paths to binaries and data
INDEXSERVER_BIN="$PROJECT_HOME/bin/indexserver"
WEBSERVER_BIN="$PROJECT_HOME/bin/webserver"
INDEX_FILE_0="$PROJECT_HOME/test/data/invertedindex-medium-0.txt"
INDEX_FILE_1="$PROJECT_HOME/test/data/invertedindex-medium-1.txt"
HTML_PATH="$PROJECT_HOME/web/static/index.html"

# Start indexserver shards in background
echo "Starting indexserver shard 0 on 127.0.0.1:9090..."
"$INDEXSERVER_BIN" -rpc_addr="127.0.0.1:9090" -index_files="$INDEX_FILE_0" &
PID1=$!

echo "Starting indexserver shard 1 on 127.0.0.1:9091..."
"$INDEXSERVER_BIN" -rpc_addr="127.0.0.1:9091" -index_files="$INDEX_FILE_1" &
PID2=$!

# Wait a moment to ensure index servers are up
sleep 2

# Start webserver with both shards
echo "Starting webserver on 0.0.0.0:8080..."
"$WEBSERVER_BIN" -addr="0.0.0.0:8080" -shards="127.0.0.1:9090,127.0.0.1:9091" -htmlPath="$HTML_PATH" -topk=100 &
PID3=$!

# Wait a moment to ensure webserver is up
sleep 2

# Test query 1
echo "Testing search for 'redemption'..."
curl "http://localhost:8080/api/search?q=redemption"

# Test query 2
echo "Testing search for 'adventure grief'..."
curl -G --data-urlencode "q=adventure grief" "http://localhost:8080/api/search"

# Cleanup
echo "Cleaning up..."
kill $PID1 $PID2 $PID3
wait $PID1 $PID2 $PID3 2>/dev/null

echo "Test complete."
