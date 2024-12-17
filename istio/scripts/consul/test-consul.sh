#!/bin/bash

CONSUL_ADDRESS="10.80.50.44:8500"

register_service() {
  for i in $(seq -w 0001 1000); do
    SERVICE_NAME="consultest$i"
    SERVICE_ID="consultest$i"
    echo "Registering service $SERVICE_NAME with ID $SERVICE_ID"
    curl -X PUT -H "Content-Type: application/json" \
      http://$CONSUL_ADDRESS/v1/agent/service/register \
      -d '{
        "ID": "'$SERVICE_ID'",
        "Name": "'$SERVICE_NAME'",
        "Address": "127.0.0.1",
        "Port": 8080
      }'
  done
}

deregister_service() {
  for i in $(seq -w 0001 1000); do
    SERVICE_ID="consultest$i"
    echo "Deregistering service with ID $SERVICE_ID"
    curl -X PUT http://$CONSUL_ADDRESS/v1/agent/service/deregister/$SERVICE_ID
  done
}

if [ "$1" == "start" ]; then
  register_service
elif [ "$1" == "stop" ]; then
  deregister_service
else
  echo "Usage: $0 <start|stop>"
  exit 1
fi
