#!/bin/bash

NAMESPACE=$1  # Get namespace from command line argument

if [ -z "$NAMESPACE" ]; then
    echo "Usage: $0 <namespace>"
    exit 1
fi

OUTPUT_FILE="${NAMESPACE}-port-list.csv"

echo -e "DEPLOYMENT,PROTOCOL,PORT" > $OUTPUT_FILE

kubectl --kubeconfig ~/.kube/dev-config -n $NAMESPACE get svc -o json |
jq -r '.items[] | .spec.ports[] | (.name + "-port:" + (.port|tostring))' |
sort | uniq >> $OUTPUT_FILE
