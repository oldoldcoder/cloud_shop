#!/bin/bash

# 确保脚本在遇到错误时退出
set -e
cd ..
# 创建服务目录及子目录
echo "创建服务目录结构..."
mkdir -p services/user-service/src
touch services/user-service/pom.xml
touch services/user-service/Dockerfile
touch services/user-service/README.md

mkdir -p services/product-service/cmd services/product-service/internal
touch services/product-service/go.mod
touch services/product-service/Dockerfile
touch services/product-service/README.md

mkdir -p services/order-service/src
touch services/order-service/pom.xml
touch services/order-service/Dockerfile
touch services/order-service/README.md

mkdir -p services/payment-service
touch services/payment-service/app.py
touch services/payment-service/requirements.txt
touch services/payment-service/Dockerfile
touch services/payment-service/README.md

# 创建基础设施目录
echo "创建基础设施目录结构..."
mkdir -p infrastructure/docker
touch infrastructure/docker/docker-compose.yml
touch infrastructure/docker/README.md

mkdir -p infrastructure/k8s
touch infrastructure/k8s/user-service.yaml
touch infrastructure/k8s/product-service.yaml
touch infrastructure/k8s/order-service.yaml
touch infrastructure/k8s/payment-service.yaml

# 创建文档目录
echo "创建文档目录结构..."
mkdir -p docs/api-specs docs/decisions
touch docs/architecture.md
touch docs/api-specs/user-service.yaml
touch docs/api-specs/product-service.yaml
touch docs/decisions/0001-use-spring-boot.md
touch docs/decisions/0002-use-kafka-vs-rabbitmq.md

# 创建CI配置目录和文件
echo "创建CI配置文件..."
mkdir -p .github/workflows
touch .github/workflows/ci.yml
touch .gitlab-ci.yml

# 创建项目根文件
echo "创建项目根文件..."
touch README.md
# 如果当前目录没有LICENSE文件，复制过来；否则创建空文件
if [ -f "../LICENSE" ]; then
    cp ../LICENSE .
else
    touch LICENSE
fi

echo "目录结构创建完成！"
cd ..
echo "项目结构已生成在 $(pwd)/cloudshop 目录下"
