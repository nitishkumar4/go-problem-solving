#!/bin/bash

# Exit on error
set -e

# Check input
if [ -z "$1" ]; then
  echo "Usage: $0 path/to/problem-name (e.g. array/two_sum)"
  exit 1
fi

FULL_PATH="$1"
DIR="$FULL_PATH"
BASENAME=$(basename "$FULL_PATH")  # Extracts just the final part like "two_sum"

# Create the directory structure
mkdir -p "$DIR"
mkdir -p "cmd/$DIR"

# Copy templates and rename them with problem name
cp template.go "$DIR/${BASENAME}.go"
cp template_test.go "$DIR/${BASENAME}_test.go"
cp cmd/template_main.go "cmd/$DIR/main.go"

# Replace placeholders inside the files
sed -i "s/__NAME__/${BASENAME}/g" "$DIR/${BASENAME}.go"
sed -i "s/__NAME__/${BASENAME}/g" "$DIR/${BASENAME}_test.go"

echo "Created: $DIR/${BASENAME}.go and ${BASENAME}_test.go and cmd/$DIR/main.go"
