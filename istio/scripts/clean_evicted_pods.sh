#!/bin/bash

# 遍历所有命名空间
kubectl ctx dev
# 获取所有命名空间中状态为 Evicted 的 Pod，并逐一删除
kubectl get pods -A --field-selector=status.phase=Failed -o jsonpath='{range .items[?(@.status.reason=="Evicted")]}{.metadata.namespace} {.metadata.name}{"\n"}{end}' | while read -r namespace pod_name; do
    if [[ -n "$namespace" && -n "$pod_name" ]]; then
        echo "Deleting evicted pod: $pod_name in namespace: $namespace"
        kubectl delete pod "$pod_name" -n "$namespace"
    fi
done

echo "Completed cleanup of evicted pods in all namespaces."

