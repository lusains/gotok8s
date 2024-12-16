#!/bin/bash

# input
NAMESPACE="$1"
PROJECT_NAME="$2"

# add tls info to envoy accesslog
kubectl aaply -f - <<EOF
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tls-log-${PROJECT_NAME}
  namespace: $NAMESPACE
spec:
  workloadSelector:
    labels:
      app: $PROJECT_NAME
  configPatches:
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_OUTBOUND
      listener:
        filterChain:
          filter:
            name: "envoy.filters.network.http_connection_manager"
    patch:
      operation: MERGE
      value:
        typed_config:
          "@type": "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager"
          access_log:
          - name: envoy.file_access_log
            typed_config:
              "@type": "type.googleapis.com/envoy.extensions.access_loggers.stream.v3.StdoutAccessLog"
              path: /dev/stdout
              log_format:
                json_format:
                  downstream_tls_version: "%DOWNSTREAM_TLS_VERSION%"
