#!/bin/bash

# 筛选符合条件的 Pod，并获取它们对应的 Deployment 名称
kdev="kubectl --kubeconfig ~/.kube/config-dev"

kdev get pods -A | grep -v "Running" | grep -v "Evicted" | grep "develop" | grep -v "Completed" | while read -r line; do
    # 获取 Pod 名称和命名空间
    pod_name=$(echo "$line" | awk '{print $2}')
    namespace=$(echo "$line" | awk '{print $1}')

    # 根据 Pod 名称去掉最后两段字符，推测 Deployment 名称
    deployment_name=$(echo "$pod_name" | sed 's/-[a-z0-9]\{9\}-[a-z0-9]\{5\}$//')

    # 检查是否成功获取 Deployment 名称
    if [[ -n "$deployment_name" ]]; then
        echo "$deployment_name"
    else
        echo "No deployment information found for pod $pod_name"
    fi
done
