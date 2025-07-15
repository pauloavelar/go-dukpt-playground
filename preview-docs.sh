#!/bin/bash

# Local documentation preview server
# Serves the docs folder on http://localhost:8080

set -e

DOCS_DIR="$(dirname "$0")/docs"
PORT="${PORT:-8080}"

if [ ! -d "$DOCS_DIR" ]; then
    echo "Error: docs directory not found at $DOCS_DIR"
    exit 1
fi

echo "Starting documentation preview server..."
echo "📖 Documentation will be available at: http://localhost:$PORT"
echo "📁 Serving from: $DOCS_DIR"
echo ""
echo "Press Ctrl+C to stop the server"
echo ""

# Try Python 3 first, then Python 2, then fallback to other options
if command -v python3 >/dev/null 2>&1; then
    cd "$DOCS_DIR" && python3 -m http.server "$PORT"
elif command -v python >/dev/null 2>&1; then
    cd "$DOCS_DIR" && python -m SimpleHTTPServer "$PORT"
elif command -v ruby >/dev/null 2>&1; then
    cd "$DOCS_DIR" && ruby -run -e httpd . -p "$PORT"
elif command -v php >/dev/null 2>&1; then
    cd "$DOCS_DIR" && php -S "localhost:$PORT"
elif command -v node >/dev/null 2>&1; then
    echo "Installing basic HTTP server via npm..."
    npx http-server "$DOCS_DIR" -p "$PORT" -o
else
    echo "Error: No suitable HTTP server found."
    echo "Please install one of: python3, python, ruby, php, or node.js"
    exit 1
fi