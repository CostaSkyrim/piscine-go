#!/bin/bash
 

curl -s  "https://platform.zone01.gr/assets/superhero/all.json" | jq -r '.[] | select(.name | contains("Chameleon"))' | jq -r '.name'
curl -s  "https://platform.zone01.gr/assets/superhero/all.json" | jq -r '.[119] | .powerstats.power'
curl -s  "https://platform.zone01.gr/assets/superhero/all.json" | jq -r '.[119] | .appearance.gender'
