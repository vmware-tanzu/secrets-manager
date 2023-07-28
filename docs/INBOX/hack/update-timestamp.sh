#!/usr/bin/env bash

# Get latest commit hash (first 6 characters)
COMMIT_HASH=$(git rev-parse --short=6 HEAD)

# Get last commit time
COMMIT_TIME=$(git log -1 --format="%cd" --date=short)

DOC=./doks-theme/_includes/site-version.html

echo "$COMMIT_HASH"
echo "$COMMIT_TIME"

# Replace placeholders in your HTML file

# Any other Unix in the world:
## sed -i "s|<span data-type=\"commit-hash\">[^<]*</span>|<span data-type=\"commit-hash\">'$COMMIT_HASH'</span>|g" $DOC
## sed -i "s|<span data-type=\"commit-time\">[^<]*</span>|<span data-type=\"commit-time\">$COMMIT_TIME</span>|g" $DOC

# Mac, BSD sed
sed -i '' "s|<span data-type=\"commit-hash\">[^<]*</span>|<span data-type=\"commit-hash\">$COMMIT_HASH</span>|g" $DOC
sed -i '' "s|<span data-type=\"commit-time\">[^<]*</span>|<span data-type=\"commit-time\">$COMMIT_TIME</span>|g" $DOC
