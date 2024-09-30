#!/bin/bash

pids=()

for dir in */; do
    if [ -d "$dir" ]; then
        echo "Entering directory: $dir"
        cd "$dir"
        if [ -f "docker-compose.yml" ]; then
            echo "Starting Docker Compose down in $dir"
            docker compose down &
            pids+=($!)
        else
            echo "No docker-compose.yml found in $dir"
        fi
        cd ..
    fi
done

for pid in "${pids[@]}"; do
    wait $pid
done

echo "All Docker Compose down commands completed"
