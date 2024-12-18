#!/bin/bash

# 脚本使用说明
if [ $# -ne 3 ]; then
  echo "Usage: $0 <service-name> <namespace> <lb-port>"
  exit 1
fi

# 获取输入参数
SERVICE_NAME=$1      # 服务名称，例如 apollo-configservice
NAMESPACE=$2         # 命名空间，例如 default
LB_PORT=$3           # LoadBalancer 端口，例如 8080

# 集群A和集群B的kubectl context名称
CTX_A="test"
CTX_B="test-mesh"

# 切换到集群A，获取LoadBalancer的IP
echo "Switching to cluster test: $CTX_A"
kubectl ctx $CTX_A

LB_IP=$(kubectl get svc $SERVICE_NAME -n $NAMESPACE -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

if [ -z "$LB_IP" ]; then
  echo "Failed to retrieve LoadBalancer IP for service $SERVICE_NAME in cluster test."
  exit 1
fi

echo "LoadBalancer IP for service $SERVICE_NAME is $LB_IP"

# 切换到集群B，创建假的Service和Endpoints
echo "Switching to cluster test-mesh: $CTX_B"
kubectl ctx $CTX_B

# 创建假的Service
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: $SERVICE_NAME
  namespace: $NAMESPACE
spec:
  ports:
  - port: $LB_PORT
    targetPort: $LB_PORT
    protocol: TCP
  clusterIP: None
EOF

# 创建Endpoints，指向集群A的LB IP
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Endpoints
metadata:
  name: $SERVICE_NAME
  namespace: $NAMESPACE
subsets:
  - addresses:
      - ip: $LB_IP
    ports:
      - port: $LB_PORT
EOF

echo "Service and Endpoints created in cluster test-mesh pointing to LoadBalancer IP: $LB_IP"
