
#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_HOME="$(cd "$SCRIPT_DIR/../.." && pwd)"

WEBSERVER_BIN="$PROJECT_HOME/bin/webserver"
INDEX_FILE="$PROJECT_HOME/test/data/invertedindex-medium.txt"
HTML_PATH="$PROJECT_HOME/web/static/index.html"

echo "Starting server on 0.0.0.0:8080..."
"$WEBSERVER_BIN" -addr="0.0.0.0:8080" -index="$INDEX_FILE" -htmlPath="$HTML_PATH" -topk=10 &
SERVER_PID=$!

sleep 2

echo "Testing search API..."
curl -s "http://localhost:8080/api/search?q=adventure"

echo "Cleaning up..."
kill $SERVER_PID
wait $SERVER_PID 2>/dev/null

echo "Done."
