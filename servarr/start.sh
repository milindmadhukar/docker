#!/bin/bash

for dir in */; do
    if [ -d "$dir" ]; then
        echo "Entering directory: $dir"
        cd "$dir"
        if [ -f "docker-compose.yml" ]; then
            echo "Starting Docker Compose in $dir"
            docker compose up -d
        else
            echo "No docker-compose.yml found in $dir"
        fi
        cd ..
    fi
done
