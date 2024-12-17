#!/bin/bash

# Set context names for A and B clusters
A_CLUSTER_CONTEXT="dev-mesh"
B_CLUSTER_CONTEXT="cluster-B"
NAMESPACE="develop"

# Temporary files to store the deployment names
A_DEPLOYMENTS_FILE="/tmp/a_deployments.txt"
B_DEPLOYMENTS_FILE="/tmp/b_deployments.txt"

# Get deployments with Istio sidecar in A cluster
kubectl --context="$A_CLUSTER_CONTEXT" get deployments -n "$NAMESPACE" -o json | \
  jq -r '.items[] | select(.spec.template.metadata.annotations."sidecar.istio.io/status") | .metadata.name' > "$A_DEPLOYMENTS_FILE"

# Get deployments with Istio sidecar in B cluster
kubectl --context="$B_CLUSTER_CONTEXT" get deployments -n "$NAMESPACE" -o json | \
  jq -r '.items[] | select(.spec.template.metadata.annotations."sidecar.istio.io/status") | .metadata.name' > "$B_DEPLOYMENTS_FILE"

# Compare the deployments, output the ones only in B cluster
echo "Deployments in B cluster but not in A cluster:"
comm -23 <(sort "$B_DEPLOYMENTS_FILE") <(sort "$A_DEPLOYMENTS_FILE")

# Clean up the temporary files
rm "$A_DEPLOYMENTS_FILE" "$B_DEPLOYMENTS_FILE"
