#!/bin/bash
set -e

COMPOSE_FILE="docker-compose.redis.yml"
WORK_DIR="$(dirname "$0")/../infrastructure/docker"

cd "$WORK_DIR"

case "$1" in
  start)
    echo "🚀 启动 Redis..."
    docker compose -f $COMPOSE_FILE up -d
    echo "✅ Redis 已启动"
    echo "👉 默认端口: 127.0.0.1:6379"
    ;;
  
  stop)
    echo "🛑 停止 Redis..."
    docker compose -f $COMPOSE_FILE down
    echo "✅ Redis 已停止"
    ;;
  
  restart)
    echo "🔄 重启 Redis..."
    docker compose -f $COMPOSE_FILE down
    docker compose -f $COMPOSE_FILE up -d
    echo "✅ Redis 已重启"
    echo "👉 默认端口: 127.0.0.1:6379"
    ;;
  
  clean)
    echo "🧹 停止并清理 Redis（包括数据卷）..."
    docker compose -f $COMPOSE_FILE down -v
    echo "✅ Redis 已彻底清理"
    ;;
  
  *)
    echo "用法: $0 {start|stop|restart|clean}"
    exit 1
    ;;
esac
