#!/bin/bash

echo "开始检测Kubernetes 1.20升级到1.22的兼容性问题..."

# 检测v1beta1 API 版本的资源
V1BETA1_APIS="CustomResourceDefinition,APIService,TokenReview,SubjectAccessReview,CertificateSigningRequest"
for api in ${V1BETA1_APIS}; do
    if kubectl get ${api}.v1beta1 --all-namespaces 2>/dev/null | grep -q .; then
        echo "警告: 您有使用 v1beta1 API 版本的 ${api}。在 1.22 中，这些将不再受支持。"
    fi
done

# 检测Ingress在v1beta1中的使用
if kubectl get ingress.v1beta1.extensions --all-namespaces 2>/dev/null | grep -q .; then
    echo "警告: 您有使用 v1beta1 API 版本的 Ingress。请迁移到 networking.k8s.io/v1。"
fi

echo "检测完毕。"