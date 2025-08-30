#!/bin/bash
set -e

echo "🧹 停止并清理旧的 Kafka + Zookeeper + RabbitMQ..."
cd "$(dirname "$0")/../infrastructure/docker"

# 停止并删除容器和数据卷
docker compose -f docker-compose.kafka.rabbitmq.yml down -v || true

echo "🚀 启动新的 Kafka + Zookeeper + RabbitMQ..."
docker compose -f docker-compose.kafka.rabbitmq.yml up -d

echo "✅ 所有服务已重新启动"
echo "Kafka 地址: PLAINTEXT://127.0.0.1:9092"
echo "RabbitMQ 管理界面: http://127.0.0.1:15672 (账号 guest / 密码 guest)"
