#!/bin/bash


curl -s "https://platform.zone01.gr/assets/superhero/all.json" | jq -r '.[52] | .connections.relatives'
