#!/bin/bash

# Loop through namespaces ending with '-sl' and also the 'develop' namespace
for namespace in $(kubectl --kubeconfig ~/.kube/config-dev-mesh get namespaces --no-headers | awk '/-sl$|tls-test/ {print $1}'); do
    # Get all services in the namespace
    for svc in $(kubectl --kubeconfig ~/.kube/config-dev-mesh get svc -n "$namespace" --no-headers | awk '{print $1}'); do
        # Annotate the service with its own name as the value for the specific annotation
        kubectl --kubeconfig ~/.kube/config-dev-mesh annotate svc "$svc" -n "$namespace" alpha.istio.io/kubernetes-serviceaccounts="$svc" --overwrite
    done
done