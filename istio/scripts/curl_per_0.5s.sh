#!/bin/bash

# URL for the request
URL='https://10.80.233.97:8080/ip'

# Headers for the request
HEADERS=(
  -H "Content-Type: application/json"
  -H "Cookie: gotok8s_device_id=bfc83076-3bcf-4463-9cf5-a4aa612a5f6b"
)

# Data payload for the request
DATA='{
  "name": "test",
  "time": 1,
  "uuid": "x"
}'

# Loop to execute the request every 0.5 seconds
while true; do
  # Send the request and capture the HTTP status code
  status_code=$(curl --location -s -o /dev/null -w "%{http_code}" "${URL}" "${HEADERS[@]}" --data "${DATA}")

  # Print the HTTP status code
  echo "HTTP Status Code: $status_code"

  # Wait for 0.5 seconds before the next request
done
