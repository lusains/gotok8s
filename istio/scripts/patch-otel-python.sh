#!/bin/bash

kubectl -n develop patch deployment javakit-benchmark --patch "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"instrumentation.opentelemetry.io/container-names\":\"javakit-test\"}}}}}"

k -n develop patch deployment javakit-demo --patch "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"instrumentation.opentelemetry.io/container-names\":\"javakit-demo\"}}}}}"