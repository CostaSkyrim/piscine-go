#!/bin/bash

files=$(find . -type f | wc -l)

folders=$(find . -type d | wc -l)

echo "$((files + folders))"
