#!/bin/bash

# URL and headers
URL='https://test-service-grid-b.jidutest.com/grpc_c_sleep'
HEADERS=(
    '-H' 'Content-Type: application/json'
)

# Data payload
DATA='{
    "name": "test",
    "time": 1,
    "uuid": "x"
}'

# Perform 15 iterations
for ((i=1; i<=15; i++))
do
    echo "Request $i:"
    curl -s -X POST "${HEADERS[@]}" \
         "$URL" \
         --data "$DATA"

    echo -e "\n" # Add a newline for readability
    sleep 1
done