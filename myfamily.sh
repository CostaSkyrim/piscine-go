#!/bin/bash



curl -s "https://platform.zone01.gr/assets/superhero/all.json" | jq -r --argjson HERO_ID "$HERO_ID" '.[] | select(.id == ($HERO_ID | tonumber)) | .connections.relatives'

