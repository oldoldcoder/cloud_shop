#!/bin/bash
set -e

echo "🛑 停止 Kafka + Zookeeper + RabbitMQ..."

cd "$(dirname "$0")/../infrastructure/docker"

docker compose -f docker-compose.kafka.rabbitmq.yml down

echo "✅ 服务已全部停止"
