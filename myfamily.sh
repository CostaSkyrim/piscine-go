#!/bin/bash

# Ensure that HERO_ID is set
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID environment variable is not set."
  exit 1
fi

# Subtract 1 from HERO_ID
PREV_HERO_ID=$((HERO_ID - 1))

# Fetch the JSON and use jq to extract the relatives
curl -s "https://platform.zone01.gr/assets/superhero/all.json" | jq -r --argjson HERO_ID "$PREV_HERO_ID" '.[$HERO_ID] | .connections.relatives'
