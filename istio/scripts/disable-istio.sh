#!/bin/bash

NAMESPACE="istio-sl"
DEPLOYMENTS=("javakit-demo")
COLOR="istio-sl"

# 为 Deployment 的 Pod 模板添加标签和注解
for DEPLOY in "${DEPLOYMENTS[@]}"; do
    # 更新 Pod 模板以添加 version 和 sidecar.istio.io/inject 标签
    kubectl patch deployment $DEPLOY --namespace=$NAMESPACE \
        --patch "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"version\":null, \"sidecar.istio.io/inject\":null}}}}}"

    # 更新 Pod 模板以添加不劫持端口 annotations
    kubectl patch deployment $DEPLOY --namespace=$NAMESPACE \
        --patch "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"traffic.sidecar.istio.io/excludeOutboundPorts\":null}}}}}"

    # 滚动更新重启应用
    kubectl rollout restart deployment $DEPLOY --namespace=$NAMESPACE
done

echo "操作完成"
