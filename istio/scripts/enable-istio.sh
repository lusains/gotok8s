#!/bin/bash

NAMESPACE="mno"
DEPLOYMENTS=$(kubectl get deployments -n $NAMESPACE -o=jsonpath='{.items[*].metadata.name}')
COLOR="mno"

# 为 Deployment 的 Pod 模板添加标签和注解
for DEPLOY in "${DEPLOYMENTS[@]}"; do
    # 更新 Pod 模板以添加 version 和 sidecar.istio.io/inject 标签
    kubectl patch deployment $DEPLOY --namespace=$NAMESPACE \
        --patch "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"sidecar.istio.io/inject\":\"true\"}}}}}"

    # 更新 Pod 模板以添加不劫持端口 annotations
    #kubectl patch deployment $DEPLOY --namespace=$NAMESPACE \
    #    --patch "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"traffic.sidecar.istio.io/excludeOutboundPorts\":\"3306,6379,9876,9200,9092,1883,8500,4317,9411\"}}}}}"

    # 滚动更新重启应用
    #kubectl rollout restart deployment $DEPLOY --namespace=$NAMESPACE
done

kubectl patch deployment mesh-python-demo -n istio-sl --patch "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"instrumentation.opentelemetry.io/inject-python\":\"gotok8s-dev-instrumentation-beta\"}}}}}"kubectl patch deployment mesh-python-demo -n istio-sl --patch "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"instrumentation.opentelemetry.io/inject-python\":\"gotok8s-dev-instrumentation-beta\"}}}}}"

echo "操作完成"
