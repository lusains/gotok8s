#!/bin/bash

# 切换kubectl context
kubectl ctx test-mesh

# 查询 develop 命名空间下所有已注入 Istio sidecar 的 Deployment
deployments=$(kubectl get deployments -n develop -o json | jq -r '.items[] | select(.spec.template.metadata.labels."sidecar.istio.io/inject" == "true") | .metadata.name')

# 遍历每个 Deployment
for deployment in $deployments; do
  # 获取该 Deployment 的 app-name 标签
  app_name=$(kubectl get deployment "$deployment" -n develop -o json | jq -r '.metadata.labels["app-name"]')

  # 如果没有 app-name 标签，跳过此 Deployment
  if [ -z "$app_name" ]; then
    echo "Skipping deployment $deployment: no app-name label found."
    continue
  fi

  # 创建 ConfigMap 名称，格式为 mesh-app-name
  configmap_name="mesh-$(echo "$app_name" | tr '_' '-')"

  # 生成 ConfigMap
  kubectl create configmap "$configmap_name" -n develop \
    --from-literal=istio_enable=true \
    --dry-run=client -o yaml | kubectl apply -f -

  echo "Created/Updated ConfigMap: $configmap_name"
done

echo "Finished processing deployments in develop namespace."
